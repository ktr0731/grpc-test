package server

import (
	"github.com/pkg/errors"
)

type opt struct {
	protocol Protocol
	addr     string

	tls        bool
	reflection bool
	verbose    bool
}

type Option func(*opt)

var defaultOption = opt{
	addr: ":50051",
}

type Protocol int

const (
	ProtocolGRPC Protocol = iota
	ProtocolImprobableGRPCWeb
	ProtocolUndefined
)

func WithProtocol(p Protocol) Option {
	if p >= ProtocolUndefined || p < ProtocolGRPC {
		panic(errors.Errorf("unknown protocol '%d'", p))
	}
	return func(o *opt) {
		o.protocol = p
	}
}

func WithAddr(addr string) Option {
	return func(o *opt) {
		o.addr = addr
	}
}

func WithTLS() Option {
	return func(o *opt) {
		o.tls = true
	}
}

func WithReflection() Option {
	return func(o *opt) {
		o.reflection = true
	}
}

func WithVerbose() Option {
	return func(o *opt) {
		o.verbose = true
	}
}
