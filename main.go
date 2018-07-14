package main

import (
	"flag"
	"net"

	"github.com/ktr0731/grpc-test/server"
)

func main() {
	web := flag.Bool("web", false, "works as a gRPC-Web server")
	flag.Parse()

	var l net.Listener
	var err error
	if !*web {
		l, err = net.Listen("tcp", ":50051")
		if err != nil {
			panic(err)
		}
	}

	defer server.New().Serve(l, *web).Stop()
	for {
		// do nothing
	}
}
