package main

import (
	"context"
	"log"
	"net"

	pb "github.com/yuji0316/gRPC/hellovoicy"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Protocol Buffersで定義された処理を実装するためのサーバーオブジェクト
type server struct{}

// サーバーオブジェクトに定義されたメソッドの処理を実装
func (s *server) Hello(ctx context.Context, in *pb.HelloVoicyRequest) (*pb.HelloVoicyReply, error) {
	return &pb.HelloVoicyReply{Message: in.Message + " World"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	// サーバーを起動
	s := grpc.NewServer()
	pb.RegisterHelloVoicyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("register serve error: %v", err)
	}
}
