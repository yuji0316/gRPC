package common

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// GetMetadataValue メタデータから指定のキーに紐づく値を返す
func GetMetadataValue(ctx context.Context, key string) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	vs := md[key]
	if len(vs) == 0 {
		return "", false
	}
	return vs[0], true
}
