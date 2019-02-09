package main

import (
	"log"

	server "github.com/yuji0316/gRPC-sample/client/server"
	pb "github.com/yuji0316/gRPC-sample/hellovoicy"
)

const (
	host = "localhost"
	port = "50051"
)

func main() {
	// クライアントオブジェクト取得
	ctx, client, err := server.GetClient()
	if err != nil {
		log.Fatalf("client error: %v", err)
	}
	defer server.Close()

	// 処理実行
	r, err := client.Hello(ctx, &pb.HelloVoicyRequest{Message: "Hello"})
	if err != nil {
		log.Fatalf("execution error: %v", err)
	}

	log.Printf("%s", r.Message)
}
