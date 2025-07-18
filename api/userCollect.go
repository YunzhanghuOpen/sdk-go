package api

import (
	"context"
)

// UserCollectService 用户信息收集
type UserCollectService interface {
	// GetUserCollectPhoneStatus 查询手机号码绑定状态
	GetUserCollectPhoneStatus(context.Context, *GetUserCollectPhoneStatusRequest) (*GetUserCollectPhoneStatusResponse, error)
	// GetUserCollectPhoneUrl 获取收集手机号码页面
	GetUserCollectPhoneUrl(context.Context, *GetUserCollectPhoneUrlRequest) (*GetUserCollectPhoneUrlResponse, error)
}

// userCollectServiceImpl UserCollectService 接口实现
type userCollectServiceImpl struct {
	cc Invoker
}

// NewUserCollectService 创建客户端
func NewUserCollectService(cc Invoker) UserCollectService {
	return &userCollectServiceImpl{cc}
}

// GetUserCollectPhoneStatus 查询手机号码绑定状态
func (c *userCollectServiceImpl) GetUserCollectPhoneStatus(ctx context.Context, in *GetUserCollectPhoneStatusRequest) (*GetUserCollectPhoneStatusResponse, error) {
	out := new(GetUserCollectPhoneStatusResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/user/v1/collect/phone/status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserCollectPhoneUrl 获取收集手机号码页面
func (c *userCollectServiceImpl) GetUserCollectPhoneUrl(ctx context.Context, in *GetUserCollectPhoneUrlRequest) (*GetUserCollectPhoneUrlResponse, error) {
	out := new(GetUserCollectPhoneUrlResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/user/v1/collect/phone/url", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserCollectPhoneStatusRequest 查询手机号码绑定状态请求
type GetUserCollectPhoneStatusRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 证件类型编码
	CertificateType int32 `json:"certificate_type,omitempty"`
}

// GetUserCollectPhoneStatusResponse 查询手机号码绑定状态返回
type GetUserCollectPhoneStatusResponse struct {
	// 手机号码收集 Token
	Token string `json:"token,omitempty"`
	// 绑定状态
	Status int32 `json:"status,omitempty"`
}

// GetUserCollectPhoneUrlRequest 获取收集手机号码页面请求
type GetUserCollectPhoneUrlRequest struct {
	// 手机号码收集 Token
	Token string `json:"token,omitempty"`
	// 主题颜色
	Color string `json:"color,omitempty"`
	// 回调地址
	URL string `json:"url,omitempty"`
	// 跳转 URL
	RedirectURL string `json:"redirect_url,omitempty"`
}

// GetUserCollectPhoneUrlResponse 获取收集手机号码页面返回
type GetUserCollectPhoneUrlResponse struct {
	// 收集手机号码页面 URL
	URL string `json:"url,omitempty"`
}

// NotifyUserCollectPhoneRequest 收集手机号码结果回调通知
type NotifyUserCollectPhoneRequest struct {
	// 平台企业用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
	// 手机号码绑定状态
	Status int32 `json:"status,omitempty"`
}
