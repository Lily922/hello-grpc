package main

import (
	"flag"

	"log"
	"net"
	"os"

	pb "github.com/LilyFaFa/hello-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	listenAddr := flag.String("listen-addr", ":10000", "hello server listen address")
	flag.Parse()
	logger := log.New(os.Stderr, "hello:", log.LstdFlags)
	logger.Printf("server is staring...")
	In, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		logger.Fatalf("could not listen:%v", err)

	}

	grpcServer := grpc.NewServer()
	helloServer, err := NewHelloServer()
	if err != nil {
		logger.Fatalf("%v", err)
	}

	pb.RegisterGreeterServer(grpcServer, helloServer)
	logger.Printf("server is listening on port:%s ...", *listenAddr)
	grpcServer.Serve(In)
}
