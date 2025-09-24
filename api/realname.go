package api

import (
	"context"
)

// RealNameService 实名信息收集
type RealNameService interface {
	// CollectRealNameInfo 用户实名认证信息收集
	CollectRealNameInfo(context.Context, *CollectRealNameInfoRequest) (*CollectRealNameInfoResponse, error)
	// QueryRealNameInfo 用户实名认证信息查询
	QueryRealNameInfo(context.Context, *QueryRealNameInfoRequest) (*QueryRealNameInfoResponse, error)
}

// realNameServiceImpl RealNameService 接口实现
type realNameServiceImpl struct {
	cc Invoker
}

// NewRealNameService 创建客户端
func NewRealNameService(cc Invoker) RealNameService {
	return &realNameServiceImpl{cc}
}

// CollectRealNameInfo 用户实名认证信息收集
func (c *realNameServiceImpl) CollectRealNameInfo(ctx context.Context, in *CollectRealNameInfoRequest) (*CollectRealNameInfoResponse, error) {
	out := new(CollectRealNameInfoResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/user/v1/collect/realname/info", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryRealNameInfo 用户实名认证信息查询
func (c *realNameServiceImpl) QueryRealNameInfo(ctx context.Context, in *QueryRealNameInfoRequest) (*QueryRealNameInfoResponse, error) {
	out := new(QueryRealNameInfoResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/user/v1/query/realname/info", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type CollectRealNameInfoRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 证件号码
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 实名认证结果
	RealnameResult int32 `json:"realname_result,omitempty"`
	// 实名认证通过时间
	RealnameTime string `json:"realname_time,omitempty"`
	// 实名认证方式
	RealnameType int32 `json:"realname_type,omitempty"`
	// 实名认证唯一可追溯编码
	RealnameTraceID string `json:"realname_trace_id,omitempty"`
	// 认证平台
	RealnamePlatform string `json:"realname_platform,omitempty"`
	// 人脸照片
	FaceImage string `json:"face_image,omitempty"`
	// 人脸识别验证分数
	FaceVerifyScore string `json:"face_verify_score,omitempty"`
	// 银行卡号
	BankNo string `json:"bank_no,omitempty"`
	// 银行预留手机号
	BankPhone string `json:"bank_phone,omitempty"`
	// 平台企业审核人
	Reviewer string `json:"reviewer,omitempty"`
	// 人脸照片收集类型
	FaceImageCollectType int32 `json:"face_image_collect_type,omitempty"`
}

type CollectRealNameInfoResponse struct {
	// 录入状态
	Status string `json:"status,omitempty"`
}

type QueryRealNameInfoRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 证件号码
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
}

type QueryRealNameInfoResponse struct {
	// 实名认证结果
	RealnameResult int32 `json:"realname_result,omitempty"`
	// 实名认证通过时间
	RealnameTime string `json:"realname_time,omitempty"`
	// 实名认证方式
	RealnameType int32 `json:"realname_type,omitempty"`
	// 实名认证唯一可追溯编码
	RealnameTraceID string `json:"realname_trace_id,omitempty"`
	// 认证平台
	RealnamePlatform string `json:"realname_platform,omitempty"`
	// 是否存在人脸照片
	FaceImage string `json:"face_image,omitempty"`
	// 人脸识别验证分数
	FaceVerifyScore string `json:"face_verify_score,omitempty"`
	// 银行卡号
	BankNo string `json:"bank_no,omitempty"`
	// 银行预留手机号
	BankPhone string `json:"bank_phone,omitempty"`
	// 平台企业审核人
	Reviewer string `json:"reviewer,omitempty"`
}
