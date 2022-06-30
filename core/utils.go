package core

import (
	"context"
	"log"
)

// Logger 日志接口
type Logger interface {
	Logf(ctx context.Context, format string, args ...interface{})
}

type defaultLogger struct{}

// Log 默认日志
func (l defaultLogger) Logf(ctx context.Context, format string, args ...interface{}) {
	requestID := RequestID(ctx)
	if requestID != "" {
		format = "[" + requestID + "] " + format
	}
	log.Printf(format, args...)
}
