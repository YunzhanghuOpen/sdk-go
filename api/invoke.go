package api

import (
	"context"
)

// Invoker invoke 接口
type Invoker interface {
	Invoke(ctx context.Context, method string, urlStr string, respEncrypted bool, req interface{}, rsp interface{}) error
}
