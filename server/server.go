package server

import (
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/ktr0731/grpc-test/api"
	"google.golang.org/grpc"
)

type Server struct {
	s *grpc.Server
}

func New() *Server {
	s := grpc.NewServer()
	api.RegisterExampleServer(s, &ExampleService{})

	return &Server{
		s: s,
	}
}

func (s *Server) Serve(l net.Listener, web bool) error {
	if web {
		ws := grpcweb.WrapServer(s.s, grpcweb.WithWebsockets(false))
		mux := http.NewServeMux()
		mux.Handle("/", ws)
		srv := &http.Server{
			Addr:    ":50051",
			Handler: mux,
		}
		log.Println("works as a gRPC Web server")
		return srv.ListenAndServe()
	}

	log.Println("works as a gRPC server")
	return s.s.Serve(l)
}

type ExampleService struct {
	api.ExampleServer
}
