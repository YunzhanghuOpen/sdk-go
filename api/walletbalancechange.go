package api

// NotifyWalletBalanceChangeRequest 钱包余额变更结果回调通知请求
type NotifyWalletBalanceChangeRequest struct {
	// 通知类型
	NotifyType string `json:"notify_type,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 劳动者信息
	UserInfo *WalletBalanceQueryUserInfo `json:"user_info,omitempty"`
	// 钱包 ID
	WalletID string `json:"wallet_id,omitempty"`
	// 余额变化批次 ID
	ChangeID string `json:"change_id,omitempty"`
	// 钱包余额变更金额
	ChangeAmount string `json:"change_amount,omitempty"`
	// 需补缴个税
	UserDebtRepaymentPersonalAmount string `json:"user_debt_repayment_personal_amount,omitempty"`
	// 需补缴增附税
	UserDebtRepaymentAddedAmount string `json:"user_debt_repayment_added_amount,omitempty"`
	// 余额变化时间
	ChangedAt string `json:"changed_at,omitempty"`
	// 钱包余额信息
	WalletBalance *WalletBalanceQueryWalletBalance `json:"wallet_balance,omitempty"`
}
