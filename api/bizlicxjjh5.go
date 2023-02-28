package api

import (
	"context"
)

// BizlicXjjH5Service 云账户新经济 H5
type BizlicXjjH5Service interface {
	// H5GetStartUrl 预启动
	H5GetStartUrl(context.Context, *H5GetStartUrlRequest) (*H5GetStartUrlResponse, error)
	// H5EcoCityAicStatus 查询个体工商户状态
	H5EcoCityAicStatus(context.Context, *H5EcoCityAicStatusRequest) (*H5EcoCityAicStatusResponse, error)
}

// bizlicXjjH5ServiceImpl BizlicXjjH5Service 接口实现
type bizlicXjjH5ServiceImpl struct {
	cc Invoker
}

// NewBizlicXjjH5Service 创建客户端
func NewBizlicXjjH5Service(cc Invoker) BizlicXjjH5Service {
	return &bizlicXjjH5ServiceImpl{cc}
}

// H5GetStartUrl 预启动
func (c *bizlicXjjH5ServiceImpl) H5GetStartUrl(ctx context.Context, in *H5GetStartUrlRequest) (*H5GetStartUrlResponse, error) {
	out := new(H5GetStartUrlResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/aic/new-economy/h5/v1/h5url", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// H5EcoCityAicStatus 查询个体工商户状态
func (c *bizlicXjjH5ServiceImpl) H5EcoCityAicStatus(ctx context.Context, in *H5EcoCityAicStatusRequest) (*H5EcoCityAicStatusResponse, error) {
	out := new(H5EcoCityAicStatusResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/aic/new-economy/h5/v1/status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// H5GetStartUrlRequest 预启动请求
type H5GetStartUrlRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业端的用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
	// 客户端类型
	ClientType int32 `json:"client_type,omitempty"`
	// 异步通知 URL
	NotifyURL string `json:"notify_url,omitempty"`
	// H5 页面主题颜色
	Color string `json:"color,omitempty"`
	// 跳转 URL
	ReturnURL string `json:"return_url,omitempty"`
	// H5 页面 Title
	CustomerTitle int32 `json:"customer_title,omitempty"`
}

// H5GetStartUrlResponse 预启动返回
type H5GetStartUrlResponse struct {
	// 跳转 URL
	H5URL string `json:"h5_url,omitempty"`
}

// H5EcoCityAicStatusRequest 查询个体工商户状态请求
type H5EcoCityAicStatusRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业端的用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 用户唯一标识
	OpenID string `json:"open_id,omitempty"`
}

// H5EcoCityAicStatusResponse 查询个体工商户状态返回
type H5EcoCityAicStatusResponse struct {
	// 用户签约状态
	Status int32 `json:"status,omitempty"`
	// 注册状态描述
	StatusMessage string `json:"status_message,omitempty"`
	// 注册详情状态码
	StatusDetail int32 `json:"status_detail,omitempty"`
	// 注册详情状态码描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 注册发起时间
	ApplyedAt string `json:"applyed_at,omitempty"`
	// 注册完成时间
	RegistedAt string `json:"registed_at,omitempty"`
	// 统一社会信用代码
	Uscc string `json:"uscc,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
}

// NotifyH5EcoCityAicRequest 结果回调
type NotifyH5EcoCityAicRequest struct {
	// 用户唯一标识
	OpenID string `json:"open_id,omitempty"`
	// 平台企业端的用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
	// 注册/注销提交时间
	SubmitAt string `json:"submit_at,omitempty"`
	// 注册/注销完成时间
	RegistedAt string `json:"registed_at,omitempty"`
	// 用户签约状态
	Status int32 `json:"status,omitempty"`
	// 注册状态描述
	StatusMessage string `json:"status_message,omitempty"`
	// 注册详情状态码
	StatusDetail int32 `json:"status_detail,omitempty"`
	// 注册详情状态码描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 统一社会信用代码
	Uscc string `json:"uscc,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 回调类型
	Type int32 `json:"type,omitempty"`
}
