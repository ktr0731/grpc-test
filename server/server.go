package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/ktr0731/dept/logger"
	"github.com/ktr0731/grpc-test/api"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	_ "github.com/ktr0731/grpc-test/statik"
)

type Server struct {
	logger *log.Logger

	sCloseCh <-chan error
	s        *grpc.Server

	wsCloseCh <-chan error
	ws        *http.Server

	opts *opt
}

func New(opts ...Option) *Server {
	opt := defaultOption
	for _, o := range opts {
		o(&opt)
	}

	logger := newLogger(&opt)

	var grpcOpts []grpc.ServerOption
	if opt.tls {
		tlsCfg := newTLSConfig()
		creds := credentials.NewTLS(tlsCfg)
		grpcOpts = append(grpcOpts, grpc.Creds(creds))
		logger.Println("TLS enabled")
	}
	s := grpc.NewServer(grpcOpts...)
	api.RegisterExampleServer(s, &ExampleService{})

	if opt.reflection {
		reflection.Register(s)
		logger.Println("gRPC reflection enabled")
	}

	return &Server{
		s:      s,
		logger: logger,
		opts:   &opt,
	}
}

func (s *Server) Serve() *Server {
	addr := fmt.Sprintf("%s:%s", s.opts.host, s.opts.port)

	if isGRPCWeb(s.opts.protocol) {
		ws := grpcweb.WrapServer(s.s,
			grpcweb.WithWebsockets(true),
			grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
			grpcweb.WithCorsForRegisteredEndpointsOnly(false))
		mux := http.NewServeMux()
		mux.Handle("/", ws)
		s.ws = &http.Server{
			Addr:    addr,
			Handler: mux,
		}

		s.logger.Println("works as a gRPC-Web server")
		closeCh := make(chan error)
		s.sCloseCh = closeCh
		go func() {
			s.logger.Printf("listen at %s", addr)
			if s.opts.tls {
				panic("TODO: gRPC-Web + TLS is not supported yet")
			} else {
				closeCh <- s.ws.ListenAndServe()
			}
		}()

		return s
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Fatalf("failed to listen a tcp port for gRPC conn: %s", err)
	}
	s.logger.Println("works as a gRPC server")
	s.logger.Printf("listen at %s", addr)
	closeCh := make(chan error)
	s.wsCloseCh = closeCh
	go func() {
		closeCh <- s.s.Serve(l)
	}()

	return s
}

func (s *Server) Stop() error {
	// Receive the error if s.Serve has an error.
	select {
	case err := <-s.sCloseCh:
		return err
	case err := <-s.wsCloseCh:
		return err
	default:
		// no errors
	}

	if isGRPCWeb(s.opts.protocol) {
		return s.ws.Shutdown(context.Background())
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		s.s.GracefulStop()
	}()
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		s.logger.Println("graceful stop deadline exceeded. use Stop instead of GracefulStop.")
		s.s.Stop()
	}

	select {
	case err := <-s.sCloseCh:
		return err
	case err := <-s.wsCloseCh:
		return err
	}
}

type ExampleService struct {
	api.ExampleServer
}

func newTLSConfig() *tls.Config {
	statikFS, err := fs.New()
	if err != nil {
		logger.Fatal(err)
	}

	certPEM, err := statikFS.Open("/localhost.pem")
	if err != nil {
		logger.Fatal(err)
	}
	keyPEM, err := statikFS.Open("/localhost-key.pem")
	if err != nil {
		logger.Fatal(err)
	}
	certPEMBytes, err := ioutil.ReadAll(certPEM)
	if err != nil {
		logger.Fatal(err)
	}
	keyPEMBytes, err := ioutil.ReadAll(keyPEM)
	if err != nil {
		logger.Fatal(err)
	}
	cert, err := tls.X509KeyPair(certPEMBytes, keyPEMBytes)
	if err != nil {
		logger.Fatal(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

func newLogger(opt *opt) *log.Logger {
	var logWriter io.Writer
	if opt.verbose {
		logWriter = os.Stderr
	} else {
		logWriter = ioutil.Discard
	}
	return log.New(logWriter, "grpc-test: ", log.LstdFlags|log.Lshortfile)
}

func isGRPCWeb(p Protocol) bool {
	return p == ProtocolImprobableGRPCWeb
}
