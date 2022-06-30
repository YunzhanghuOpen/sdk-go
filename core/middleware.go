package core

import "context"

// Handler 中间件使用 Handler
type Handler func(ctx context.Context, req interface{}) (resp interface{}, err error)

// Middleware 中间件
type Middleware func(Handler) Handler

// Chain 中间件链
func Chain(m ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next
	}
}
