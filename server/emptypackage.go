package server

import (
	"context"
	"fmt"
	"log"

	"github.com/ktr0731/grpc-test/api/emptypackage"
)

type EmptyPackageService struct {
	logger *log.Logger
	emptypackage.UnimplementedEmptyPackageServiceServer
}

func (s *EmptyPackageService) Unary(ctx context.Context, req *emptypackage.SimpleRequest) (*emptypackage.SimpleResponse, error) {
	return &emptypackage.SimpleResponse{
		Message: fmt.Sprintf("hello, %s", req.GetName()),
	}, nil
}
