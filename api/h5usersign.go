package api

import (
	"context"
)

// H5UserSignService H5 签约
type H5UserSignService interface {
	// H5UserPresign 预申请签约
	H5UserPresign(context.Context, *H5UserPresignRequest) (*H5UserPresignResponse, error)
	// H5UserSign 申请签约
	H5UserSign(context.Context, *H5UserSignRequest) (*H5UserSignResponse, error)
	// GetH5UserSignStatus 获取用户签约状态
	GetH5UserSignStatus(context.Context, *GetH5UserSignStatusRequest) (*GetH5UserSignStatusResponse, error)
	// H5UserRelease 用户解约（测试账号专用接口）
	H5UserRelease(context.Context, *H5UserReleaseRequest) (*H5UserReleaseResponse, error)
}

// h5UserSignServiceImpl H5UserSignService 接口实现
type h5UserSignServiceImpl struct {
	cc Invoker
}

// NewH5UserSignService 创建客户端
func NewH5UserSignService(cc Invoker) H5UserSignService {
	return &h5UserSignServiceImpl{cc}
}

// H5UserPresign 预申请签约
func (c *h5UserSignServiceImpl) H5UserPresign(ctx context.Context, in *H5UserPresignRequest) (*H5UserPresignResponse, error) {
	out := new(H5UserPresignResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/sdk/v1/presign", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// H5UserSign 申请签约
func (c *h5UserSignServiceImpl) H5UserSign(ctx context.Context, in *H5UserSignRequest) (*H5UserSignResponse, error) {
	out := new(H5UserSignResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/sdk/v1/sign/h5", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetH5UserSignStatus 获取用户签约状态
func (c *h5UserSignServiceImpl) GetH5UserSignStatus(ctx context.Context, in *GetH5UserSignStatusRequest) (*GetH5UserSignStatusResponse, error) {
	out := new(GetH5UserSignStatusResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/sdk/v1/sign/user/status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// H5UserRelease 用户解约（测试账号专用接口）
func (c *h5UserSignServiceImpl) H5UserRelease(ctx context.Context, in *H5UserReleaseRequest) (*H5UserReleaseResponse, error) {
	out := new(H5UserReleaseResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/sdk/v1/sign/release", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// H5UserPresignRequest 预申请签约请求
type H5UserPresignRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 证件类型  0：身份证 2：港澳居民来往内地通行证 3：护照 5：台湾居民来往大陆通行证 9：港澳居民居住证 10：台湾居民居住证 11：外国人永久居留身份证（外国人永久居留证） 15：中华人民共和国外国人工作许可证（A类） 16：中华人民共和国外国人工作许可证（B类） 17：中华人民共和国外国人工作许可证（C类） 18：港澳居民来往内地通行证（非中国国籍）
	CertificateType int32 `json:"certificate_type,omitempty"`
	// 是否收集手机号码 0：不收集（默认） 1：收集手机号码
	CollectPhoneNo int32 `json:"collect_phone_no,omitempty"`
	// 签约页面打开方式 1：微信小程序打开签约页面
	PageOpenWay int32 `json:"page_open_way,omitempty"`
}

// H5UserPresignResponse 预申请签约返回
type H5UserPresignResponse struct {
	// 用户 ID（废弃字段）
	Uid string `json:"uid,omitempty"`
	// H5 签约 token
	Token string `json:"token,omitempty"`
	// 签约状态
	Status int32 `json:"status,omitempty"`
}

// H5UserSignRequest 申请签约请求
type H5UserSignRequest struct {
	// H5 签约 token
	Token string `json:"token,omitempty"`
	// H5 页面主题颜色
	Color string `json:"color,omitempty"`
	// 回调 URL 地址
	URL string `json:"url,omitempty"`
	// 跳转 URL
	RedirectURL string `json:"redirect_url,omitempty"`
}

// H5UserSignResponse 申请签约返回
type H5UserSignResponse struct {
	// H5 签约页面 URL
	URL string `json:"url,omitempty"`
}

// GetH5UserSignStatusRequest 获取用户签约状态请求
type GetH5UserSignStatusRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
}

// GetH5UserSignStatusResponse 获取用户签约状态返回
type GetH5UserSignStatusResponse struct {
	// 签约时间
	SignedAt string `json:"signed_at,omitempty"`
	// 用户签约状态
	Status int32 `json:"status,omitempty"`
}

// H5UserReleaseRequest 用户解约（测试账号专用接口）请求
type H5UserReleaseRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 证件类型  0：身份证 2：港澳居民来往内地通行证 3：护照 5：台湾居民来往大陆通行证 9：港澳居民居住证 10：台湾居民居住证 11：外国人永久居留身份证（外国人永久居留证） 15：中华人民共和国外国人工作许可证（A类） 16：中华人民共和国外国人工作许可证（B类） 17：中华人民共和国外国人工作许可证（C类） 18：港澳居民来往内地通行证（非中国国籍）
	CertificateType int32 `json:"certificate_type,omitempty"`
}

// H5UserReleaseResponse 用户解约（测试账号专用接口）返回
type H5UserReleaseResponse struct {
	// 是否解约成功
	Status string `json:"status,omitempty"`
}

// NotifyH5UserSignRequest 签约回调
type NotifyH5UserSignRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 预签约手机号
	Phone string `json:"phone,omitempty"`
}
