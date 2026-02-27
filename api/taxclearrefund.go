package api

import (
	"context"
)

// TaxClearRefundService 连续劳务税费退补
type TaxClearRefundService interface {
	// GetClearTaxInfo 查询税费清缴完成结果
	GetClearTaxInfo(context.Context, *GetClearTaxInfoRequest) (*GetClearTaxInfoResponse, error)
	// GetClearTaxFile 查询税费清缴明细文件
	GetClearTaxFile(context.Context, *GetClearTaxFileRequest) (*GetClearTaxFileResponse, error)
	// GetRefundTaxInfo 查询税费退补完成结果
	GetRefundTaxInfo(context.Context, *GetRefundTaxInfoRequest) (*RefundTaxData, error)
	// GetRefundTaxLaborInfo 查询税费退补涉及劳动者
	GetRefundTaxLaborInfo(context.Context, *GetRefundTaxLaborInfoRequest) (*GetRefundTaxLaborInfoResponse, error)
}

// taxClearRefundImpl TaxClearRefundService 接口实现
type taxClearRefundImpl struct {
	cc Invoker
}

// NewTaxClearRefundService 创建客户端
func NewTaxClearRefundService(cc Invoker) TaxClearRefundService {
	return &taxClearRefundImpl{cc}
}

// GetClearTaxInfo 查询税费清缴完成结果
func (c *taxClearRefundImpl) GetClearTaxInfo(ctx context.Context, in *GetClearTaxInfoRequest) (*GetClearTaxInfoResponse, error) {
	out := new(GetClearTaxInfoResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/query-clear-tax", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetClearTaxFile 查询税费清缴明细文件
func (c *taxClearRefundImpl) GetClearTaxFile(ctx context.Context, in *GetClearTaxFileRequest) (*GetClearTaxFileResponse, error) {
	out := new(GetClearTaxFileResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/query-clear-tax/file", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetRefundTaxInfo 查询税费退补完成结果
func (c *taxClearRefundImpl) GetRefundTaxInfo(ctx context.Context, in *GetRefundTaxInfoRequest) (*RefundTaxData, error) {
	out := new(RefundTaxData)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/query-clear-status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetRefundTaxLaborInfo 查询税费退补涉及劳动者
func (c *taxClearRefundImpl) GetRefundTaxLaborInfo(ctx context.Context, in *GetRefundTaxLaborInfoRequest) (*GetRefundTaxLaborInfoResponse, error) {
	out := new(GetRefundTaxLaborInfoResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/query-clear-labor-info", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyClearTaxDoneRequest 连续劳务税费清缴完成通知
type NotifyClearTaxDoneRequest struct {
	// 通知 ID
	NotifyID string `json:"notify_id,omitempty"`
	// 通知时间
	NotifyTime string `json:"notify_time,omitempty"`
	// 返回数据
	Data *ClearTaxData `json:"data,omitempty"`
}

// ClearTaxData 连续劳务税费清缴完成数据
type ClearTaxData struct {
	// 报税属期
	TaxMonth string `json:"tax_month,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 清缴次数
	TaxClearNum string `json:"tax_clear_num,omitempty"`
	// 退补税用户数量
	RefundTaxLaborNum string `json:"refund_tax_labor_num,omitempty"`
	// 退补税订单数量
	RefundTaxOrderNum string `json:"refund_tax_order_num,omitempty"`
	// 订单总金额
	TotalAmount string `json:"total_amount,omitempty"`
	// 本批次退补税费总额
	CurTotalRefundTax string `json:"cur_total_refund_tax,omitempty"`
	// 退补税费总额
	TotalRefundTax string `json:"total_refund_tax,omitempty"`
	// 历史已退补税费总额
	HistoryRefundTax string `json:"history_refund_tax,omitempty"`
	// 本批次预扣税费总额
	TotalTax string `json:"total_tax,omitempty"`
	// 本批次实缴税费总额
	ReceiveTotalTax string `json:"receive_total_tax,omitempty"`
	// 本批次退补给用户税费总额
	CurTotalRefundLaborTax string `json:"cur_total_refund_labor_tax,omitempty"`
	// 本批次退补给平台企业税费总额
	CurTotalRefundDealerTax string `json:"cur_total_refund_dealer_tax,omitempty"`
	// 本批次退补给云账户税费总额
	CurTotalRefundBrokerTax string `json:"cur_total_refund_broker_tax,omitempty"`
	// 批次号
	BatchID string `json:"batch_id,omitempty"`
	// 批次生成时间
	BatchCreateTime string `json:"batch_create_time,omitempty"`
}

// GetClearTaxInfoRequest 查询税费清缴完成结果请求
type GetClearTaxInfoRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 报税属期
	TaxMonth string `json:"tax_month,omitempty"`
}

// GetClearTaxInfoResponse 查询税费清缴完成结果返回
type GetClearTaxInfoResponse struct {
	// 清缴批次列表
	BatchList []*ClearTaxData `json:"batch_list,omitempty"`
}

// GetClearTaxFileRequest 查询税费清缴明细文件请求
type GetClearTaxFileRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 报税属期
	TaxMonth string `json:"tax_month,omitempty"`
	// 批次号
	BatchID string `json:"batch_id,omitempty"`
}

// GetClearTaxFileResponse 查询税费清缴明细文件返回
type GetClearTaxFileResponse struct {
	// 下载地址
	URL string `json:"url,omitempty"`
}

// NotifyRefundTaxDoneRequest 连续劳务税费退补完成通知
type NotifyRefundTaxDoneRequest struct {
	// 通知 ID
	NotifyID string `json:"notify_id,omitempty"`
	// 通知时间
	NotifyTime string `json:"notify_time,omitempty"`
	// 返回数据
	Data *RefundTaxData `json:"data,omitempty"`
}

// RefundTaxData 连续劳务税费退补完成数据
type RefundTaxData struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 清缴次数
	TaxClearNum string `json:"tax_clear_num,omitempty"`
	// 报税属期
	TaxMonth string `json:"tax_month,omitempty"`
	// 退补税用户数量
	RefundTaxLaborNum string `json:"refund_tax_labor_num,omitempty"`
	// 退补税订单数量
	RefundTaxOrderNum string `json:"refund_tax_order_num,omitempty"`
	// 订单总金额
	TotalAmount string `json:"total_amount,omitempty"`
	// 本批次退补税费总额
	CurTotalRefundTax string `json:"cur_total_refund_tax,omitempty"`
	// 退补税费总额
	TotalRefundTax string `json:"total_refund_tax,omitempty"`
	// 历史已退补税费总额
	HistoryRefundTax string `json:"history_refund_tax,omitempty"`
	// 本批次预扣税费总额
	TotalTax string `json:"total_tax,omitempty"`
	// 本批次实缴税费总额
	ReceiveTotalTax string `json:"receive_total_tax,omitempty"`
	// 本批次退补给用户税费总额
	CurTotalRefundLaborTax string `json:"cur_total_refund_labor_tax,omitempty"`
	// 本批次退补给平台企业税费总额
	CurTotalRefundDealerTax string `json:"cur_total_refund_dealer_tax,omitempty"`
	// 本批次退补给云账户税费总额
	CurTotalRefundBrokerTax string `json:"cur_total_refund_broker_tax,omitempty"`
	// 批次号
	BatchID string `json:"batch_id,omitempty"`
	// 批次退补税状态
	BatchRefundTaxStatus string `json:"batch_refund_tax_status,omitempty"`
	// 批次生成时间
	BatchCreateTime string `json:"batch_create_time,omitempty"`
	// 批次退补税成功时间
	BatchRefundTaxFinishedTime string `json:"batch_refund_tax_finished_time,omitempty"`
	// 已完成税费退补的用户数量
	RefundTaxFinishedLaborNum string `json:"refund_tax_finished_labor_num,omitempty"`
	// 已完成的税费退补总额
	RefundTaxFinishedAmount string `json:"refund_tax_finished_amount,omitempty"`
}

// GetRefundTaxInfoRequest 查询税费退补完成结果请求
type GetRefundTaxInfoRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 报税属期
	TaxMonth string `json:"tax_month,omitempty"`
	// 批次号
	BatchID string `json:"batch_id,omitempty"`
}

// GetRefundTaxLaborInfoRequest 查询税费退补涉及劳动者请求
type GetRefundTaxLaborInfoRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 批次号
	BatchID string `json:"batch_id,omitempty"`
	// 税款所属期
	TaxMonth string `json:"tax_month,omitempty"`
	// 偏移量
	Offset int32 `json:"offset,omitempty"`
	// 每页返回条数
	Length int32 `json:"length,omitempty"`
}

// GetRefundTaxLaborInfoResponse 查询税费退补涉及劳动者返回
type GetRefundTaxLaborInfoResponse struct {
	// 税款所属期
	TaxMonth string `json:"tax_month,omitempty"`
	// 批次号
	BatchID string `json:"batch_id,omitempty"`
	// 批次生成时间
	BatchCreateTime string `json:"batch_create_time,omitempty"`
	// 退补税劳动者数量
	LaborNum string `json:"labor_num,omitempty"`
	// 退补税订单数量
	OrderNum string `json:"order_num,omitempty"`
	// 总数据条数
	TotalNum string `json:"total_num,omitempty"`
	// 退补税劳动者明细
	LaborRefundInfo []*LaborRefundInfo `json:"labor_refund_info,omitempty"`
}

// LaborRefundInfo 退补税劳动者明细
type LaborRefundInfo struct {
	// 劳动者姓名
	RealName string `json:"real_name,omitempty" mask:"real_name"`
	// 劳动者证件号
	IDCard string `json:"id_card,omitempty" mask:"id_card"`
	// 本批次退补给劳动者税费总额
	RefundTax string `json:"refund_tax,omitempty"`
	// 退补税状态
	TaxRefundStatus string `json:"tax_refund_status,omitempty"`
	// 劳动者收款账户
	ReceivingAccount string `json:"receiving_account,omitempty"`
	// 劳动者收款账号
	ReceivingChannel string `json:"receiving_channel,omitempty"`
	// 退补税费时间
	RefundTaxFinishedTime string `json:"refund_tax_finished_time,omitempty"`
}
