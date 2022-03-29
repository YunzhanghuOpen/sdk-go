package core

import (
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/crypto"
)

// Option 配置
type Option func(*Core)

// Core client 用以实现 Invoker
type Core struct {
	httpClient  *http.Client
	err         error
	logger      Logger
	debug       bool
	host        string
	dealerID    string
	signType    string
	middlewares []Middleware
	signer      crypto.Signer
	encoder     *crypto.Des3Encoding
}

// New 新建 Core
func New(option ...Option) (*Core, error) {
	o := &Core{
		httpClient: &http.Client{},
		logger:     &defaultLogger{},
	}
	for _, opt := range option {
		opt(o)
	}
	if o.err != nil {
		return nil, o.err
	}
	return o, nil
}

// WithHost 设置访问地址
func WithHost(host string) Option {
	return func(o *Core) {
		o.host = host
	}
}

// WithMiddleware 中间件
func WithMiddleware(m ...Middleware) Option {
	return func(o *Core) {
		o.middlewares = m
	}
}

// EnDebug 开启debug模式
func EnDebug() Option {
	return func(o *Core) {
		o.debug = true
	}
}

// WithRsaSign rsa签名
func WithRsaSign(privateKey, appKey string) Option {
	return func(o *Core) {
		signer, err := crypto.NewRsaSigner(privateKey, appKey)
		if err != nil {
			o.err = fmt.Errorf("rsa privatekey wrong error: %w", err)
			return
		}
		o.signType = "rsa"
		o.signer = signer
	}
}

// WithHmacSign hmac签名
func WithHmacSign(appKey string) Option {
	return func(o *Core) {
		signer, err := crypto.NewHmacSigner(appKey)
		if err != nil {
			o.err = fmt.Errorf("rsa privatekey wrong error: %w", err)
			return
		}
		o.signType = "sha256"
		o.signer = signer
	}
}

// WithDealerID 设置dealerID
func WithDealerID(dealerID string) Option {
	return func(o *Core) {
		o.dealerID = dealerID
	}
}

// WithDes3Encoding 加解密
func WithDes3Encoding(des3Key string) Option {
	return func(o *Core) {
		o.encoder = crypto.NewDes3Encoding(des3Key)
	}
}

// WithHttpClient http client
func WithHttpClient(client *http.Client) Option {
	return func(o *Core) {
		o.httpClient = client
	}
}

// WithLogger 日志
func WithLogger(logger Logger) Option {
	return func(o *Core) {
		o.logger = logger
	}
}
