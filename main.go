package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ktr0731/grpc-test/server"
	"github.com/spf13/pflag"
)

func main() {
	web := pflag.Bool("web", false, "works as a gRPC-Web server")
	reflection := pflag.BoolP("reflection", "r", false, "use gRPC reflection")
	verbose := pflag.BoolP("verbose", "v", true, "verbose")
	tls := pflag.BoolP("tls", "t", false, "use TLS")
	cert := pflag.String("cert", "", "cert")
	certKey := pflag.String("cert-key", "", "cert key")
	rootCACert := pflag.String("root-ca-cert", "", "cert key")
	emptyPackageService := pflag.Bool("empty-package-service", false, "register empty package service")
	pflag.Parse()

	var opts []server.Option
	if *verbose {
		opts = append(opts, server.WithVerbose())
	}
	if *reflection {
		opts = append(opts, server.WithReflection())
	}
	if *tls {
		opts = append(opts, server.WithTLS())
	}
	if *web {
		opts = append(opts, server.WithProtocol(server.ProtocolImprobableGRPCWeb))
	}
	if len(*cert) != len(*certKey) {
		log.Fatalf("--cert and --cert-key are required")
	}
	if *cert != "" {
		opts = append(opts, server.WithCert(*cert, *certKey))
	}
	if *rootCACert != "" {
		opts = append(opts, server.WithRootCACert(*rootCACert))
	}
	if *emptyPackageService {
		opts = append(opts, server.WithEmptyPackageService())
	}

	defer server.New(opts...).Serve().Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig
	log.Println("signal received")
}
