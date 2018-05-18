package server

import (
	"context"
	"fmt"
	"net"
	"strings"

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

func (s *ExampleService) Unary(ctx context.Context, req *api.UnaryRequest) (*api.UnaryResponse, error) {
	return &api.UnaryResponse{
		Message: fmt.Sprintf("hello, %s", req.GetName()),
	}, nil
}

func (s *ExampleService) UnaryMessage(ctx context.Context, req *api.UnaryMessageRequest) (*api.UnaryResponse, error) {
	n := req.GetName()
	return &api.UnaryResponse{
		Message: fmt.Sprintf("hello, %s %s", n.GetFirstName(), n.GetLastName()),
	}, nil
}

func (s *ExampleService) UnaryRepeated(ctx context.Context, req *api.UnaryRepeatedRequest) (*api.UnaryResponse, error) {
	return &api.UnaryResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(req.GetName(), ", ")),
	}, nil
}

func (s *ExampleService) UnaryRepeatedMessage(ctx context.Context, req *api.UnaryRepeatedMessageRequest) (*api.UnaryResponse, error) {
	names := make([]string, 0, len(req.GetName()))
	for _, n := range req.GetName() {
		names = append(names, fmt.Sprintf("%s %s", n.GetFirstName(), n.GetLastName()))
	}
	return &api.UnaryResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(names, ", ")),
	}, nil
}

func (s *ExampleService) UnaryMap(ctx context.Context, req *api.UnaryMapRequest) (*api.UnaryResponse, error) {
	var names []string
	for k, v := range req.GetKvs() {
		names = append(names, fmt.Sprintf("%s (nickname: %s)", v, k))
	}
	return &api.UnaryResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(names, ", ")),
	}, nil
}

func (s *ExampleService) UnaryMapMessage(ctx context.Context, req *api.UnaryMapMessageRequest) (*api.UnaryResponse, error) {
	var names []string
	for k, v := range req.GetKvs() {
		names = append(names, fmt.Sprintf("%s %s (nickname: %s)", v.GetFirstName(), v.GetLastName(), k))
	}
	return &api.UnaryResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(names, ", ")),
	}, nil
}

func (s *ExampleService) UnaryOneof(ctx context.Context, req *api.UnaryOneofRequest) (*api.UnaryResponse, error) {
	var name string
	switch {
	case req.GetMsg() != nil:
		name = fmt.Sprintf("%s %s", req.GetMsg().GetFirstName(), req.GetMsg().GetLastName())
	default:
		name = req.GetPlain()
	}
	return &api.UnaryResponse{
		Message: fmt.Sprintf("hello, %s", name),
	}, nil
}

func (s *ExampleService) UnaryEnum(ctx context.Context, req *api.UnaryEnumRequest) (*api.UnaryResponse, error) {
	var msg string
	switch req.GetGender() {
	case api.Gender_Male:
		msg = "M"
	default:
		msg = "F"
	}
	return &api.UnaryResponse{
		Message: msg,
	}, nil
}
