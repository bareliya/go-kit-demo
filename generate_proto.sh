#!/bin/bash

# Run protoc to generate Go and gRPC code
#protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false ./demo_pb/demo.proto
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./demo_pb/demo.proto
