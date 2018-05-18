SHELL := /bin/bash

.PHONY: proto
proto:
	protoc --proto_path=$(GOPATH)/src --proto_path api --go_out=plugins=grpc:api api/api.proto
