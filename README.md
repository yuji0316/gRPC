# gRPC sample with Golang

## Create Environment

1. Get Golang library for gRPC

    go get -u google.golang.org/grpc

1. Install Protocol Buffers(for Mac)

    brew install protobuf

1. Install golang plugin of Protocol Buffers because basic Protocol Buffers doesn't support golang.

    go get -u github.com/golang/protobuf/protoc-gen-go

1. Set PATH to the `$GOPATH/bin`.

    export PATH=$PATH:$GOPATH/bin

## Create Golang gRPC source code from Protocol Buffers.

1. Defince gRPC interface by Protocol Buffers.  
[Source code](./hellovoicy/hellovoicy.proto)

1. Create Golan gRPC sourcde code form Protocol Buffers.

    protoc -I hellovoicy/ hellovoicy/hellovoicy.proto --go_out=plugins=grpc:hellovoicy  
    [Created code](./hellovoicy/hellovoicy.pb.go)

## Server side sample

[Source code](./server/main.go)

Execute

    go run server/main.go

## Client side sample

[Source code](./client/main.go)

Execute

    go run client/main.go

## Others

Refer to [this](https://github.com/grpc/grpc-go).

[This software includes the work that is distributed in the Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)