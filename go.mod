module github.com/ktr0731/grpc-test

require (
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/improbable-eng/grpc-web v0.9.1
	github.com/ktr0731/dept v0.1.1
	github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rakyll/statik v0.1.5
	github.com/rs/cors v1.6.0 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.3.0 // indirect
	go.uber.org/goleak v0.10.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1 // indirect
	google.golang.org/grpc v1.21.1
)

replace github.com/ktr0731/grpc-test/statik v0.0.0 => ./statik

go 1.13
