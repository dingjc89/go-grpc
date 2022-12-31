# How to use grpc in golang? I will introduce you for it step by step

## Step1
Use brew install protobuf, protoc-gen-go and protoc-gen-go-grpc

```shell
$ brew install protobuf
$ brew install protoc-gen-go
$ brew install protoc-gen-go-grpc
```

## Step2
Touch a file named as "service.proto", and put the following codes into it
```proto
syntax = "proto3";

package proto;

option go_package = "go-grpc/proto/service";

message HelloRequest {
    string name =1;
}

message HelloReply {
    string message=1;
}

service Greeter {
    rpc SayHello(HelloRequest) returns  (HelloReply) {}
    rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
}
```

## Step3
Generating client and server code. Details can be found in Makefile, if you are unfamiliar with its syntax, you can find related data at the https://www.gnu.org/software/make/manual/make.html website.

```shell
$ make protos
```

After the above command ended, two files were included in the relative path, and are named "service.pb.go" and "service_grpc.pb.go"

## Step4 
Creating the server

## Step5
Creating the client
