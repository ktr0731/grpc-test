package main

import (
	"net"

	"github.com/ktr0731/grpc-test/server"
)

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := server.New().Serve(l); err != nil {
		panic(err)
	}
}
