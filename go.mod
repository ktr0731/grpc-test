module github.com/ktr0731/grpc-test

require (
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/golang/protobuf v1.3.3
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/improbable-eng/grpc-web v0.12.0
	github.com/ktr0731/dept v0.1.3
	github.com/ktr0731/go-multierror v0.0.0-20171204182908-b7773ae21874 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.6
	github.com/rs/cors v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.3.0 // indirect
	go.uber.org/goleak v0.10.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sys v0.0.0-20200202164722-d101bd2416d5 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200204235621-fb4a7afc5178 // indirect
	google.golang.org/grpc v1.27.0
)

replace github.com/ktr0731/grpc-test/statik v0.0.0 => ./statik

go 1.13
