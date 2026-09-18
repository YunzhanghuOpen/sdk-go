package api

import (
	"context"
)

// WalletIncomeService 钱包余额入账
type WalletIncomeService interface {
	// CreateWalletIncome 发起钱包余额入账
	CreateWalletIncome(context.Context, *CreateWalletIncomeRequest) (*CreateWalletIncomeResponse, error)
	// QueryWalletIncome 查询钱包余额入账结果
	QueryWalletIncome(context.Context, *QueryWalletIncomeRequest) (*QueryWalletIncomeResponse, error)
	// CancelWalletIncome 取消钱包收入计税订单
	CancelWalletIncome(context.Context, *CancelWalletIncomeRequest) (*CancelWalletIncomeResponse, error)
	// RetryWalletIncome 重试挂起的计税订单
	RetryWalletIncome(context.Context, *RetryWalletIncomeRequest) (*RetryWalletIncomeResponse, error)
}

// walletIncomeServiceImpl WalletIncomeService 接口实现
type walletIncomeServiceImpl struct {
	cc Invoker
}

// NewWalletIncomeService 创建客户端
func NewWalletIncomeService(cc Invoker) WalletIncomeService {
	return &walletIncomeServiceImpl{cc}
}

// CreateWalletIncome 发起钱包余额入账
func (c *walletIncomeServiceImpl) CreateWalletIncome(ctx context.Context, in *CreateWalletIncomeRequest) (*CreateWalletIncomeResponse, error) {
	out := new(CreateWalletIncomeResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/income/v1/create", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryWalletIncome 查询钱包余额入账结果
func (c *walletIncomeServiceImpl) QueryWalletIncome(ctx context.Context, in *QueryWalletIncomeRequest) (*QueryWalletIncomeResponse, error) {
	out := new(QueryWalletIncomeResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/income/v1/query", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CancelWalletIncome 取消钱包收入计税订单
func (c *walletIncomeServiceImpl) CancelWalletIncome(ctx context.Context, in *CancelWalletIncomeRequest) (*CancelWalletIncomeResponse, error) {
	out := new(CancelWalletIncomeResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/income/v1/cancel-order", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetryWalletIncome 重试挂起的计税订单
func (c *walletIncomeServiceImpl) RetryWalletIncome(ctx context.Context, in *RetryWalletIncomeRequest) (*RetryWalletIncomeResponse, error) {
	out := new(RetryWalletIncomeResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/income/v1/retry-order", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateWalletIncomeRequest 发起钱包余额入账请求
type CreateWalletIncomeRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletIncomeUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台信息
	PlatformInfo *WalletIncomePlatformInfo `json:"platform_info,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 下单金额
	Amount string `json:"amount,omitempty"`
	// 获得收入时间
	EarnedAt string `json:"earned_at,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 通知地址
	NotifyURL string `json:"notify_url,omitempty"`
}

// CreateWalletIncomeResponse 发起钱包余额入账返回
type CreateWalletIncomeResponse struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包入账订单号
	Ref string `json:"ref,omitempty"`
	// 税前收入金额
	Amount string `json:"amount,omitempty"`
}

// QueryWalletIncomeRequest 查询钱包余额入账结果请求
type QueryWalletIncomeRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包入账订单号
	Ref string `json:"ref,omitempty"`
}

// QueryWalletIncomeResponse 查询钱包余额入账结果返回
type QueryWalletIncomeResponse struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletIncomeUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台信息
	PlatformInfo *WalletIncomePlatformInfo `json:"platform_info,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包入账订单号
	Ref string `json:"ref,omitempty"`
	// 下单金额
	Amount string `json:"amount,omitempty"`
	// 税前金额
	BeforeTaxAmount string `json:"before_tax_amount,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 获得收入时间
	EarnedAt string `json:"earned_at,omitempty"`
	// 订单状态
	Status string `json:"status,omitempty"`
	// 订单状态详情
	StatusDetail string `json:"status_detail,omitempty"`
	// 订单状态描述
	StatusMessage string `json:"status_message,omitempty"`
	// 订单状态详情描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 完成时间
	FinishedAt string `json:"finished_at,omitempty"`
	// 服务费信息
	FeeInfo *WalletIncomeFeeInfo `json:"fee_info,omitempty"`
	// 税费信息
	TaxInfo *WalletIncomeTaxInfo `json:"tax_info,omitempty"`
	// 劳动者历史订单需补缴税费金额
	UserDebtRepaymentAmount string `json:"user_debt_repayment_amount,omitempty"`
	// 劳动者历史订单需补缴个税金额
	UserDebtRepaymentPersonalAmount string `json:"user_debt_repayment_personal_amount,omitempty"`
	// 劳动者历史订单需补缴增附税金额
	UserDebtRepaymentAddedAmount string `json:"user_debt_repayment_added_amount,omitempty"`
	// 钱包入账金额
	WalletInflowAmount string `json:"wallet_inflow_amount,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletIncomeWalletBalance `json:"wallet_balance,omitempty"`
}

// CancelWalletIncomeRequest 取消钱包收入计税订单请求
type CancelWalletIncomeRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包入账订单号
	Ref string `json:"ref,omitempty"`
	// 取消钱包收入计税订单号
	CancelOrderID string `json:"cancel_order_id,omitempty"`
}

// CancelWalletIncomeResponse 取消钱包收入计税订单返回
type CancelWalletIncomeResponse struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包入账订单号
	Ref string `json:"ref,omitempty"`
	// 取消钱包收入计税订单号
	CancelOrderID string `json:"cancel_order_id,omitempty"`
	// 取消结果类型
	CancelResultType string `json:"cancel_result_type,omitempty"`
	// 取消明细
	CancelDetail *WalletIncomeCancelDetail `json:"cancel_detail,omitempty"`
}

// RetryWalletIncomeRequest 重试挂起的计税订单请求
type RetryWalletIncomeRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包入账订单号
	Ref string `json:"ref,omitempty"`
}

// RetryWalletIncomeResponse 重试挂起的计税订单返回
type RetryWalletIncomeResponse struct {
}

// NotifyWalletIncomeRequest 钱包余额入账结果回调通知请求
type NotifyWalletIncomeRequest struct {
	// 通知类型
	NotifyType string `json:"notify_type,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletIncomeUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台信息
	PlatformInfo *WalletIncomePlatformInfo `json:"platform_info,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包入账订单号
	Ref string `json:"ref,omitempty"`
	// 下单金额
	Amount string `json:"amount,omitempty"`
	// 税前金额
	BeforeTaxAmount string `json:"before_tax_amount,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 获得收入时间
	EarnedAt string `json:"earned_at,omitempty"`
	// 订单状态
	Status string `json:"status,omitempty"`
	// 订单状态详情
	StatusDetail string `json:"status_detail,omitempty"`
	// 订单状态描述
	StatusMessage string `json:"status_message,omitempty"`
	// 订单状态详情描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 完成时间
	FinishedAt string `json:"finished_at,omitempty"`
	// 服务费信息
	FeeInfo *WalletIncomeFeeInfo `json:"fee_info,omitempty"`
	// 税费信息
	TaxInfo *WalletIncomeTaxInfo `json:"tax_info,omitempty"`
	// 劳动者历史订单需补缴税费金额
	UserDebtRepaymentAmount string `json:"user_debt_repayment_amount,omitempty"`
	// 劳动者历史订单需补缴个税金额
	UserDebtRepaymentPersonalAmount string `json:"user_debt_repayment_personal_amount,omitempty"`
	// 劳动者历史订单需补缴增附税金额
	UserDebtRepaymentAddedAmount string `json:"user_debt_repayment_added_amount,omitempty"`
	// 钱包入账金额
	WalletInflowAmount string `json:"wallet_inflow_amount,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletIncomeWalletBalance `json:"wallet_balance,omitempty"`
}

// WalletIncomeUserInfo 劳动者信息
type WalletIncomeUserInfo struct {
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 证件号
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 证件类型编码
	CardType string `json:"card_type,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty" mask:"phone_no"`
}

// WalletIncomePlatformInfo 平台信息
type WalletIncomePlatformInfo struct {
	// 互联网平台名称
	PlatformName string `json:"platform_name,omitempty"`
	// 劳动者 ID
	UserID string `json:"user_id,omitempty"`
	// 劳动者名称或昵称
	UserNickname string `json:"user_nickname,omitempty"`
}

// WalletIncomeWalletBalance 钱包余额信息
type WalletIncomeWalletBalance struct {
	// 钱包总余额
	TotalBalance string `json:"total_balance,omitempty"`
	// 可用余额
	AvailableBalance string `json:"available_balance,omitempty"`
	// 冻结余额
	FrozenBalance string `json:"frozen_balance,omitempty"`
	// 版本号
	Version string `json:"version,omitempty"`
}

// WalletIncomeFeeInfo 服务费信息
type WalletIncomeFeeInfo struct {
	// 总服务费
	TotalFee string `json:"total_fee,omitempty"`
	// 总服务费率
	TotalFeeRate string `json:"total_fee_rate,omitempty"`
	// 平台企业加成服务费
	DealerFee string `json:"dealer_fee,omitempty"`
	// 平台企业加成服务费率
	DealerFeeRate string `json:"dealer_fee_rate,omitempty"`
	// 抵扣账户支付的加成服务费
	DealerDeductFee string `json:"dealer_deduct_fee,omitempty"`
	// 抵扣后应支付的加成服务费
	DealerPayableFee string `json:"dealer_payable_fee,omitempty"`
	// 劳动者加成服务费
	UserFee string `json:"user_fee,omitempty"`
	// 劳动者加成服务费率
	UserFeeRate string `json:"user_fee_rate,omitempty"`
}

// WalletIncomeTaxInfo 税费信息
type WalletIncomeTaxInfo struct {
	// 下单计税信息
	Original *WalletIncomeTaxDetail `json:"original,omitempty"`
	// 当前计税信息
	Current *WalletIncomeTaxDetail `json:"current,omitempty"`
}

// WalletIncomeTaxDetail 计税信息
type WalletIncomeTaxDetail struct {
	// 个税税率
	PersonalTaxRate string `json:"personal_tax_rate,omitempty"`
	// 个税速算扣除数
	DeductTax string `json:"deduct_tax,omitempty"`
	// 基本减除费用扣除
	BasicDeducted string `json:"basic_deducted,omitempty"`
	// 税费总额及明细
	Total *WalletIncomeTaxParty `json:"total,omitempty"`
	// 劳动者承担税费
	User *WalletIncomeTaxParty `json:"user,omitempty"`
	// 平台企业承担税费
	Dealer *WalletIncomeTaxParty `json:"dealer,omitempty"`
	// 云账户承担税费
	Broker *WalletIncomeTaxParty `json:"broker,omitempty"`
}

// WalletIncomeTaxParty 税费承担方明细
type WalletIncomeTaxParty struct {
	// 税费总额
	TotalTax string `json:"total_tax,omitempty"`
	// 个人所得税
	PersonalTax string `json:"personal_tax,omitempty"`
	// 增值税
	ValueAddedTax string `json:"value_added_tax,omitempty"`
	// 附加税
	AdditionalTax string `json:"additional_tax,omitempty"`
	// 城市维护建设税
	AdditionalUrbanTax string `json:"additional_urban_tax,omitempty"`
	// 教育费附加
	AdditionalTuitionTax string `json:"additional_tuition_tax,omitempty"`
	// 地方教育附加
	AdditionalLocalTuitionTax string `json:"additional_local_tuition_tax,omitempty"`
}

// WalletIncomeCancelDetail 取消计税处理明细
type WalletIncomeCancelDetail struct {
	// 服务费信息
	FeeInfo *WalletIncomeCancelFeeInfo `json:"fee_info,omitempty"`
	// 税费信息
	TaxInfo *WalletIncomeCancelTaxInfo `json:"tax_info,omitempty"`
	// 劳动者钱包扣减金额
	WalletOutflowAmount string `json:"wallet_outflow_amount,omitempty"`
	// 历史订单补缴税费退回金额
	UserDebtRepaymentAmount string `json:"user_debt_repayment_amount,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletIncomeWalletBalance `json:"wallet_balance,omitempty"`
}

// WalletIncomeCancelFeeInfo 取消计税退还服务费信息
type WalletIncomeCancelFeeInfo struct {
	// 总服务费
	TotalFee string `json:"total_fee,omitempty"`
	// 平台企业加成服务费
	DealerFee string `json:"dealer_fee,omitempty"`
	// 抵扣账户支付的加成服务费
	DealerDeductFee string `json:"dealer_deduct_fee,omitempty"`
	// 抵扣后应支付的加成服务费
	DealerPayableFee string `json:"dealer_payable_fee,omitempty"`
	// 劳动者加成服务费
	UserFee string `json:"user_fee,omitempty"`
}

// WalletIncomeCancelTaxInfo 取消计税退还税费信息
type WalletIncomeCancelTaxInfo struct {
	// 税费总额及明细
	Total *WalletIncomeTaxParty `json:"total,omitempty"`
	// 劳动者承担税费
	User *WalletIncomeTaxParty `json:"user,omitempty"`
	// 平台企业承担税费
	Dealer *WalletIncomeTaxParty `json:"dealer,omitempty"`
	// 云账户承担税费
	Broker *WalletIncomeTaxParty `json:"broker,omitempty"`
}
