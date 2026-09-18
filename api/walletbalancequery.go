package api

import (
	"context"
)

// WalletBalanceQueryService 钱包余额查询
type WalletBalanceQueryService interface {
	// QueryWalletBalance 查询钱包余额
	QueryWalletBalance(context.Context, *QueryWalletBalanceRequest) (*QueryWalletBalanceResponse, error)
}

// walletBalanceQueryServiceImpl WalletBalanceQueryService 接口实现
type walletBalanceQueryServiceImpl struct {
	cc Invoker
}

// NewWalletBalanceQueryService 创建客户端
func NewWalletBalanceQueryService(cc Invoker) WalletBalanceQueryService {
	return &walletBalanceQueryServiceImpl{cc}
}

// QueryWalletBalance 查询钱包余额
func (c *walletBalanceQueryServiceImpl) QueryWalletBalance(ctx context.Context, in *QueryWalletBalanceRequest) (*QueryWalletBalanceResponse, error) {
	out := new(QueryWalletBalanceResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/wallet/v1/balance", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryWalletBalanceRequest 查询钱包余额请求
type QueryWalletBalanceRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletBalanceQueryUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
}

// QueryWalletBalanceResponse 查询钱包余额返回
type QueryWalletBalanceResponse struct {
	// 钱包总余额
	TotalBalance string `json:"total_balance,omitempty"`
	// 可用余额
	AvailableBalance string `json:"available_balance,omitempty"`
	// 冻结余额
	FrozenBalance string `json:"frozen_balance,omitempty"`
	// 版本号
	Version string `json:"version,omitempty"`
}

// WalletBalanceQueryUserInfo 劳动者信息
type WalletBalanceQueryUserInfo struct {
	// 姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 证件号
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 证件类型编码
	CardType string `json:"card_type,omitempty"`
}

// WalletBalanceQueryWalletBalance 钱包余额信息
type WalletBalanceQueryWalletBalance struct {
	// 钱包总余额
	TotalBalance string `json:"total_balance,omitempty"`
	// 可用余额
	AvailableBalance string `json:"available_balance,omitempty"`
	// 冻结余额
	FrozenBalance string `json:"frozen_balance,omitempty"`
	// 版本号
	Version string `json:"version,omitempty"`
}
