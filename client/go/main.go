package main

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/ktr0731/grpc-test/api"
	"google.golang.org/grpc"
)

var (
	server = flag.Bool("server", true, "server streaming")
	client = flag.Bool("client", false, "client streaming")
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	panicIfErr(err)

	flag.Parse()

	client := api.NewExampleClient(conn)
	if *server {
		req := &api.SimpleRequest{"ktr"}
		client.Unary()
		stream, err := client.ServerStreaming(context.Background(), req)
		panicIfErr(err)
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			panicIfErr(err)
			fmt.Println(res.GetMessage())
		}
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
