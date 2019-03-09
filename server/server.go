package server

import (
	"context"
	"crypto/tls"
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

var (
	waitTime = 1 * time.Second
)

func SetWaitTime(t time.Duration) {
	waitTime = t
}

type Server struct {
	startAsWeb bool

	logger *log.Logger

	sCloseCh <-chan error
	s        *grpc.Server

	wsCloseCh <-chan error
	ws        *http.Server

	useTLS bool
}

func New(verbose, useReflection, useTLS bool) *Server {
	var logWriter io.Writer
	if verbose {
		logWriter = os.Stderr
	} else {
		logWriter = ioutil.Discard
	}
	logger := log.New(logWriter, "[grpc-test] ", log.LstdFlags|log.Lshortfile)

	var opts []grpc.ServerOption
	if useTLS {
		tlsCfg := newTLSConfig()
		creds := credentials.NewTLS(tlsCfg)
		opts = append(opts, grpc.Creds(creds))
		logger.Println("TLS enabled")
	}
	s := grpc.NewServer(opts...)
	api.RegisterExampleServer(s, &ExampleService{})

	if useReflection {
		reflection.Register(s)
		logger.Println("gRPC reflection enabled")
	}

	return &Server{
		s:      s,
		logger: logger,
		useTLS: useTLS,
	}
}

func (s *Server) Serve(l net.Listener, web bool) *Server {
	if web {
		ws := grpcweb.WrapServer(s.s,
			grpcweb.WithWebsockets(true),
			grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
			grpcweb.WithCorsForRegisteredEndpointsOnly(false))
		mux := http.NewServeMux()
		mux.Handle("/", ws)
		s.ws = &http.Server{
			Addr:    ":50051",
			Handler: mux,
		}
		s.startAsWeb = true

		s.logger.Println("works as a gRPC Web server")
		closeCh := make(chan error)
		s.sCloseCh = closeCh
		go func() {
			if s.useTLS {
				s.ws.TLSConfig = newTLSConfig()
				closeCh <- s.ws.ListenAndServeTLS("", "")
			} else {
				closeCh <- s.ws.ListenAndServe()
			}
		}()

		return s
	}

	s.logger.Println("works as a gRPC server")
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

	if s.startAsWeb {
		return s.ws.Shutdown(context.Background())
	}
	s.s.GracefulStop()

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
