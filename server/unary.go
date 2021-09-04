package server

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/ktr0731/grpc-test/api"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (s *ExampleService) Unary(ctx context.Context, req *api.SimpleRequest) (*api.SimpleResponse, error) {
	if err := grpc.SendHeader(ctx, metadata.New(map[string]string{"header_key1": "header_val1", "header_key2": "header_val2"})); err != nil {
		return nil, err
	}
	if err := grpc.SetTrailer(ctx, metadata.New(map[string]string{"trailer_key1": "trailer_val1", "trailer_key2": "trailer_val2"})); err != nil {
		return nil, err
	}
	return &api.SimpleResponse{
		Message: req.GetName(),
	}, nil
}

func (s *ExampleService) UnaryMessage(ctx context.Context, req *api.UnaryMessageRequest) (*api.SimpleResponse, error) {
	n := req.GetName()
	return &api.SimpleResponse{
		Message: fmt.Sprintf("%s %s", n.GetFirstName(), n.GetLastName()),
	}, nil
}

func (s *ExampleService) UnaryRepeated(ctx context.Context, req *api.UnaryRepeatedRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{
		Message: strings.Join(req.GetName(), ", "),
	}, nil
}

func (s *ExampleService) UnaryRepeatedMessage(ctx context.Context, req *api.UnaryRepeatedMessageRequest) (*api.SimpleResponse, error) {
	names := make([]string, 0, len(req.GetName()))
	for _, n := range req.GetName() {
		names = append(names, fmt.Sprintf("%s %s", n.GetFirstName(), n.GetLastName()))
	}
	return &api.SimpleResponse{
		Message: strings.Join(names, ", "),
	}, nil
}

func (s *ExampleService) UnaryRepeatedEnum(ctx context.Context, req *api.UnaryRepeatedEnumRequest) (*api.SimpleResponse, error) {
	var c1, c2 int
	for _, c := range req.GetChoices() {
		if c == api.Choices_Choice1 {
			c1++
		} else {
			c2++
		}
	}
	return &api.SimpleResponse{
		Message: fmt.Sprintf("1: %d, 2:%d", c1, c2),
	}, nil
}

func (s *ExampleService) UnarySelf(ctx context.Context, req *api.UnarySelfRequest) (*api.SimpleResponse, error) {
	var processPerson func(*api.Person) string
	processPerson = func(p *api.Person) string {
		txt := fmt.Sprintf("%s %s (%s), friends: ", p.GetName().GetFirstName(), p.GetName().GetLastName(), p.GetNickname())
		friends := p.GetFriends()
		names := make([]string, len(friends))
		for i, friend := range friends {
			names[i] = friend.GetNickname()
		}
		txt += strings.Join(names, ", ") + " → "
		for _, friend := range friends {
			txt += processPerson(friend)
		}
		return txt
	}
	return &api.SimpleResponse{
		Message: processPerson(req.GetYou()),
	}, nil
}

func (s *ExampleService) UnaryMap(ctx context.Context, req *api.UnaryMapRequest) (*api.SimpleResponse, error) {
	var kvs []string
	for k, v := range req.GetKvs() {
		kvs = append(kvs, fmt.Sprintf("key = %s, value = %s", k, v))
	}
	return &api.SimpleResponse{
		Message: fmt.Sprintf("passed kvs: %s", strings.Join(kvs, ", ")),
	}, nil
}

func (s *ExampleService) UnaryMapMessage(ctx context.Context, req *api.UnaryMapMessageRequest) (*api.SimpleResponse, error) {
	var names []string
	for k, v := range req.GetKvs() {
		names = append(names, fmt.Sprintf("%s %s (nickname: %s)", v.GetFirstName(), v.GetLastName(), k))
	}
	return &api.SimpleResponse{
		Message: strings.Join(names, ", "),
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
		Message: name,
	}, nil
}

func (s *ExampleService) UnaryEnum(ctx context.Context, req *api.UnaryEnumRequest) (*api.SimpleResponse, error) {
	var msg string
	switch req.GetChoice() {
	case api.Choices_Choice1:
		msg = "1"
	default:
		msg = "2"
	}
	return &api.SimpleResponse{
		Message: msg,
	}, nil
}

func (s *ExampleService) UnaryBytes(ctx context.Context, req *api.UnaryBytesRequest) (*api.SimpleResponse, error) {
	data := req.GetData()
	msg := fmt.Sprintf("received: (bytes) % x, (string) %s", data, data)
	s.Logger.Println(msg)
	return &api.SimpleResponse{
		Message: msg,
	}, nil
}

func (s *ExampleService) UnaryHeader(ctx context.Context, req *api.UnaryHeaderRequest) (*api.SimpleResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("cannot get headers")
	}
	var msg string
	for k, v := range md {
		msg += fmt.Sprintf("key = %s, val = %s\n", k, strings.Join(v, ", "))
	}
	return &api.SimpleResponse{
		Message: msg,
	}, nil
}

func (s *ExampleService) UnaryWithMapResponse(ctx context.Context, req *api.SimpleRequest) (*api.MapResponse, error) {
	m := map[string]*api.Name{}
	m[req.GetName()] = nil
	return &api.MapResponse{
		Names: m,
	}, nil
}

func (s *ExampleService) UnaryHeaderTrailer(ctx context.Context, req *api.SimpleRequest) (*api.SimpleResponse, error) {
	grpc.SendHeader(ctx, metadata.New(map[string]string{"header_key1": "header_val1", "header_key2": "header_val2"}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{"trailer_key1": "trailer_val1", "trailer_key2": "trailer_val2"}))
	return &api.SimpleResponse{Message: "response"}, nil
}

func (s *ExampleService) UnaryHeaderTrailerFailure(ctx context.Context, req *api.SimpleRequest) (*api.SimpleResponse, error) {
	grpc.SendHeader(ctx, metadata.New(map[string]string{"header_key1": "header_val1", "header_key2": "header_val2"}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{"trailer_key1": "trailer_val1", "trailer_key2": "trailer_val2"}))
	st := status.New(codes.Internal, "internal error")
	st, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "field",
					Description: "description",
				},
			},
		},
		&errdetails.PreconditionFailure{
			Violations: []*errdetails.PreconditionFailure_Violation{
				{
					Type:        "type",
					Subject:     "subject",
					Description: "description",
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return nil, st.Err()
}

func (s *ExampleService) UnaryEcho(ctx context.Context, req *api.UnaryMessageRequest) (*api.SimpleResponse, error) {
	b, err := protojson.Marshal(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.SimpleResponse{
		Message: string(b),
	}, nil
}
