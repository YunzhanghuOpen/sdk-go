package api

import (
	"errors"

	"github.com/YunzhanghuOpen/sdk-go/core"
)

// Client 客户端
type Client struct {
	Payment                // 实时下单接口
	Tax                    // 个税服务接口
	Authentication         // 用户信息验证接口
	DataService            // 数据接口
	Invoice                // 发票接口
	ApiUserSignService     // API 签约
	H5UserSignService      // H5 签约
	UploadUserSignService  // 签约信息上传
	BizlicXjjH5Service     // 新经济个体户注册 H5
	BizlicXjjH5APIService  // 新经济个体户注册 H5+API
	BizlicGxV2H5Service    // 共享大额个体户注册 H5
	BizlicGxV2H5APIService // 共享大额个体户注册 H5+API
	CustomService          // 通用请求接口
	UserCollectService     // 用户信息收集
	CalculateLaborService  // 连续劳务税费试算
	RealNameService        // 实名信息收集
}

// Config Client 配置
type Config struct {
	Host       string // 请求域名，必传。参见“API 文档 > 接口定义 > 请求 URL”
	DealerID   string // 平台企业 ID，必传。登录云账户综合服务平台，选择“业务中心 > 业务管理 > 对接信息”获取
	PrivateKey string // 平台企业私钥，必传。同时登录云账户综合服务平台，选择“业务中心 > 业务管理 > 对接信息”，单击页面右上角的“编辑”，完成平台企业公钥配置
	AppKey     string // App Key，必传。登录云账户综合服务平台，选择“业务中心 > 业务管理 > 对接信息”获取
	Des3Key    string // 3DES Key，必传。登录云账户综合服务平台，选择“业务中心 > 业务管理 > 对接信息”获取
}

// New 新建 Client
func New(cfg *Config, options ...core.Option) (*Client, error) {
	if cfg == nil {
		return nil, errors.New("config is empty")
	}

	if cfg.Host == "" {
		return nil, errors.New("config.host is empty")
	}

	if cfg.DealerID == "" {
		return nil, errors.New("config.dealerID is empty")
	}

	if cfg.Des3Key == "" {
		return nil, errors.New("config.des3Key is empty")
	}

	opts := []core.Option{
		core.WithHost(cfg.Host),
		core.WithDealerID(cfg.DealerID),
		core.WithDes3Encoding(cfg.Des3Key),
	}
	opts = append(opts, options...)
	co, err := core.New(opts...)
	if err != nil {
		return nil, err
	}
	if !co.IsSetSigner() {
		core.WithRsaSign(cfg.PrivateKey, cfg.AppKey)(co)
	}
	return &Client{
		Payment:                NewPayment(co),
		Tax:                    NewTax(co),
		Authentication:         NewAuthentication(co),
		DataService:            NewDataService(co),
		Invoice:                NewInvoice(co),
		ApiUserSignService:     NewApiUserSignService(co),
		H5UserSignService:      NewH5UserSignService(co),
		UploadUserSignService:  NewUploadUserSignService(co),
		BizlicXjjH5Service:     NewBizlicXjjH5Service(co),
		BizlicXjjH5APIService:  NewBizlicXjjH5APIService(co),
		BizlicGxV2H5Service:    NewBizlicGxV2H5Service(co),
		BizlicGxV2H5APIService: NewBizlicGxV2H5APIService(co),
		CustomService:          NewCustomService(co),
		UserCollectService:     NewUserCollectService(co),
		CalculateLaborService:  NewCalculateLaborService(co),
		RealNameService:        NewRealNameService(co),
	}, nil
}
