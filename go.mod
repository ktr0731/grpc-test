module github.com/ktr0731/grpc-test

require (
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/golang/protobuf v1.5.1
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/improbable-eng/grpc-web v0.12.0
	github.com/jhump/protoreflect v1.8.2
	github.com/ktr0731/dept v0.1.3
	github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.6
	github.com/rs/cors v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5
	go.uber.org/goleak v0.10.0
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/ktr0731/grpc-test/statik v0.0.0 => ./statik

go 1.13
