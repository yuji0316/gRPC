package auth

import (
	"context"

	"github.com/yuji0316/gRPC-sample/server/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// AuthenticationTokenKey 認証トークンの保存キー
	authenticationTokenKey = "authentication"
)

// NewAuthenticationIntercepter 認証用Intercepter作成
// 特定の認証トークンがメタデータに存在するか判定
// 全てのサービス呼び出し時に実行される
func NewAuthenticationIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 認証エラー時は不正アクセスを考慮して、ヒントを与えない為にも詳細なメッセージは返さない

	// メタデータから認証トークン取得
	token, ok := common.GetMetadataValue(ctx, authenticationTokenKey)
	if !ok {
		// metadata not found
		return nil, status.Error(codes.Unauthenticated, "")
	}

	// 取得した値が不正な場合はエラー
	if token != "authauth" {
		// Unauthenticated
		return nil, status.Error(codes.Unauthenticated, "")
	}

	return handler(ctx, req)
}
