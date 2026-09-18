package api

import (
	"context"
)

// WalletDeductService 钱包余额扣减
type WalletDeductService interface {
	// CreateWalletDeduct 申请钱包余额扣减
	CreateWalletDeduct(context.Context, *CreateWalletDeductRequest) (*CreateWalletDeductResponse, error)
	// QueryWalletDeduct 查询钱包余额扣减申请结果
	QueryWalletDeduct(context.Context, *QueryWalletDeductRequest) (*QueryWalletDeductResponse, error)
	// CompleteWalletDeduct 提交钱包余额扣减结果
	CompleteWalletDeduct(context.Context, *CompleteWalletDeductRequest) (*CompleteWalletDeductResponse, error)
}

// walletDeductServiceImpl WalletDeductService 接口实现
type walletDeductServiceImpl struct {
	cc Invoker
}

// NewWalletDeductService 创建客户端
func NewWalletDeductService(cc Invoker) WalletDeductService {
	return &walletDeductServiceImpl{cc}
}

// CreateWalletDeduct 申请钱包余额扣减
func (c *walletDeductServiceImpl) CreateWalletDeduct(ctx context.Context, in *CreateWalletDeductRequest) (*CreateWalletDeductResponse, error) {
	out := new(CreateWalletDeductResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payout/v1/direct/create", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryWalletDeduct 查询钱包余额扣减申请结果
func (c *walletDeductServiceImpl) QueryWalletDeduct(ctx context.Context, in *QueryWalletDeductRequest) (*QueryWalletDeductResponse, error) {
	out := new(QueryWalletDeductResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payout/v1/direct/query", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompleteWalletDeduct 提交钱包余额扣减结果
func (c *walletDeductServiceImpl) CompleteWalletDeduct(ctx context.Context, in *CompleteWalletDeductRequest) (*CompleteWalletDeductResponse, error) {
	out := new(CompleteWalletDeductResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payout/v1/direct/complete", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateWalletDeductRequest 申请钱包余额扣减请求
type CreateWalletDeductRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletDeductUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 业务场景
	Scene string `json:"scene,omitempty"`
	// 申请扣减金额
	Amount string `json:"amount,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 通知地址
	NotifyURL string `json:"notify_url,omitempty"`
}

// CreateWalletDeductResponse 申请钱包余额扣减返回
type CreateWalletDeductResponse struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额扣减订单号
	Ref string `json:"ref,omitempty"`
	// 扣减金额
	Amount string `json:"amount,omitempty"`
}

// QueryWalletDeductRequest 查询钱包余额扣减申请结果请求
type QueryWalletDeductRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额扣减订单号
	Ref string `json:"ref,omitempty"`
}

// QueryWalletDeductResponse 查询钱包余额扣减申请结果返回
type QueryWalletDeductResponse struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletDeductUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 业务场景
	Scene string `json:"scene,omitempty"`
	// 云账户钱包余额扣减订单号
	Ref string `json:"ref,omitempty"`
	// 扣减金额
	Amount string `json:"amount,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 申请处理状态
	Status string `json:"status,omitempty"`
	// 申请处理状态详情
	StatusDetail string `json:"status_detail,omitempty"`
	// 申请处理状态描述
	StatusMessage string `json:"status_message,omitempty"`
	// 申请处理状态详情描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 处理完成时间
	FinishedAt string `json:"finished_at,omitempty"`
	// 劳动者实收金额
	UserReceivedAmount string `json:"user_received_amount,omitempty"`
	// 劳动者历史订单需补缴税费金额
	UserDebtRepaymentAmount string `json:"user_debt_repayment_amount,omitempty"`
	// 钱包出账金额
	WalletOutflowAmount string `json:"wallet_outflow_amount,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletDeductWalletBalance `json:"wallet_balance,omitempty"`
}

// CompleteWalletDeductRequest 提交钱包余额扣减结果请求
type CompleteWalletDeductRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额扣减订单号
	Ref string `json:"ref,omitempty"`
	// 结算状态
	Status string `json:"status,omitempty"`
	// 平台企业扣减交易流水号
	TradeNo string `json:"trade_no,omitempty"`
	// 支付完成时间
	FinishedAt string `json:"finished_at,omitempty"`
}

// CompleteWalletDeductResponse 提交钱包余额扣减结果返回
type CompleteWalletDeductResponse struct {
}

// NotifyWalletDeductRequest 钱包余额扣减申请结果回调通知请求
type NotifyWalletDeductRequest struct {
	// 通知类型
	NotifyType string `json:"notify_type,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletDeductUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 业务场景
	Scene string `json:"scene,omitempty"`
	// 云账户钱包余额扣减订单号
	Ref string `json:"ref,omitempty"`
	// 扣减金额
	Amount string `json:"amount,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 申请处理状态
	Status string `json:"status,omitempty"`
	// 申请处理状态详情
	StatusDetail string `json:"status_detail,omitempty"`
	// 申请处理状态描述
	StatusMessage string `json:"status_message,omitempty"`
	// 申请处理状态详情描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 处理完成时间
	FinishedAt string `json:"finished_at,omitempty"`
	// 劳动者实收金额
	UserReceivedAmount string `json:"user_received_amount,omitempty"`
	// 劳动者历史订单需补缴税费金额
	UserDebtRepaymentAmount string `json:"user_debt_repayment_amount,omitempty"`
	// 钱包出账金额
	WalletOutflowAmount string `json:"wallet_outflow_amount,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletDeductWalletBalance `json:"wallet_balance,omitempty"`
}

// WalletDeductUserInfo 劳动者信息
type WalletDeductUserInfo struct {
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 证件号
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 证件类型编码
	CardType string `json:"card_type,omitempty"`
}

// WalletDeductWalletBalance 钱包余额信息
type WalletDeductWalletBalance struct {
	// 钱包总余额
	TotalBalance string `json:"total_balance,omitempty"`
	// 可用余额
	AvailableBalance string `json:"available_balance,omitempty"`
	// 冻结余额
	FrozenBalance string `json:"frozen_balance,omitempty"`
	// 版本号
	Version string `json:"version,omitempty"`
}
