package main

import (
	"context"

	pb "github.com/LilyFaFa/hello-grpc/proto"
	"google.golang.org/grpc"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (h *Client) connect(serverAddr string) (conn *grpc.ClientConn, err error) {
	if conn, err = grpc.Dial(serverAddr, grpc.WithInsecure()); err != nil {
		return nil, err
	}
	return conn, nil
}

// Executes tiller.ListReleases RPC.
func (h *Client) sayHello(serverAddr string, req *pb.Request) (*pb.Response, error) {
	c, err := h.connect(serverAddr)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	clt := pb.NewGreeterClient(c)
	res, err := clt.SayHello(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, err
}
