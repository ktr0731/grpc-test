package server

import (
	"net"

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

func (s *Server) Serve(l net.Listener) error {
	return s.s.Serve(l)
}

type ExampleService struct {
	api.ExampleServer
}
