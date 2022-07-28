package api

import (
	"errors"

	"github.com/YunzhanghuOpen/sdk-go/core"
)

// Client 客户端
type Client struct {
	Payment            // 实时下单接口
	Tax                // 个税服务接口
	Authentication     // 用户信息验证接口
	DataService        // 数据接口
	Invoice            // 发票接口
	ApiUserSignService // API签约
	H5UserSignService  // H5签约
}

// Config Client 配置
type Config struct {
	Host       string // 服务地址 必传  生产环境地址 https://api-service.yunzhanghu.com  沙箱环境地址 https://api-service.yunzhanghu.com/sandbox
	DealerID   string // 平台企业 ID 必传
	PrivateKey string // 平台企业 私钥 必传
	AppKey     string // 平台企业 appKey 必传
	Des3Key    string // 平台企业 des3Key 必传
}

// New 新建 Client
func New(cfg *Config, options ...core.Option) (*Client, error) {
	if cfg.Host == "" {
		return nil, errors.New("config.host is empty")
	}

	if cfg.DealerID == "" {
		return nil, errors.New("config.dealerID is empty")
	}

	if cfg.PrivateKey == "" {
		return nil, errors.New("config.privateKey is empty")
	}

	if cfg.AppKey == "" {
		return nil, errors.New("config.appKey is empty")
	}

	if cfg.Des3Key == "" {
		return nil, errors.New("config.des3Key is empty")
	}

	opts := []core.Option{
		core.WithHost(cfg.Host),
		core.WithDealerID(cfg.DealerID),
		core.WithDes3Encoding(cfg.Des3Key),
		core.WithRsaSign(cfg.PrivateKey, cfg.AppKey),
	}
	opts = append(opts, options...)
	co, err := core.New(opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		Payment:            NewPayment(co),
		Tax:                NewTax(co),
		Authentication:     NewAuthentication(co),
		DataService:        NewDataService(co),
		Invoice:            NewInvoice(co),
		ApiUserSignService: NewApiUserSignService(co),
		H5UserSignService:  NewH5UserSignService(co),
	}, nil
}
