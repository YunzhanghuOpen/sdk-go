package core

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
)

// Logger 日志接口
type Logger interface {
	Logf(ctx context.Context, format string, args ...interface{})
}

type defaultLogger struct{}

// Logf 默认日志
func (l defaultLogger) Logf(ctx context.Context, format string, args ...interface{}) {
	requestID := RequestID(ctx)
	if requestID != "" {
		format = "[" + requestID + "] " + format
	}
	log.Printf(format, args...)
}

// userAgent
// 示例：yunzhanghu-sdk-go/1.0.0/1.20"
func userAgent() string {
	return fmt.Sprintf("yunzhanghu-sdk-go/%s/%s", Version, strings.Trim(runtime.Version(), "go"))
}
