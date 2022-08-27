package core

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
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

var sys_info string // 系统信息 eg:Darwin 21.6.0 arm64

func init() {
	// default
	sys_info = fmt.Sprintf("%s - %s", runtime.GOOS, runtime.GOARCH)

	if out, err := exec.Command("uname", "-s", "-r", "-m").CombinedOutput(); err == nil {
		sys_info = string(bytes.TrimSpace(out))
	}
}

var go_version = strings.Trim(runtime.Version(), "go") // go语言版本 eg:1.16

// userAgent
// 示例：yunzhanghu-sdk-go/1.0.0/Darwin 21.6.0 arm64/1.16"
func userAgent() string {
	return fmt.Sprintf("yunzhanghu-sdk-go/%s/%s/%s", Version, sys_info, go_version)
}
