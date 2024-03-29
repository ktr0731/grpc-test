SHELL := /bin/bash

export GOBIN := $(PWD)/_tools

.PHONY: tools
tools:
	@cat tools/tools.go | grep -E '^\s*_\s.*' | awk '{ print $$2 }' | xargs go install

.PHONY: proto
proto:
	protoc --proto_path=$(GOPATH)/src --proto_path api --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=api --go-grpc_out=api api/api.proto api/independent*.proto
	protoc --proto_path=$(GOPATH)/src --proto_path api --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=api/emptypackage --go-grpc_out=api/emptypackage api/emptypackage.proto
