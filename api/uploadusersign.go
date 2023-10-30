package api

import (
	"context"
)

// UploadUserSignService 签约信息上传
type UploadUserSignService interface {
	// UploadUserSign 用户签约信息上传
	UploadUserSign(context.Context, *UploadUserSignRequest) (*UploadUserSignResponse, error)
	// GetUploadUserSignStatus 获取用户签约状态
	GetUploadUserSignStatus(context.Context, *GetUploadUserSignStatusRequest) (*GetUploadUserSignStatusResponse, error)
}

// uploadUserSignServiceImpl UploadUserSignService 接口实现
type uploadUserSignServiceImpl struct {
	cc Invoker
}

// NewUploadUserSignService 创建客户端
func NewUploadUserSignService(cc Invoker) UploadUserSignService {
	return &uploadUserSignServiceImpl{cc}
}

// UploadUserSign 用户签约信息上传
func (c *uploadUserSignServiceImpl) UploadUserSign(ctx context.Context, in *UploadUserSignRequest) (*UploadUserSignResponse, error) {
	out := new(UploadUserSignResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/sign/user", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUploadUserSignStatus 获取用户签约状态
func (c *uploadUserSignServiceImpl) GetUploadUserSignStatus(ctx context.Context, in *GetUploadUserSignStatusRequest) (*GetUploadUserSignStatusResponse, error) {
	out := new(GetUploadUserSignStatusResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/sign/user/status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadUserSignRequest 用户签约信息上传请求
type UploadUserSignRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 是否是海外用户
	IsAbroad bool `json:"is_abroad,omitempty"`
	// 签约回调地址
	NotifyURL string `json:"notify_url,omitempty"`
}

// UploadUserSignResponse 用户签约信息上传返回
type UploadUserSignResponse struct {
	// 上传状态
	Status string `json:"status,omitempty"`
}

// GetUploadUserSignStatusRequest 获取用户签约状态请求
type GetUploadUserSignStatusRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
}

// GetUploadUserSignStatusResponse 获取用户签约状态返回
type GetUploadUserSignStatusResponse struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 签约状态
	Status int32 `json:"status,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// NotifyUploadUserSignRequest 签约成功状态回调通知
type NotifyUploadUserSignRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
}
