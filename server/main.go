package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "go-grpc/service"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello," + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again," + in.GetName()}, nil
}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.ErrorStatus, error) {
	return &pb.ErrorStatus{
		Message: in.Title,
		Details: nil,
	}, nil
}

func (s *server) GetList(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	result := new(pb.SearchResponse_Result)
	result.Title = "GRPC lesson"
	result.Url = "https://developers.google.com/protocol-buffers/docs/proto3"
	result.Snippets = []string{"lesson1", "lesson2", "lesson3"}

	return &pb.SearchResponse{
		Results: []*pb.SearchResponse_Result{result},
	}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
