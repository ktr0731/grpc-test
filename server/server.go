package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/ktr0731/grpc-test/api"
	"google.golang.org/grpc"
)

type Server struct {
	startAsWeb bool

	s  *grpc.Server
	ws *http.Server
}

func New() *Server {
	s := grpc.NewServer()
	api.RegisterExampleServer(s, &ExampleService{})

	return &Server{
		s: s,
	}
}

func (s *Server) Serve(l net.Listener, web bool) *Server {
	if web {
		ws := grpcweb.WrapServer(s.s, grpcweb.WithWebsockets(false))
		mux := http.NewServeMux()
		mux.Handle("/", ws)
		s.ws = &http.Server{
			Addr:    ":50051",
			Handler: mux,
		}
		s.startAsWeb = true

		log.Println("works as a gRPC Web server")
		go func() {
			if err := s.ws.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}()

		return s
	}

	log.Println("works as a gRPC server")
	go func() {
		if err := s.s.Serve(l); err != nil {
			log.Println(err)
		}
	}()
	return s
}

func (s *Server) Stop() error {
	if s.startAsWeb {
		return s.ws.Shutdown(context.Background())
	}
	s.s.GracefulStop()
	return nil
}

type ExampleService struct {
	api.ExampleServer
}
