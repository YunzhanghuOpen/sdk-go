package api

import (
	"context"
)

// FaceAuthService 人脸识别实名核验
type FaceAuthService interface {
	// ApplyFaceAuth 申请人脸识别实名核验
	ApplyFaceAuth(context.Context, *ApplyFaceAuthRequest) (*ApplyFaceAuthResponse, error)
	// GetFaceAuthResult 查询人脸识别实名核验结果
	GetFaceAuthResult(context.Context, *GetFaceAuthResultRequest) (*GetFaceAuthResultResponse, error)
}

// faceAuthServiceImpl FaceAuthService 接口实现
type faceAuthServiceImpl struct {
	cc Invoker
}

// NewFaceAuthService 创建客户端
func NewFaceAuthService(cc Invoker) FaceAuthService {
	return &faceAuthServiceImpl{cc}
}

// ApplyFaceAuth 申请人脸识别实名核验
func (c *faceAuthServiceImpl) ApplyFaceAuth(ctx context.Context, in *ApplyFaceAuthRequest) (*ApplyFaceAuthResponse, error) {
	out := new(ApplyFaceAuthResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/user/v1/face/auth", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetFaceAuthResult 查询人脸识别实名核验结果
func (c *faceAuthServiceImpl) GetFaceAuthResult(ctx context.Context, in *GetFaceAuthResultRequest) (*GetFaceAuthResultResponse, error) {
	out := new(GetFaceAuthResultResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/user/v1/face/auth_result", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplyFaceAuthRequest 申请人脸识别实名核验请求
type ApplyFaceAuthRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业实名核验 ID
	VerificationID string `json:"verification_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 回调地址
	CallbackURL string `json:"callback_url,omitempty"`
	// 跳转 URL
	RedirectURL string `json:"redirect_url,omitempty"`
	// 主题颜色
	Color string `json:"color,omitempty"`
}

// ApplyFaceAuthResponse 申请人脸识别实名核验返回
type ApplyFaceAuthResponse struct {
	// 人脸识别实名核验唯一 ID
	RecordID string `json:"record_id,omitempty"`
	// 平台企业实名核验 ID
	VerificationID string `json:"verification_id,omitempty"`
	// 人脸识别实名核验 H5 页面地址
	VerificationURL string `json:"verification_url,omitempty"`
}

// GetFaceAuthResultRequest 查询人脸识别实名核验结果请求
type GetFaceAuthResultRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 人脸识别实名核验唯一 ID
	RecordID string `json:"record_id,omitempty"`
	// 平台企业实名核验 ID
	VerificationID string `json:"verification_id,omitempty"`
}

// GetFaceAuthResultResponse 查询人脸识别实名核验结果返回
type GetFaceAuthResultResponse struct {
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 人脸识别实名核验唯一 ID
	RecordID string `json:"record_id,omitempty"`
	// 平台企业实名核验 ID
	VerificationID string `json:"verification_id,omitempty"`
	// 实名核验状态
	Status string `json:"status,omitempty"`
	// 实名核验完成时间
	VerifyTime string `json:"verify_time,omitempty"`
	// 实名核验失败详情
	Detail *FaceAuthDetail `json:"detail,omitempty"`
}

// FaceAuthDetail 人脸识别实名核验失败详情
type FaceAuthDetail struct {
	// 实名核验失败原因
	FailReason string `json:"fail_reason,omitempty"`
}

// NotifyFaceAuthRequest 人脸识别实名核验结果回调通知请求
type NotifyFaceAuthRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 人脸识别实名核验唯一 ID
	RecordID string `json:"record_id,omitempty"`
	// 平台企业实名核验 ID
	VerificationID string `json:"verification_id,omitempty"`
	// 实名核验状态
	Status string `json:"status,omitempty"`
	// 实名核验完成时间
	VerifyTime string `json:"verify_time,omitempty"`
	// 实名核验失败详情
	Detail *FaceAuthDetail `json:"detail,omitempty"`
}
