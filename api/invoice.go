package api

import (
	"context"
)

// Invoice Invoice 发票接口
type Invoice interface {
	// GetInvoiceStat 查询平台企业已开具和待开具发票金额
	GetInvoiceStat(context.Context, *GetInvoiceStatRequest) (*GetInvoiceStatResponse, error)
	// GetInvoiceAmount 查询可开票额度和开票信息
	GetInvoiceAmount(context.Context, *GetInvoiceAmountRequest) (*GetInvoiceAmountResponse, error)
	// ApplyInvoice 开票申请
	ApplyInvoice(context.Context, *ApplyInvoiceRequest) (*ApplyInvoiceResponse, error)
	// GetInvoiceStatus 查询开票申请状态
	GetInvoiceStatus(context.Context, *GetInvoiceStatusRequest) (*GetInvoiceStatusResponse, error)
	// GetInvoiceFile 下载发票 PDF
	GetInvoiceFile(context.Context, *GetInvoiceFileRequest) (*GetInvoiceFileResponse, error)
	// SendReminderEmail 发送发票扫描件压缩包下载链接邮件
	SendReminderEmail(context.Context, *SendReminderEmailRequest) (*SendReminderEmailResponse, error)
}

// invoiceImpl Invoice 接口实现
type invoiceImpl struct {
	cc Invoker
}

// NewInvoice 创建客户端
func NewInvoice(cc Invoker) Invoice {
	return &invoiceImpl{cc}
}

// GetInvoiceStat 查询平台企业已开具和待开具发票金额
func (c *invoiceImpl) GetInvoiceStat(ctx context.Context, in *GetInvoiceStatRequest) (*GetInvoiceStatResponse, error) {
	out := new(GetInvoiceStatResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/payment/v1/invoice-stat", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetInvoiceAmount 查询可开票额度和开票信息
func (c *invoiceImpl) GetInvoiceAmount(ctx context.Context, in *GetInvoiceAmountRequest) (*GetInvoiceAmountResponse, error) {
	out := new(GetInvoiceAmountResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/invoice/v2/invoice-amount", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplyInvoice 开票申请
func (c *invoiceImpl) ApplyInvoice(ctx context.Context, in *ApplyInvoiceRequest) (*ApplyInvoiceResponse, error) {
	out := new(ApplyInvoiceResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/invoice/v2/apply", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetInvoiceStatus 查询开票申请状态
func (c *invoiceImpl) GetInvoiceStatus(ctx context.Context, in *GetInvoiceStatusRequest) (*GetInvoiceStatusResponse, error) {
	out := new(GetInvoiceStatusResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/invoice/v2/invoice/invoice-status", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetInvoiceFile 下载发票 PDF
func (c *invoiceImpl) GetInvoiceFile(ctx context.Context, in *GetInvoiceFileRequest) (*GetInvoiceFileResponse, error) {
	out := new(GetInvoiceFileResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/invoice/v2/invoice/invoice-pdf", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendReminderEmail 发送发票扫描件压缩包下载链接邮件
func (c *invoiceImpl) SendReminderEmail(ctx context.Context, in *SendReminderEmailRequest) (*SendReminderEmailResponse, error) {
	out := new(SendReminderEmailResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/invoice/v2/invoice/reminder/email", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetInvoiceStatRequest 查询平台企业已开具和待开具发票金额请求
type GetInvoiceStatRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 商户 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 查询年份
	Year int32 `json:"year,omitempty"`
}

// GetInvoiceStatResponse 查询平台企业已开具和待开具发票金额响应
type GetInvoiceStatResponse struct {
	// 商户 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 已开具发票金额
	Invoiced string `json:"invoiced,omitempty"`
	// 待开具发票金额
	NotInvoiced string `json:"not_invoiced,omitempty"`
}

// GetInvoiceAmountRequest 查询可开票额度和开票信息请求
type GetInvoiceAmountRequest struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 商户 ID
	DealerID string `json:"dealer_id,omitempty"`
}

// GetInvoiceAmountResponse 查询可开票额度和开票信息响应
type GetInvoiceAmountResponse struct {
	// 可开票额度
	Amount string `json:"amount,omitempty"`
	// 系统支持的开户行及账号
	BankNameAccount []*BankNameAccount `json:"bank_name_account,omitempty"`
	// 系统支持的货物或应税劳务、服务名称
	GoodsServicesName []*GoodsServicesName `json:"goods_services_name,omitempty"`
}

// ApplyInvoiceRequest 开票申请请求
type ApplyInvoiceRequest struct {
	// 发票申请编号
	InvoiceApplyID string `json:"invoice_apply_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 商户 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 申请开票金额
	Amount string `json:"amount,omitempty"`
	// 发票类型
	InvoiceType string `json:"invoice_type,omitempty"`
	// 开户行及账号
	BankNameAccount string `json:"bank_name_account,omitempty"`
	// 货物或应税劳务、服务名称
	GoodsServicesName string `json:"goods_services_name,omitempty"`
	// 发票备注
	Remark string `json:"remark,omitempty"`
}

// ApplyInvoiceResponse 开票申请响应
type ApplyInvoiceResponse struct {
	// 发票申请单 ID
	ApplicationID string `json:"application_id,omitempty"`
	// 发票张数
	Count int64 `json:"count,omitempty"`
}

// GetInvoiceStatusRequest 查询开票申请状态请求
type GetInvoiceStatusRequest struct {
	// 发票申请编号
	InvoiceApplyID string `json:"invoice_apply_id,omitempty"`
	// 发票申请单 ID
	ApplicationID string `json:"application_id,omitempty"`
}

// GetInvoiceStatusResponse 查询开票申请状态响应
type GetInvoiceStatusResponse struct {
	// 申请结果
	Status string `json:"status,omitempty"`
	// 发票张数
	Count int64 `json:"count,omitempty"`
	// 价税合计
	PriceTaxAmount string `json:"price_tax_amount,omitempty"`
	// 不含税金额
	PriceAmount string `json:"price_amount,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty"`
	// 发票类型
	InvoiceType string `json:"invoice_type,omitempty"`
	// 购方名称
	CustomerName string `json:"customer_name,omitempty"`
	// 纳税人识别号
	CustomerTaxNum string `json:"customer_tax_num,omitempty"`
	// 购方地址、电话
	CustomerAddressTel string `json:"customer_address_tel,omitempty"`
	// 开户行及账号
	BankNameAccount string `json:"bank_name_account,omitempty"`
	// 货物或应税劳务、服务名称
	GoodsServicesName string `json:"goods_services_name,omitempty"`
	// 发票备注
	Remark string `json:"remark,omitempty"`
	// 邮寄类型
	PostType string `json:"post_type,omitempty"`
	// 快递单号
	WaybillNumber []string `json:"waybill_number,omitempty"`
}

// BankNameAccount 系统支持的开户行及账号
type BankNameAccount struct {
	// 开户行及账号
	Item string `json:"item,omitempty"`
	// 是否为默认值
	Default bool `json:"default,omitempty"`
}

// GoodsServicesName 统支持的货物或应税劳务、服务名称
type GoodsServicesName struct {
	// 货物或应税劳务、服务名称
	Item string `json:"item,omitempty"`
	// 是否为默认值
	Default bool `json:"default,omitempty"`
}

// GetInvoiceFileRequest 下载发票 PDF 请求
type GetInvoiceFileRequest struct {
	// 发票申请编号
	InvoiceApplyID string `json:"invoice_apply_id,omitempty"`
	// 发票申请单 ID
	ApplicationID string `json:"application_id,omitempty"`
}

// GetInvoiceFileResponse 下载发票 PDF 响应
type GetInvoiceFileResponse struct {
	// 下载地址
	URL string `json:"url,omitempty"`
	// 文件名称
	Name string `json:"name,omitempty"`
}

// SendReminderEmailRequest 发送发票扫描件压缩包下载链接邮件请求
type SendReminderEmailRequest struct {
	// 发票申请编号
	InvoiceApplyID string `json:"invoice_apply_id,omitempty"`
	// 发票申请单 ID
	ApplicationID string `json:"application_id,omitempty"`
}

// SendReminderEmailResponse 发送发票扫描件压缩包下载链接邮件响应
type SendReminderEmailResponse struct {
}

// NotifyInvoiceDoneRequest 开票完成通知
type NotifyInvoiceDoneRequest struct {
	// 发票申请编号
	ApplicationID string `json:"application_id,omitempty"`
	// 发票申请编号
	InvoiceApplyID string `json:"invoice_apply_id,omitempty"`
	// 申请结果
	Status string `json:"status,omitempty"`
	// 发票张数
	Count int64 `json:"count,omitempty"`
	// 价税合计
	PriceTaxAmount string `json:"price_tax_amount,omitempty"`
	// 不含税金额
	PriceAmount string `json:"price_amount,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty"`
	// 发票类型
	InvoiceType string `json:"invoice_type,omitempty"`
	// 购方名称
	CustomerName string `json:"customer_name,omitempty"`
	// 纳税人识别号
	CustomerTaxNum string `json:"customer_tax_num,omitempty"`
	// 购方地址、电话
	CustomerAddressTel string `json:"customer_address_tel,omitempty"`
	// 开户行及账号
	BankNameAccount string `json:"bank_name_account,omitempty"`
	// 货物或应税劳务、服务名称
	GoodsServicesName string `json:"goods_services_name,omitempty"`
	// 发票备注
	Remark string `json:"remark,omitempty"`
	// 邮寄类型
	PostType string `json:"post_type,omitempty"`
	// 快递单号
	WaybillNumber []string `json:"waybill_number,omitempty"`
	// 驳回原因
	RejectReason string `json:"reject_reason,omitempty"`
}
