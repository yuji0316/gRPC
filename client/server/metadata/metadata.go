package metadata

import "context"

// New Metadataオブジェクト作成
func New() Metadata {
	return Metadata{}
}

// Metadata gRPC通信時に送るメタデータ
type Metadata struct {
	Auth string
}

// GetRequestMetadata PerRPCCredentialsに定義されたメソッド
func (c *Metadata) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authentication": c.Auth,
	}, nil
}

// RequireTransportSecurity PerRPCCredentialsに定義されたメソッド
func (c *Metadata) RequireTransportSecurity() bool {
	return false
}
