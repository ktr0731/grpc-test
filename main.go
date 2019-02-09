package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/ktr0731/grpc-test/server"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	web := flag.Bool("web", false, "works as a gRPC-Web server")
	reflection := flag.Bool("r", false, "use gRPC reflection")
	verbose := flag.Bool("v", true, "verbose")
	waitTime := flag.Duration("wait", 1*time.Second, "wait time for each server streaming response")
	flag.Parse()

	var l net.Listener
	var err error
	if !*web {
		l, err = net.Listen("tcp", ":50051")
		if err != nil {
			panic(err)
		}
	}

	server.SetWaitTime(*waitTime)
	defer server.New(*verbose, *reflection).Serve(l, *web).Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig
	log.Println("signal received")
}
