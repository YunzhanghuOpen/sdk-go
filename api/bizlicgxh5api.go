package api

import (
	"context"
)

// BizlicGxH5APIService 云账户共享大额 H5+API
type BizlicGxH5APIService interface {
	// GxH5APIPreCollectBizlicMsg 工商实名信息录入
	GxH5APIPreCollectBizlicMsg(context.Context, *GxH5APIPreCollectBizlicMsgRequest) (*GxH5APIPreCollectBizlicMsgResponse, error)
	// GxH5APIGetStartUrl 预启动
	GxH5APIGetStartUrl(context.Context, *GxH5APIGetStartUrlRequest) (*GxH5APIGetStartUrlResponse, error)
	// GxH5APIEcoCityAicStatus 查询个体工商户状态
	GxH5APIEcoCityAicStatus(context.Context, *GxH5APIEcoCityAicStatusRequest) (*GxH5APIEcoCityAicStatusResponse, error)
}

// bizlicGxH5APIServiceImpl BizlicGxH5APIService 接口实现
type bizlicGxH5APIServiceImpl struct {
	cc Invoker
}

// NewBizlicGxH5APIService 创建客户端
func NewBizlicGxH5APIService(cc Invoker) BizlicGxH5APIService {
	return &bizlicGxH5APIServiceImpl{cc}
}

// GxH5APIPreCollectBizlicMsg 工商实名信息录入
func (c *bizlicGxH5APIServiceImpl) GxH5APIPreCollectBizlicMsg(ctx context.Context, in *GxH5APIPreCollectBizlicMsgRequest) (*GxH5APIPreCollectBizlicMsgResponse, error) {
	out := new(GxH5APIPreCollectBizlicMsgResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/aic/sharing-economy/api-h5/v1/collect", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GxH5APIGetStartUrl 预启动
func (c *bizlicGxH5APIServiceImpl) GxH5APIGetStartUrl(ctx context.Context, in *GxH5APIGetStartUrlRequest) (*GxH5APIGetStartUrlResponse, error) {
	out := new(GxH5APIGetStartUrlResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/aic/sharing-economy/api-h5/v1/h5url", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GxH5APIEcoCityAicStatus 查询个体工商户状态
func (c *bizlicGxH5APIServiceImpl) GxH5APIEcoCityAicStatus(ctx context.Context, in *GxH5APIEcoCityAicStatusRequest) (*GxH5APIEcoCityAicStatusResponse, error) {
	out := new(GxH5APIEcoCityAicStatusResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/aic/sharing-economy/api-h5/v1/status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GxH5APIPreCollectBizlicMsgRequest 工商实名信息录入请求
type GxH5APIPreCollectBizlicMsgRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业端的用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 身份证住址
	IDCardAddress string `json:"id_card_address,omitempty"`
	// 身份证签发机关
	IDCardAgency string `json:"id_card_agency,omitempty"`
	// 身份证民族
	IDCardNation string `json:"id_card_nation,omitempty"`
	// 身份证有效期开始时间
	IDCardValidityStart string `json:"id_card_validity_start,omitempty"`
	// 身份证有效期结束时间
	IDCardValidityEnd string `json:"id_card_validity_end,omitempty"`
}

// GxH5APIPreCollectBizlicMsgResponse 工商实名信息录入返回
type GxH5APIPreCollectBizlicMsgResponse struct {
	// 平台企业端的用户 ID
	DealerUserID string `json:"dealer_user_id,omitempty"`
}

// GxH5APIGetStartUrlRequest 预启动请求
type GxH5APIGetStartUrlRequest struct {
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

// GxH5APIGetStartUrlResponse 预启动返回
type GxH5APIGetStartUrlResponse struct {
	// 跳转 URL
	H5URL string `json:"h5_url,omitempty"`
}

// GxH5APIEcoCityAicStatusRequest 查询个体工商户状态请求
type GxH5APIEcoCityAicStatusRequest struct {
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

// GxH5APIEcoCityAicStatusResponse 查询个体工商户状态返回
type GxH5APIEcoCityAicStatusResponse struct {
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

// NotifyGxH5APIEcoCityAicRequest 个体工商户注册结果回调通知
type NotifyGxH5APIEcoCityAicRequest struct {
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
