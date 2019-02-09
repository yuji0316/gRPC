package main

import (
	"context"
	"log"
	"time"

	pb "github.com/yuji0316/gRPC-sample/hellovoicy"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = "50051"
)

func main() {
	// サーバーへの接続設定
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloVoicyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 処理実行
	r, err := c.Hello(ctx, &pb.HelloVoicyRequest{Message: "Hello"})
	if err != nil {
		log.Fatalf("execution error: %v", err)
	}

	log.Printf("%s", r.Message)
}
