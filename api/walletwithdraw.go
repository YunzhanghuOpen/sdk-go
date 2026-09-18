package api

import (
	"context"
)

// WalletWithdrawService 钱包余额提现
type WalletWithdrawService interface {
	// CreateWalletWithdraw 发起钱包余额提现
	CreateWalletWithdraw(context.Context, *CreateWalletWithdrawRequest) (*CreateWalletWithdrawResponse, error)
	// QueryWalletWithdraw 查询钱包余额提现结果
	QueryWalletWithdraw(context.Context, *QueryWalletWithdrawRequest) (*QueryWalletWithdrawResponse, error)
	// CancelWalletWithdraw 取消挂起的钱包余额提现订单
	CancelWalletWithdraw(context.Context, *CancelWalletWithdrawRequest) (*CancelWalletWithdrawResponse, error)
	// RetryWalletWithdraw 重试挂起的钱包余额提现订单
	RetryWalletWithdraw(context.Context, *RetryWalletWithdrawRequest) (*RetryWalletWithdrawResponse, error)
	// GetWalletWithdrawReceiptFile 查询钱包余额提现电子回单
	GetWalletWithdrawReceiptFile(context.Context, *GetWalletWithdrawReceiptFileRequest) (*GetWalletWithdrawReceiptFileResponse, error)
}

// walletWithdrawServiceImpl WalletWithdrawService 接口实现
type walletWithdrawServiceImpl struct {
	cc Invoker
}

// NewWalletWithdrawService 创建客户端
func NewWalletWithdrawService(cc Invoker) WalletWithdrawService {
	return &walletWithdrawServiceImpl{cc}
}

// CreateWalletWithdraw 发起钱包余额提现
func (c *walletWithdrawServiceImpl) CreateWalletWithdraw(ctx context.Context, in *CreateWalletWithdrawRequest) (*CreateWalletWithdrawResponse, error) {
	out := new(CreateWalletWithdrawResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payout/v1/create", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryWalletWithdraw 查询钱包余额提现结果
func (c *walletWithdrawServiceImpl) QueryWalletWithdraw(ctx context.Context, in *QueryWalletWithdrawRequest) (*QueryWalletWithdrawResponse, error) {
	out := new(QueryWalletWithdrawResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payout/v1/query", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CancelWalletWithdraw 取消挂起的钱包余额提现订单
func (c *walletWithdrawServiceImpl) CancelWalletWithdraw(ctx context.Context, in *CancelWalletWithdrawRequest) (*CancelWalletWithdrawResponse, error) {
	out := new(CancelWalletWithdrawResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payout/v1/cancel-order", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetryWalletWithdraw 重试挂起的钱包余额提现订单
func (c *walletWithdrawServiceImpl) RetryWalletWithdraw(ctx context.Context, in *RetryWalletWithdrawRequest) (*RetryWalletWithdrawResponse, error) {
	out := new(RetryWalletWithdrawResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payout/v1/retry-order", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetWalletWithdrawReceiptFile 查询钱包余额提现电子回单
func (c *walletWithdrawServiceImpl) GetWalletWithdrawReceiptFile(ctx context.Context, in *GetWalletWithdrawReceiptFileRequest) (*GetWalletWithdrawReceiptFileResponse, error) {
	out := new(GetWalletWithdrawReceiptFileResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payout/v1/receipt-file", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateWalletWithdrawRequest 发起钱包余额提现请求
type CreateWalletWithdrawRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletWithdrawUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 平台企业的微信 AppID
	WxAppID string `json:"wx_app_id,omitempty"`
	// 提现金额
	Amount string `json:"amount,omitempty"`
	// 提现渠道
	Channel string `json:"channel,omitempty"`
	// 收款账号
	Account string `json:"account,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 通知地址
	NotifyURL string `json:"notify_url,omitempty"`
}

// CreateWalletWithdrawResponse 发起钱包余额提现返回
type CreateWalletWithdrawResponse struct {
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额提现订单号
	Ref string `json:"ref,omitempty"`
	// 提现金额
	Amount string `json:"amount,omitempty"`
}

// QueryWalletWithdrawRequest 查询钱包余额提现结果请求
type QueryWalletWithdrawRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 提现渠道
	Channel string `json:"channel,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额提现订单号
	Ref string `json:"ref,omitempty"`
}

// QueryWalletWithdrawResponse 查询钱包余额提现结果返回
type QueryWalletWithdrawResponse struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletWithdrawUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额提现订单号
	Ref string `json:"ref,omitempty"`
	// 提现金额
	Amount string `json:"amount,omitempty"`
	// 提现渠道
	Channel string `json:"channel,omitempty"`
	// 收款账号
	Account string `json:"account,omitempty"`
	// 平台企业的微信 AppID
	WxAppID string `json:"wx_app_id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
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
	// 劳动者实收金额
	UserReceivedAmount string `json:"user_received_amount,omitempty"`
	// 劳动者历史订单需补缴税费金额
	UserDebtRepaymentAmount string `json:"user_debt_repayment_amount,omitempty"`
	// 劳动者历史订单需补缴个税金额
	UserDebtRepaymentPersonalAmount string `json:"user_debt_repayment_personal_amount,omitempty"`
	// 劳动者历史订单需补缴增附税金额
	UserDebtRepaymentAddedAmount string `json:"user_debt_repayment_added_amount,omitempty"`
	// 钱包出账金额
	WalletOutflowAmount string `json:"wallet_outflow_amount,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletWithdrawWalletBalance `json:"wallet_balance,omitempty"`
}

// CancelWalletWithdrawRequest 取消挂起的钱包余额提现订单请求
type CancelWalletWithdrawRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额提现订单号
	Ref string `json:"ref,omitempty"`
}

// CancelWalletWithdrawResponse 取消挂起的钱包余额提现订单返回
type CancelWalletWithdrawResponse struct {
}

// RetryWalletWithdrawRequest 重试挂起的钱包余额提现订单请求
type RetryWalletWithdrawRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额提现订单号
	Ref string `json:"ref,omitempty"`
}

// RetryWalletWithdrawResponse 重试挂起的钱包余额提现订单返回
type RetryWalletWithdrawResponse struct {
}

// GetWalletWithdrawReceiptFileRequest 查询钱包余额提现电子回单请求
type GetWalletWithdrawReceiptFileRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包余额提现订单号
	Ref string `json:"ref,omitempty"`
	// 回单类型
	ReceiptType string `json:"receipt_type,omitempty"`
}

// GetWalletWithdrawReceiptFileResponse 查询钱包余额提现电子回单返回
type GetWalletWithdrawReceiptFileResponse struct {
	// 链接失效时间
	ExpireTime string `json:"expire_time,omitempty"`
	// 回单名
	FileName string `json:"file_name,omitempty"`
	// 下载链接
	URL string `json:"url,omitempty"`
}

// NotifyWalletWithdrawRequest 通知钱包余额提现结果回调通知请求
type NotifyWalletWithdrawRequest struct {
	// 通知类型
	NotifyType string `json:"notify_type,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletWithdrawUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 云账户钱包提现订单号
	Ref string `json:"ref,omitempty"`
	// 提现金额支付金额
	Amount string `json:"amount,omitempty"`
	// 提现渠道
	Channel string `json:"channel,omitempty"`
	// 收款账号
	Account string `json:"account,omitempty"`
	// 平台企业的微信 AppID
	WxAppID string `json:"wx_app_id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
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
	// 劳动者实收金额
	UserReceivedAmount string `json:"user_received_amount,omitempty"`
	// 劳动者历史订单需补缴税费金额
	UserDebtRepaymentAmount string `json:"user_debt_repayment_amount,omitempty"`
	// 劳动者历史订单需补缴个税金额
	UserDebtRepaymentPersonalAmount string `json:"user_debt_repayment_personal_amount,omitempty"`
	// 劳动者历史订单需补缴增附税金额
	UserDebtRepaymentAddedAmount string `json:"user_debt_repayment_added_amount,omitempty"`
	// 钱包出账金额
	WalletOutflowAmount string `json:"wallet_outflow_amount,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletWithdrawWalletBalance `json:"wallet_balance,omitempty"`
}

// WalletWithdrawUserInfo 劳动者信息
type WalletWithdrawUserInfo struct {
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 证件号
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 证件类型编码
	CardType string `json:"card_type,omitempty"`
}

// WalletWithdrawWalletBalance 钱包余额信息
type WalletWithdrawWalletBalance struct {
	// 钱包总余额
	TotalBalance string `json:"total_balance,omitempty"`
	// 可用余额
	AvailableBalance string `json:"available_balance,omitempty"`
	// 冻结余额
	FrozenBalance string `json:"frozen_balance,omitempty"`
	// 版本号
	Version string `json:"version,omitempty"`
}
