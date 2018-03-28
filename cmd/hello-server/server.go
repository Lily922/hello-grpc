package main

import (
	"golang.org/x/net/context"

	pb "github.com/LilyFaFa/hello-grpc/proto"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: "Hello",
	}, nil
}

func NewHelloServer() (*server, error) {
	return &server{}, nil
}
