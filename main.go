package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/ktr0731/grpc-test/server"
	"github.com/pkg/profile"
	"github.com/spf13/pflag"
)

func main() {
	defer profile.Start().Stop()

	web := pflag.Bool("web", false, "works as a gRPC-Web server")
	reflection := pflag.BoolP("reflection", "r", false, "use gRPC reflection")
	verbose := pflag.Bool("v", true, "verbose")
	waitTime := pflag.Duration("wait", 1*time.Second, "wait time for each server streaming response")
	tls := pflag.BoolP("tls", "t", false, "use TLS")
	pflag.Parse()

	var l net.Listener
	var err error
	if !*web {
		l, err = net.Listen("tcp", ":50051")
		if err != nil {
			panic(err)
		}
	}

	server.SetWaitTime(*waitTime)
	defer server.New(*verbose, *reflection, *tls).Serve(l, *web).Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig
	log.Println("signal received")
}
