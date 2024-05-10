package api

import (
	"context"
)

// BizlicGxV2H5Service 云账户共享大额 H5
type BizlicGxV2H5Service interface {
	// GxV2H5GetStartUrl 预启动
	GxV2H5GetStartUrl(context.Context, *GxV2H5GetStartUrlRequest) (*GxV2H5GetStartUrlResponse, error)
	// GxV2H5GetAicStatus 查询个体工商户状态
	GxV2H5GetAicStatus(context.Context, *GxV2H5GetAicStatusRequest) (*GxV2H5GetAicStatusResponse, error)
}

// bizlicGxV2H5ServiceImpl BizlicGxV2H5Service 接口实现
type bizlicGxV2H5ServiceImpl struct {
	cc Invoker
}

// NewBizlicGxV2H5Service 创建客户端
func NewBizlicGxV2H5Service(cc Invoker) BizlicGxV2H5Service {
	return &bizlicGxV2H5ServiceImpl{cc}
}

// GxV2H5GetStartUrl 预启动
func (c *bizlicGxV2H5ServiceImpl) GxV2H5GetStartUrl(ctx context.Context, in *GxV2H5GetStartUrlRequest) (*GxV2H5GetStartUrlResponse, error) {
	out := new(GxV2H5GetStartUrlResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/aic/sharing-economy/h5/v1/h5url", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GxV2H5GetAicStatus 查询个体工商户状态
func (c *bizlicGxV2H5ServiceImpl) GxV2H5GetAicStatus(ctx context.Context, in *GxV2H5GetAicStatusRequest) (*GxV2H5GetAicStatusResponse, error) {
	out := new(GxV2H5GetAicStatusResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/aic/sharing-economy/h5/v1/status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GxV2H5GetStartUrlRequest 预启动请求
type GxV2H5GetStartUrlRequest struct {
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

// GxV2H5GetStartUrlResponse 预启动返回
type GxV2H5GetStartUrlResponse struct {
	// 跳转 URL
	H5URL string `json:"h5_url,omitempty"`
}

// GxV2H5GetAicStatusRequest 查询个体工商户状态请求
type GxV2H5GetAicStatusRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 用户唯一标识
	OpenID string `json:"open_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 平台企业端的用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
}

// GxV2H5GetAicStatusResponse 查询个体工商户状态返回
type GxV2H5GetAicStatusResponse struct {
	// 用户注册状态
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

// NotifyGxV2H5AicRequest 个体工商户注册/注销结果回调通知
type NotifyGxV2H5AicRequest struct {
	// 用户唯一标识
	OpenID string `json:"open_id,omitempty"`
	// 平台企业端的用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
	// 注册/注销提交时间
	SubmitAt string `json:"submit_at,omitempty"`
	// 注册/注销完成时间
	RegistedAt string `json:"registed_at,omitempty"`
	// 用户注册/注销状态
	Status int32 `json:"status,omitempty"`
	// 注册/注销状态描述
	StatusMessage string `json:"status_message,omitempty"`
	// 注册/注销详情状态码
	StatusDetail int32 `json:"status_detail,omitempty"`
	// 注册/注销详情状态码描述
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
