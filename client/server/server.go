package grpc

import (
	"context"
	"time"

	"github.com/yuji0316/gRPC-sample/client/server/metadata"
	pb "github.com/yuji0316/gRPC-sample/hellovoicy"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = "50051"
)

var (
	conn   *grpc.ClientConn
	cancel *context.CancelFunc
	ctx    *context.Context
	client *pb.HelloVoicyClient
)

// GetClient gRPC呼び出し用クライアントを返す
func GetClient() (context.Context, pb.HelloVoicyClient, error) {
	if ctx != nil && client != nil {
		return *ctx, *client, nil
	}

	//　接続用メタデータ作成
	meta := metadata.New()
	meta.Auth = "authauth" //　認証トークン

	// サーバー接続設定
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(&meta),
	}
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return nil, nil, err
	}

	// クライアントオブジェクト作成
	cli := pb.NewHelloVoicyClient(conn)
	context, can := context.WithTimeout(context.Background(), time.Second)
	client = &cli
	ctx = &context
	cancel = &can

	return *ctx, *client, nil
}

// Close 終了処理
func Close() {
	if conn != nil {
		conn.Close()
	}
	if cancel != nil {
		(*cancel)()
	}
	conn = nil
	cancel = nil
	ctx = nil
	client = nil
}
