package core

import (
	"context"
)

type contextRequestIDKey struct{}

// WithRequestID 注入requestID
func WithRequestID(parent context.Context, requestID string) context.Context {
	return context.WithValue(parent, contextRequestIDKey{}, requestID)
}

// RequestID 获取requestID
func RequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(contextRequestIDKey{}).(string)
	if ok {
		return requestID
	}
	return ""
}
