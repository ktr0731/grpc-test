package server

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/ktr0731/grpc-test/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
}

func New(verbose, useReflection bool) *Server {
	s := grpc.NewServer()
	api.RegisterExampleServer(s, &ExampleService{})

	var logWriter io.Writer
	if verbose {
		logWriter = os.Stderr
	} else {
		logWriter = ioutil.Discard
	}
	logger := log.New(logWriter, "[grpc-test] ", log.LstdFlags)

	if useReflection {
		reflection.Register(s)
		logger.Println("gRPC reflection enabled")
	}

	return &Server{
		s:      s,
		logger: logger,
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
			closeCh <- s.ws.ListenAndServe()
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
