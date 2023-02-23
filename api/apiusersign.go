package api

import (
	"context"
)

// ApiUserSignService API签约
type ApiUserSignService interface {
	// ApiUseSignContract 获取协议预览 URL
	ApiUseSignContract(context.Context, *ApiUseSignContractRequest) (*ApiUseSignContractResponse, error)
	// ApiUserSign 用户签约
	ApiUserSign(context.Context, *ApiUserSignRequest) (*ApiUserSignResponse, error)
	// GetApiUserSignStatus 获取用户签约状态
	GetApiUserSignStatus(context.Context, *GetApiUserSignStatusRequest) (*GetApiUserSignStatusResponse, error)
	// ApiUserSignRelease 用户解约（测试账号专用接口）
	ApiUserSignRelease(context.Context, *ApiUserSignReleaseRequest) (*ApiUserSignReleaseResponse, error)
}

// apiUserSignServiceImpl ApiUserSignService 接口实现
type apiUserSignServiceImpl struct {
	cc Invoker
}

// NewApiUserSignService 创建客户端
func NewApiUserSignService(cc Invoker) ApiUserSignService {
	return &apiUserSignServiceImpl{cc}
}

// ApiUseSignContract 获取协议预览 URL
func (c *apiUserSignServiceImpl) ApiUseSignContract(ctx context.Context, in *ApiUseSignContractRequest) (*ApiUseSignContractResponse, error) {
	out := new(ApiUseSignContractResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/sign/v1/user/contract", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiUserSign 用户签约
func (c *apiUserSignServiceImpl) ApiUserSign(ctx context.Context, in *ApiUserSignRequest) (*ApiUserSignResponse, error) {
	out := new(ApiUserSignResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/sign/v1/user/sign", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetApiUserSignStatus 获取用户签约状态
func (c *apiUserSignServiceImpl) GetApiUserSignStatus(ctx context.Context, in *GetApiUserSignStatusRequest) (*GetApiUserSignStatusResponse, error) {
	out := new(GetApiUserSignStatusResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/sign/v1/user/status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiUserSignRelease 用户解约（测试账号专用接口）
func (c *apiUserSignServiceImpl) ApiUserSignRelease(ctx context.Context, in *ApiUserSignReleaseRequest) (*ApiUserSignReleaseResponse, error) {
	out := new(ApiUserSignReleaseResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/sign/v1/user/release", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiUseSignContractRequest 获取协议预览 URL 请求
type ApiUseSignContractRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
}

// ApiUseSignContractResponse 获取协议预览 URL 返回
type ApiUseSignContractResponse struct {
	// 预览跳转 URL
	URL string `json:"url,omitempty"`
	// 协议名称
	Title string `json:"title,omitempty"`
}

// ApiUserSignRequest 用户签约请求
type ApiUserSignRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 证件类型 idcard：身份证 passport：护照 mtphkm：港澳居民来往内地通行证  mtpt：台湾居民往来大陆通行证 rphkm：中华人民共和国港澳居民居住证 rpt：中华人民共和国台湾居民居住证 fpr：外国人永久居留身份证 ffwp：中华人民共和国外国人就业许可证书
	CardType string `json:"card_type,omitempty"`
}

// ApiUserSignResponse 用户签约返回
type ApiUserSignResponse struct {
	// 是否签约成功
	Status string `json:"status,omitempty"`
}

// GetApiUserSignStatusRequest 获取用户签约状态请求
type GetApiUserSignStatusRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
}

// GetApiUserSignStatusResponse 获取用户签约状态返回
type GetApiUserSignStatusResponse struct {
	// 签约时间
	SignedAt string `json:"signed_at,omitempty"`
	// 用户签约状态
	Status string `json:"status,omitempty"`
}

// ApiUserSignReleaseRequest 用户解约（测试账号专用接口）请求
type ApiUserSignReleaseRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 证件类型 idcard：身份证 passport：护照 mtphkm：港澳居民来往内地通行证  mtpt：台湾居民往来大陆通行证 rphkm：中华人民共和国港澳居民居住证 rpt：中华人民共和国台湾居民居住证 fpr：外国人永久居留身份证 ffwp：中华人民共和国外国人就业许可证书
	CardType string `json:"card_type,omitempty"`
}

// ApiUserSignReleaseResponse 用户解约（测试账号专用接口）返回
type ApiUserSignReleaseResponse struct {
	// 是否解约成功
	Status string `json:"status,omitempty"`
}
