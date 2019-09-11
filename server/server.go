package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
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

	s *grpc.Server

	ws *http.Server

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
		var tlsCfg *tls.Config
		if opt.cert != "" {
			certPEMBytes, err := ioutil.ReadFile(opt.cert)
			if err != nil {
				logger.Fatal(err)
			}
			keyPEMBytes, err := ioutil.ReadFile(opt.certKey)
			if err != nil {
				logger.Fatal(err)
			}
			cert, err := tls.X509KeyPair(certPEMBytes, keyPEMBytes)
			if err != nil {
				logger.Fatal(err)
			}
			tlsCfg = &tls.Config{Certificates: []tls.Certificate{cert}}
		} else {
			tlsCfg = newTLSConfig()
		}
		if opt.rootCACert != "" {
			b, err := ioutil.ReadFile(opt.rootCACert)
			if err != nil {
				logger.Fatal(err)
			}
			cp := x509.NewCertPool()
			if !cp.AppendCertsFromPEM(b) {
				logger.Fatal(err)
			}
			tlsCfg.RootCAs = cp
		}
		creds := credentials.NewTLS(tlsCfg)
		grpcOpts = append(grpcOpts, grpc.Creds(creds))
		logger.Println("TLS enabled")
	}

	s := grpc.NewServer(grpcOpts...)
	api.RegisterExampleServer(s, &ExampleService{logger: logger})

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
	if isGRPCWeb(s.opts.protocol) {
		ws := grpcweb.WrapServer(s.s,
			grpcweb.WithWebsockets(true),
			grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
			grpcweb.WithCorsForRegisteredEndpointsOnly(false))
		mux := http.NewServeMux()
		mux.Handle("/", ws)
		s.ws = &http.Server{
			Addr:    s.opts.addr,
			Handler: mux,
		}

		s.logger.Println("works as a gRPC-Web server")
		go func() {
			s.logger.Printf("listen at %s", s.opts.addr)
			if s.opts.tls {
				panic("TODO: gRPC-Web + TLS is not supported yet")
			} else {
				s.ws.ListenAndServe()
			}
		}()

		return s
	}

	l, err := net.Listen("tcp", s.opts.addr)
	if err != nil {
		s.logger.Fatalf("failed to listen a tcp port for gRPC conn: %s", err)
	}
	s.logger.Println("works as a gRPC server")
	s.logger.Printf("listen at %s", s.opts.addr)
	go func() {
		s.s.Serve(l)
	}()

	return s
}

func (s *Server) Stop() error {
	if isGRPCWeb(s.opts.protocol) {
		s.logger.Println("trying to shutdown the server")
		if err := s.ws.Shutdown(context.Background()); err != nil {
			return err
		}
		s.logger.Println("shutdown succeeded")
		return nil
	}

	s.logger.Println("trying to graceful shutdown the server")
	done := make(chan struct{})
	go func() {
		defer close(done)
		s.s.GracefulStop()
	}()
	select {
	case <-done:
		s.logger.Println("graceful shutdown succeeded")
	case <-time.After(3 * time.Second):
		s.logger.Println("graceful stop deadline exceeded. use Stop instead of GracefulStop.")
		s.s.Stop()
	}

	return nil
}

type ExampleService struct {
	logger *log.Logger
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
