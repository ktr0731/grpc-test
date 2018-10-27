package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/ktr0731/grpc-test/api"
)

func (s *ExampleService) Unary(ctx context.Context, req *api.SimpleRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{
		Message: fmt.Sprintf("hello, %s", req.GetName()),
	}, nil
}

func (s *ExampleService) UnaryMessage(ctx context.Context, req *api.UnaryMessageRequest) (*api.SimpleResponse, error) {
	n := req.GetName()
	return &api.SimpleResponse{
		Message: fmt.Sprintf("hello, %s %s", n.GetFirstName(), n.GetLastName()),
	}, nil
}

func (s *ExampleService) UnaryRepeated(ctx context.Context, req *api.UnaryRepeatedRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(req.GetName(), ", ")),
	}, nil
}

func (s *ExampleService) UnaryRepeatedMessage(ctx context.Context, req *api.UnaryRepeatedMessageRequest) (*api.SimpleResponse, error) {
	names := make([]string, 0, len(req.GetName()))
	for _, n := range req.GetName() {
		names = append(names, fmt.Sprintf("%s %s", n.GetFirstName(), n.GetLastName()))
	}
	return &api.SimpleResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(names, ", ")),
	}, nil
}

func (s *ExampleService) UnaryRepeatedEnum(ctx context.Context, req *api.UnaryRepeatedEnumRequest) (*api.SimpleResponse, error) {
	var m, f int
	for _, g := range req.GetGenders() {
		if g == api.Gender_Male {
			m++
		} else {
			f++
		}
	}
	return &api.SimpleResponse{
		Message: fmt.Sprintf("M: %d, F:%d", m, f),
	}, nil
}

func (s *ExampleService) UnarySelf(ctx context.Context, req *api.UnarySelfRequest) (*api.SimpleResponse, error) {
	person := req.GetYou()
	var txt string
	for {
		txt += fmt.Sprintf("%s %s (%s) - ", person.GetName().GetFirstName(), person.GetName().GetLastName(), person.GetNickname())
		person = person.GetNeighbor()
		if person == nil {
			break
		}
	}
	return &api.SimpleResponse{
		Message: txt,
	}, nil
}

func (s *ExampleService) UnaryMap(ctx context.Context, req *api.UnaryMapRequest) (*api.SimpleResponse, error) {
	var names []string
	for k, v := range req.GetKvs() {
		names = append(names, fmt.Sprintf("%s (nickname: %s)", v, k))
	}
	return &api.SimpleResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(names, ", ")),
	}, nil
}

func (s *ExampleService) UnaryMapMessage(ctx context.Context, req *api.UnaryMapMessageRequest) (*api.SimpleResponse, error) {
	var names []string
	for k, v := range req.GetKvs() {
		names = append(names, fmt.Sprintf("%s %s (nickname: %s)", v.GetFirstName(), v.GetLastName(), k))
	}
	return &api.SimpleResponse{
		Message: fmt.Sprintf("hello, %s", strings.Join(names, ", ")),
	}, nil
}

func (s *ExampleService) UnaryOneof(ctx context.Context, req *api.UnaryOneofRequest) (*api.SimpleResponse, error) {
	var name string
	switch {
	case req.GetMsg() != nil:
		name = fmt.Sprintf("%s %s", req.GetMsg().GetFirstName(), req.GetMsg().GetLastName())
	default:
		name = req.GetPlain()
	}
	return &api.SimpleResponse{
		Message: fmt.Sprintf("hello, %s", name),
	}, nil
}

func (s *ExampleService) UnaryEnum(ctx context.Context, req *api.UnaryEnumRequest) (*api.SimpleResponse, error) {
	var msg string
	switch req.GetGender() {
	case api.Gender_Male:
		msg = "M"
	default:
		msg = "F"
	}
	return &api.SimpleResponse{
		Message: msg,
	}, nil
}

func (s *ExampleService) UnaryBytes(ctx context.Context, req *api.UnaryBytesRequest) (*api.SimpleResponse, error) {
	data := req.GetData()
	return &api.SimpleResponse{
		Message: fmt.Sprintf("received: (bytes) % x, (string) %s", data, data),
	}, nil
}
