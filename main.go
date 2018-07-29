package main

import (
	"flag"
	"net"

	"github.com/ktr0731/grpc-test/server"
)

func main() {
	web := flag.Bool("web", false, "works as a gRPC-Web server")
	reflection := flag.Bool("r", false, "use gRPC reflection")
	flag.Parse()

	var l net.Listener
	var err error
	if !*web {
		l, err = net.Listen("tcp", ":50051")
		if err != nil {
			panic(err)
		}
	}

	defer server.New(*reflection).Serve(l, *web).Stop()
	for {
		// do nothing
	}
}
