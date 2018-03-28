package main

import (
	"context"
	"flag"
	"log"
	"os"

	pb "github.com/LilyFaFa/hello-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// server address
	serverAddr := flag.String("server-addr", ":10000", "remote hello server address")
	// 解析 flag
	flag.Parse()
	// 日志抬头使用 LstdFlags
	// 日志默认输出到 Stderr
	logger := log.New(os.Stderr, "hello:", log.LstdFlags)
	// 创建 grpc 连接
	conn, err := grpc.Dial(
		*serverAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("can not connect remote server:%v", err)
	}
	// 断开连接
	defer conn.Close()

	clt := pb.NewGreeterClient(conn)
	req := &pb.Request{}
	res, err := clt.SayHello(context.Background(), req)
	if err != nil {
		logger.Fatalf("could not say hello:%v", err)
	}
	logger.Printf("remote server say:%s", res.Message)

}
