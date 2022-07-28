package api

import (
	"context"
)

// Authentication 用户信息验证接口
type Authentication interface {
	// BankCardFourAuthVerify 银行卡四要素鉴权请求（下发短信验证码）
	BankCardFourAuthVerify(context.Context, *BankCardFourAuthVerifyRequest) (*BankCardFourAuthVerifyResponse, error)
	// BankCardFourAuthConfirm 银行卡四要素确认鉴权（上传短信验证码）
	BankCardFourAuthConfirm(context.Context, *BankCardFourAuthConfirmRequest) (*BankCardFourAuthConfirmResponse, error)
	// BankCardFourVerify 银行卡四要素验证
	BankCardFourVerify(context.Context, *BankCardFourVerifyRequest) (*BankCardFourVerifyResponse, error)
	// BankCardThreeVerify 银行卡三要素验证
	BankCardThreeVerify(context.Context, *BankCardThreeVerifyRequest) (*BankCardThreeVerifyResponse, error)
	// IDCardVerify 身份证实名验证
	IDCardVerify(context.Context, *IDCardVerifyRequest) (*IDCardVerifyResponse, error)
	// UserExemptedInfo 上传免验证用户名单信息
	UserExemptedInfo(context.Context, *UserExemptedInfoRequest) (*UserExemptedInfoResponse, error)
	// UserWhiteCheck 查看免验证用户名单是否存在
	UserWhiteCheck(context.Context, *UserWhiteCheckRequest) (*UserWhiteCheckResponse, error)
	// GetBankCardInfo 银行卡信息查询接口
	GetBankCardInfo(context.Context, *GetBankCardInfoRequest) (*GetBankCardInfoResponse, error)
}

// authenticationImpl Authentication 接口实现
type authenticationImpl struct {
	cc Invoker
}

// NewAuthentication 创建客户端
func NewAuthentication(cc Invoker) Authentication {
	return &authenticationImpl{cc}
}

// BankCardFourAuthVerify 银行卡四要素鉴权请求（下发短信验证码）
func (c *authenticationImpl) BankCardFourAuthVerify(ctx context.Context, in *BankCardFourAuthVerifyRequest) (*BankCardFourAuthVerifyResponse, error) {
	out := new(BankCardFourAuthVerifyResponse)
	err := c.cc.Invoke(ctx, "POST", "/authentication/verify-request", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankCardFourAuthConfirm 银行卡四要素确认鉴权（上传短信验证码）
func (c *authenticationImpl) BankCardFourAuthConfirm(ctx context.Context, in *BankCardFourAuthConfirmRequest) (*BankCardFourAuthConfirmResponse, error) {
	out := new(BankCardFourAuthConfirmResponse)
	err := c.cc.Invoke(ctx, "POST", "/authentication/verify-confirm", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankCardFourVerify 银行卡四要素验证
func (c *authenticationImpl) BankCardFourVerify(ctx context.Context, in *BankCardFourVerifyRequest) (*BankCardFourVerifyResponse, error) {
	out := new(BankCardFourVerifyResponse)
	err := c.cc.Invoke(ctx, "POST", "/authentication/verify-bankcard-four-factor", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankCardThreeVerify 银行卡三要素验证
func (c *authenticationImpl) BankCardThreeVerify(ctx context.Context, in *BankCardThreeVerifyRequest) (*BankCardThreeVerifyResponse, error) {
	out := new(BankCardThreeVerifyResponse)
	err := c.cc.Invoke(ctx, "POST", "/authentication/verify-bankcard-three-factor", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IDCardVerify 身份证实名验证
func (c *authenticationImpl) IDCardVerify(ctx context.Context, in *IDCardVerifyRequest) (*IDCardVerifyResponse, error) {
	out := new(IDCardVerifyResponse)
	err := c.cc.Invoke(ctx, "POST", "/authentication/verify-id", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserExemptedInfo 上传免验证用户名单信息
func (c *authenticationImpl) UserExemptedInfo(ctx context.Context, in *UserExemptedInfoRequest) (*UserExemptedInfoResponse, error) {
	out := new(UserExemptedInfoResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/user/exempted/info", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserWhiteCheck 查看免验证用户名单是否存在
func (c *authenticationImpl) UserWhiteCheck(ctx context.Context, in *UserWhiteCheckRequest) (*UserWhiteCheckResponse, error) {
	out := new(UserWhiteCheckResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/user/white/check", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetBankCardInfo 银行卡信息查询接口
func (c *authenticationImpl) GetBankCardInfo(ctx context.Context, in *GetBankCardInfoRequest) (*GetBankCardInfoResponse, error) {
	out := new(GetBankCardInfoResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/card", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankCardFourAuthVerifyRequest 银行卡四要素鉴权请求
type BankCardFourAuthVerifyRequest struct {
	// 银行卡号
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 银行预留手机号
	Mobile string `json:"mobile,omitempty"`
}

// BankCardFourAuthVerifyResponse 银行卡四要素鉴权返回
type BankCardFourAuthVerifyResponse struct {
	// 交易凭证
	Ref string `json:"ref,omitempty"`
}

// BankCardFourAuthConfirmRequest 银行卡四要素确认鉴权请求
type BankCardFourAuthConfirmRequest struct {
	// 银行卡号
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 银行预留手机号
	Mobile string `json:"mobile,omitempty"`
	// 短信验证码
	Captcha string `json:"captcha,omitempty"`
	// 交易凭证
	Ref string `json:"ref,omitempty"`
}

// BankCardFourAuthConfirmResponse 银行卡四要素确认鉴权返回
type BankCardFourAuthConfirmResponse struct {
}

// BankCardFourVerifyRequest 银行卡四要素验证请求
type BankCardFourVerifyRequest struct {
	// 银行卡号
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 银行预留手机号
	Mobile string `json:"mobile,omitempty"`
}

// BankCardFourVerifyResponse 银行卡四要素验证返回
type BankCardFourVerifyResponse struct {
}

// BankCardThreeVerifyRequest 银行卡三要素验证请求
type BankCardThreeVerifyRequest struct {
	// 银行卡号
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
}

// BankCardThreeVerifyResponse 银行卡三要素验证返回
type BankCardThreeVerifyResponse struct {
}

// IDCardVerifyRequest 身份证实名验证请求
type IDCardVerifyRequest struct {
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
}

// IDCardVerifyResponse 身份证实名验证返回
type IDCardVerifyResponse struct {
}

// UserExemptedInfoRequest 上传免验证用户名单信息请求
type UserExemptedInfoRequest struct {
	// 证件类型码
	CardType string `json:"card_type,omitempty"`
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 申请备注
	CommentApply string `json:"comment_apply,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 人员信息图片
	UserImages []string `json:"user_images,omitempty"`
	// 国别（地区）代码
	Country string `json:"country,omitempty"`
	// 出生日期
	Birthday string `json:"birthday,omitempty"`
	// 性别
	Gender string `json:"gender,omitempty"`
	// 回调地址
	NotifyURL string `json:"notify_url,omitempty"`
	// 请求流水号
	Ref string `json:"ref,omitempty"`
}

// UserExemptedInfoResponse 上传免验证用户名单信息返回
type UserExemptedInfoResponse struct {
	// 是否上传成功
	Ok string `json:"ok,omitempty"`
}

// UserWhiteCheckRequest 查看免验证用户名单是否存在请求
type UserWhiteCheckRequest struct {
	// 证件号码
	IDCard string `json:"id_card,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
}

// UserWhiteCheckResponse 查看免验证用户名单是否存在返回
type UserWhiteCheckResponse struct {
	Ok bool `json:"ok,omitempty"`
}

// GetBankCardInfoRequest 银行卡信息查询请求
type GetBankCardInfoRequest struct {
	// 银行卡号
	CardNo string `json:"card_no,omitempty"`
	// 银行名称
	BankName string `json:"bank_name,omitempty"`
}

// GetBankCardInfoResponse 银行卡信息查询返回
type GetBankCardInfoResponse struct {
	// 银行代码
	BankCode string `json:"bank_code,omitempty"`
	// 银行名称
	BankName string `json:"bank_name,omitempty"`
	// 卡类型
	CardType string `json:"card_type,omitempty"`
	// 云账户是否支持向该银行支付
	IsSupport bool `json:"is_support,omitempty"`
}
