package api

import (
	"context"
)

// CalculateLaborService 连续劳务税费试算
type CalculateLaborService interface {
	// LaborCaculator 连续劳务税费试算（计算器）
	LaborCaculator(context.Context, *LaborCaculatorRequest) (*LaborCaculatorResponse, error)
	// CalcTax 订单税费试算
	CalcTax(context.Context, *CalcTaxRequest) (*CalcTaxResponse, error)
}

// calculateLaborServiceImpl CalculateLaborService 接口实现
type calculateLaborServiceImpl struct {
	cc Invoker
}

// NewCalculateLaborService 创建客户端
func NewCalculateLaborService(cc Invoker) CalculateLaborService {
	return &calculateLaborServiceImpl{cc}
}

// LaborCaculator 连续劳务税费试算（计算器）
func (c *calculateLaborServiceImpl) LaborCaculator(ctx context.Context, in *LaborCaculatorRequest) (*LaborCaculatorResponse, error) {
	out := new(LaborCaculatorResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/tax/v1/labor-caculator", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcTax 订单税费试算
func (c *calculateLaborServiceImpl) CalcTax(ctx context.Context, in *CalcTaxRequest) (*CalcTaxResponse, error) {
	out := new(CalcTaxResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/payment/v1/calc-tax", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaborCaculatorRequest 连续劳务税费试算（计算器）请求
type LaborCaculatorRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 月度收入列表
	MonthSettlementList []*MonthSettlement `json:"month_settlement_list,omitempty"`
}

// MonthSettlement 月度收入
type MonthSettlement struct {
	// 月份
	Month int32 `json:"month,omitempty"`
	// 月度收入
	MonthPreTaxAmount string `json:"month_pre_tax_amount,omitempty"`
}

// LaborCaculatorResponse 连续劳务税费试算（计算器）返回
type LaborCaculatorResponse struct {
	// 综合所得汇算清缴
	YearTaxInfo *YearTaxInfo `json:"year_tax_info,omitempty"`
	// 月度税务信息列表
	MonthTaxList []*MontTax `json:"month_tax_list,omitempty"`
}

// YearTaxInfo 综合所得汇算清缴信息
type YearTaxInfo struct {
	// 连续劳务年度个税
	ContinuousMonthPersonalTax string `json:"continuous_month_personal_tax,omitempty"`
	// 综合所得汇算清缴年度个税
	PersonalTax string `json:"personal_tax,omitempty"`
	// 年度扣除费用
	DeductCost string `json:"deduct_cost,omitempty"`
	// 个税税率
	PersonalTaxRate string `json:"personal_tax_rate,omitempty"`
	// 速算扣除数
	DeductTax string `json:"deduct_tax,omitempty"`
	// 税负率
	TotalTaxRate string `json:"total_tax_rate,omitempty"`
}

// MontTax 月度税务信息
type MontTax struct {
	// 月份
	Month int32 `json:"month,omitempty"`
	// 含增值税收入
	PreTaxAmount string `json:"pre_tax_amount,omitempty"`
	// 不含增值税收入
	ExcludingVatAmount string `json:"excluding_vat_amount,omitempty"`
	// 增值税
	ValueAddedTax string `json:"value_added_tax,omitempty"`
	// 附加税
	AdditionalTax string `json:"additional_tax,omitempty"`
	// 个税
	PersonalTax string `json:"personal_tax,omitempty"`
	// 个税税率
	PersonalTaxRate string `json:"personal_tax_rate,omitempty"`
	// 速算扣除数
	DeductTax string `json:"deduct_tax,omitempty"`
	// 税后金额
	PostTaxAmount string `json:"post_tax_amount,omitempty"`
	// 税负率
	TotalTaxRate string `json:"total_tax_rate,omitempty"`
}

// CalcTaxRequest 订单税费试算请求
type CalcTaxRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 证件号
	IDCard string `json:"id_card,omitempty"`
	// 订单金额
	Pay string `json:"pay,omitempty"`
}

// CalcTaxResponse 订单税费试算返回
type CalcTaxResponse struct {
	// 订单金额
	Pay string `json:"pay,omitempty"`
	// 税费总额
	Tax string `json:"tax,omitempty"`
	// 税后金额
	AfterTaxAmount string `json:"after_tax_amount,omitempty"`
	// 税费明细
	TaxDetail *CalcTaxDetail `json:"tax_detail,omitempty"`
}

// CalcTaxDetail 税费明细
type CalcTaxDetail struct {
	// 应纳个税
	PersonalTax string `json:"personal_tax,omitempty"`
	// 应纳增值税
	ValueAddedTax string `json:"value_added_tax,omitempty"`
	// 应纳附加税费
	AdditionalTax string `json:"additional_tax,omitempty"`
}
