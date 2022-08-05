package api

import (
	"context"
)

// Payment 实时下单接口
type Payment interface {
	// CreateBankpayOrder 银行卡实时下单
	CreateBankpayOrder(context.Context, *CreateBankpayOrderRequest) (*CreateBankpayOrderResponse, error)
	// CreateAlipayOrder 支付宝实时下单
	CreateAlipayOrder(context.Context, *CreateAlipayOrderRequest) (*CreateAlipayOrderResponse, error)
	// CreateWxpayOrder 微信实时下单
	CreateWxpayOrder(context.Context, *CreateWxpayOrderRequest) (*CreateWxpayOrderResponse, error)
	// GetOrder 查询单笔订单信息
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	// GetDealerVARechargeAccount 查询平台企业汇款信息
	GetDealerVARechargeAccount(context.Context, *GetDealerVARechargeAccountRequest) (*GetDealerVARechargeAccountResponse, error)
	// ListAccount 查询平台企业余额
	ListAccount(context.Context, *ListAccountRequest) (*ListAccountResponse, error)
	// GetEleReceiptFile 查询电子回单
	GetEleReceiptFile(context.Context, *GetEleReceiptFileRequest) (*GetEleReceiptFileResponse, error)
	// CancelOrder 取消待支付订单
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
}

// paymentImpl Payment 接口实现
type paymentImpl struct {
	cc Invoker
}

// NewPayment 创建客户端
func NewPayment(cc Invoker) Payment {
	return &paymentImpl{cc}
}

// CreateBankpayOrder 银行卡实时下单
func (c *paymentImpl) CreateBankpayOrder(ctx context.Context, in *CreateBankpayOrderRequest) (*CreateBankpayOrderResponse, error) {
	out := new(CreateBankpayOrderResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/order-bankpay", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateAlipayOrder 支付宝实时下单
func (c *paymentImpl) CreateAlipayOrder(ctx context.Context, in *CreateAlipayOrderRequest) (*CreateAlipayOrderResponse, error) {
	out := new(CreateAlipayOrderResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/order-alipay", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateWxpayOrder 微信实时下单
func (c *paymentImpl) CreateWxpayOrder(ctx context.Context, in *CreateWxpayOrderRequest) (*CreateWxpayOrderResponse, error) {
	out := new(CreateWxpayOrderResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/order-wxpay", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetOrder 查询单笔订单信息
func (c *paymentImpl) GetOrder(ctx context.Context, in *GetOrderRequest) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/query-order", in.DataType == "encryption", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDealerVARechargeAccount 查询平台企业汇款信息
func (c *paymentImpl) GetDealerVARechargeAccount(ctx context.Context, in *GetDealerVARechargeAccountRequest) (*GetDealerVARechargeAccountResponse, error) {
	out := new(GetDealerVARechargeAccountResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/va-account", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListAccount 查询平台企业余额
func (c *paymentImpl) ListAccount(ctx context.Context, in *ListAccountRequest) (*ListAccountResponse, error) {
	out := new(ListAccountResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/query-accounts", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetEleReceiptFile 查询电子回单
func (c *paymentImpl) GetEleReceiptFile(ctx context.Context, in *GetEleReceiptFileRequest) (*GetEleReceiptFileResponse, error) {
	out := new(GetEleReceiptFileResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/receipt/file", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CancelOrder 取消待支付订单
func (c *paymentImpl) CancelOrder(ctx context.Context, in *CancelOrderRequest) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/order/fail", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateBankpayOrderRequest 银行卡实时下单请求
type CreateBankpayOrderRequest struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 银行卡号
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
	// 订单备注
	PayRemark string `json:"pay_remark,omitempty"`
	// 回调地址
	NotifyURL string `json:"notify_url,omitempty"`
	// 项目标识
	ProjectID string `json:"project_id,omitempty"`
}

// CreateBankpayOrderResponse 银行卡实时下单返回
type CreateBankpayOrderResponse struct {
	OrderID string `json:"order_id,omitempty"`
	// 综合服务平台流水号
	Ref string `json:"ref,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
}

// CreateAlipayOrderRequest 支付宝实时下单请求
type CreateAlipayOrderRequest struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 支付宝账户
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
	// 订单备注
	PayRemark string `json:"pay_remark,omitempty"`
	// 回调地址
	NotifyURL string `json:"notify_url,omitempty"`
	// 项目标识
	ProjectID string `json:"project_id,omitempty"`
	// 校验支付宝账户姓名，固定值：Check
	CheckName string `json:"check_name,omitempty"`
}

// CreateAlipayOrderResponse 支付宝实时下单返回
type CreateAlipayOrderResponse struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 综合服务平台流水号
	Ref string `json:"ref,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
}

// CreateWxpayOrderRequest 微信实时下单请求
type CreateWxpayOrderRequest struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 微信用户 openid
	Openid string `json:"openid,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
	// 订单备注
	PayRemark string `json:"pay_remark,omitempty"`
	// 回调地址
	NotifyURL string `json:"notify_url,omitempty"`
	// 平台企业微信 AppID
	WxAppID string `json:"wx_app_id,omitempty"`
	// 微信支付模式，固定值：transfer
	WxpayMode string `json:"wxpay_mode,omitempty"`
	// 项目标识
	ProjectID string `json:"project_id,omitempty"`
	// 描述信息，该字段已废弃
	Notes string `json:"notes,omitempty"`
}

// CreateWxpayOrderResponse 微信实时下单返回
type CreateWxpayOrderResponse struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 综合服务平台流水号，唯一
	Ref string `json:"ref,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
}

// GetOrderRequest 查询单笔订单信息请求
type GetOrderRequest struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 支付路径名，银行卡（默认）、支付宝、微信
	Channel string `json:"channel,omitempty"`
	// 数据类型，如果为 encryption，则对返回的 data 进行加密
	DataType string `json:"data_type,omitempty"`
}

// GetOrderResponse 查询单笔订单信息返回
type GetOrderResponse struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 收款人账号
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty"`
	// 订单状态码
	Status string `json:"status,omitempty"`
	// 订单详细状态码
	StatusDetail string `json:"status_detail,omitempty"`
	// 订单状态码描述
	StatusMessage string `json:"status_message,omitempty"`
	// 订单详细状态码描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 综合服务主体支付金额
	BrokerAmount string `json:"broker_amount,omitempty"`
	// 综合服务平台流水号
	Ref string `json:"ref,omitempty"`
	// 支付交易流水号
	BrokerBankBill string `json:"broker_bank_bill,omitempty"`
	// 支付路径
	WithdrawPlatform string `json:"withdraw_platform,omitempty"`
	// 订单接收时间，精确到秒
	CreatedAt string `json:"created_at,omitempty"`
	// 订单完成时间，精确到秒
	FinishedTime string `json:"finished_time,omitempty"`
	// 综合服务主体加成服务费
	BrokerFee string `json:"broker_fee,omitempty"`
	// 余额账户支出加成服务费
	BrokerRealFee string `json:"broker_real_fee,omitempty"`
	// 抵扣账户支出加成服务费
	BrokerDeductFee string `json:"broker_deduct_fee,omitempty"`
	// 订单备注
	PayRemark string `json:"pay_remark,omitempty"`
	// 用户加成服务费
	UserFee string `json:"user_fee,omitempty"`
	// 银行名称
	BankName string `json:"bank_name,omitempty"`
	// 项目标识
	ProjectID string `json:"project_id,omitempty"`
	// 新就业形态劳动者 ID，该字段已废弃
	AnchorID string `json:"anchor_id,omitempty"`
	// 描述信息，该字段已废弃
	Notes string `json:"notes,omitempty"`
	// 系统支付金额，该字段已废弃
	SysAmount string `json:"sys_amount,omitempty"`
	// 税费，该字段已废弃
	Tax string `json:"tax,omitempty"`
	// 系统支付费用，该字段已废弃
	SysFee string `json:"sys_fee,omitempty"`
}

// GetDealerVARechargeAccountRequest 查询平台企业汇款信息请求
type GetDealerVARechargeAccountRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
}

// GetDealerVARechargeAccountResponse 查询平台企业汇款信息返回
type GetDealerVARechargeAccountResponse struct {
	// 账户名称
	AcctName string `json:"acct_name,omitempty"`
	// 专属账户
	AcctNo string `json:"acct_no,omitempty"`
	// 银行名称
	BankName string `json:"bank_name,omitempty"`
	// 付款账户
	DealerAcctName string `json:"dealer_acct_name,omitempty"`
}

// CancelOrderRequest 取消待支付订单请求
type CancelOrderRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 综合服务平台流水号
	Ref string `json:"ref,omitempty"`
	// 支付路径名，银行卡（默认）、支付宝、微信
	Channel string `json:"channel,omitempty"`
}

// CancelOrderResponse 取消待支付订单返回
type CancelOrderResponse struct {
	Ok string `json:"ok,omitempty"`
}

// ListAccountRequest 查询平台企业余额请求
type ListAccountRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
}

// ListAccountResponse 查询平台企业余额返回
type ListAccountResponse struct {
	DealerInfos []*AccountInfo `json:"dealer_infos,omitempty"`
}

// AccountInfo 账户信息
type AccountInfo struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 银行卡余额
	BankCardBalance string `json:"bank_card_balance,omitempty"`
	// 是否开通银行卡支付路径
	IsBankCard bool `json:"is_bank_card,omitempty"`
	// 支付宝余额
	AlipayBalance string `json:"alipay_balance,omitempty"`
	// 是否开通支付宝支付路径
	IsAlipay bool `json:"is_alipay,omitempty"`
	// 微信余额
	WxpayBalance string `json:"wxpay_balance,omitempty"`
	// 是否开通微信支付路径
	IsWxpay bool `json:"is_wxpay,omitempty"`
	// 加成服务费返点余额
	RebateFeeBalance string `json:"rebate_fee_balance,omitempty"`
	// 业务服务费余额
	AcctBalance string `json:"acct_balance,omitempty"`
	// 总余额
	TotalBalance string `json:"total_balance,omitempty"`
}

// GetEleReceiptFileRequest 查询电子回单请求
type GetEleReceiptFileRequest struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 综合服务平台流水号
	Ref string `json:"ref,omitempty"`
}

// GetEleReceiptFileResponse 查询电子回单返回
type GetEleReceiptFileResponse struct {
	// 链接失效时间
	ExpireTime string `json:"expire_time,omitempty"`
	// 回单名
	FileName string `json:"file_name,omitempty"`
	// 下载链接
	URL string `json:"url,omitempty"`
}

// NotifyOrderRequest 订单回调消息
type NotifyOrderRequest struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 收款人账号
	CardNo string `json:"card_no,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty"`
	// 订单状态码
	Status string `json:"status,omitempty"`
	// 订单详细状态码
	StatusDetail string `json:"status_detail,omitempty"`
	// 订单状态码描述
	StatusMessage string `json:"status_message,omitempty"`
	// 订单详细状态码描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 综合服务主体支付金额
	BrokerAmount string `json:"broker_amount,omitempty"`
	// 综合服务平台流水号
	Ref string `json:"ref,omitempty"`
	// 支付交易流水号
	BrokerBankBill string `json:"broker_bank_bill,omitempty"`
	// 支付路径
	WithdrawPlatform string `json:"withdraw_platform,omitempty"`
	// 订单接收时间，精确到秒
	CreatedAt string `json:"created_at,omitempty"`
	// 订单完成时间，精确到秒
	FinishedTime string `json:"finished_time,omitempty"`
	// 综合服务主体加成服务费
	BrokerFee string `json:"broker_fee,omitempty"`
	// 余额账户支出加成服务费
	BrokerRealFee string `json:"broker_real_fee,omitempty"`
	// 抵扣账户支出加成服务费
	BrokerDeductFee string `json:"broker_deduct_fee,omitempty"`
	// 订单备注
	PayRemark string `json:"pay_remark,omitempty"`
	// 用户加成服务费
	UserFee string `json:"user_fee,omitempty"`
	// 银行名称
	BankName string `json:"bank_name,omitempty"`
	// 项目标识
	ProjectID string `json:"project_id,omitempty"`
	// 新就业形态劳动者 ID，该字段已废弃
	AnchorID string `json:"anchor_id,omitempty"`
	// 描述信息，该字段已废弃
	Notes string `json:"notes,omitempty"`
	// 系统支付金额，该字段已废弃
	SysAmount string `json:"sys_amount,omitempty"`
	// 税费，该字段已废弃
	Tax string `json:"tax,omitempty"`
	// 系统支付费用，该字段已废弃
	SysFee string `json:"sys_fee,omitempty"`
}
