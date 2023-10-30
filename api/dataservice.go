package api

import (
	"context"
)

// DataService 对账文件获取
type DataService interface {
	// ListDailyOrder 查询日订单数据
	ListDailyOrder(context.Context, *ListDailyOrderRequest) (*ListDailyOrderResponse, error)
	// GetDailyOrderFile 查询日订单文件
	GetDailyOrderFile(context.Context, *GetDailyOrderFileRequest) (*GetDailyOrderFileResponse, error)
	// GetDailyOrderFileV2 查询日订单文件（支付和退款订单）
	GetDailyOrderFileV2(context.Context, *GetDailyOrderFileV2Request) (*GetDailyOrderFileV2Response, error)
	// ListDailyBill 查询日流水数据
	ListDailyBill(context.Context, *ListDailyBillRequest) (*ListDailyBillResponse, error)
	// GetDailyBillFileV2 查询日流水文件
	GetDailyBillFileV2(context.Context, *GetDailyBillFileV2Request) (*GetDailyBillFileV2Response, error)
	// ListDealerRechargeRecordV2 查询平台企业预付业务服务费记录
	ListDealerRechargeRecordV2(context.Context, *ListDealerRechargeRecordV2Request) (*ListDealerRechargeRecordV2Response, error)
	// ListBalanceDailyStatement 查询余额日账单数据
	ListBalanceDailyStatement(context.Context, *ListBalanceDailyStatementRequest) (*ListBalanceDailyStatementResponse, error)
}

// dataServiceImpl DataService 接口实现
type dataServiceImpl struct {
	cc Invoker
}

// NewDataService 创建客户端
func NewDataService(cc Invoker) DataService {
	return &dataServiceImpl{cc}
}

// ListDailyOrder 查询日订单数据
func (c *dataServiceImpl) ListDailyOrder(ctx context.Context, in *ListDailyOrderRequest) (*ListDailyOrderResponse, error) {
	out := new(ListDailyOrderResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/dataservice/v1/orders", in.DataType == "encryption", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDailyOrderFile 查询日订单文件
func (c *dataServiceImpl) GetDailyOrderFile(ctx context.Context, in *GetDailyOrderFileRequest) (*GetDailyOrderFileResponse, error) {
	out := new(GetDailyOrderFileResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/dataservice/v1/order/downloadurl", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDailyOrderFileV2 查询日订单文件（支付和退款订单）
func (c *dataServiceImpl) GetDailyOrderFileV2(ctx context.Context, in *GetDailyOrderFileV2Request) (*GetDailyOrderFileV2Response, error) {
	out := new(GetDailyOrderFileV2Response)
	err := c.cc.Invoke(ctx, "GET", "/api/dataservice/v1/order/day/url", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListDailyBill 查询日流水数据
func (c *dataServiceImpl) ListDailyBill(ctx context.Context, in *ListDailyBillRequest) (*ListDailyBillResponse, error) {
	out := new(ListDailyBillResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/dataservice/v1/bills", in.DataType == "encryption", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDailyBillFileV2 查询日流水文件
func (c *dataServiceImpl) GetDailyBillFileV2(ctx context.Context, in *GetDailyBillFileV2Request) (*GetDailyBillFileV2Response, error) {
	out := new(GetDailyBillFileV2Response)
	err := c.cc.Invoke(ctx, "GET", "/api/dataservice/v2/bill/downloadurl", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListDealerRechargeRecordV2 查询平台企业预付业务服务费记录
func (c *dataServiceImpl) ListDealerRechargeRecordV2(ctx context.Context, in *ListDealerRechargeRecordV2Request) (*ListDealerRechargeRecordV2Response, error) {
	out := new(ListDealerRechargeRecordV2Response)
	err := c.cc.Invoke(ctx, "GET", "/api/dataservice/v2/recharge-record", false, in, &out.Data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListBalanceDailyStatement 查询余额日账单数据
func (c *dataServiceImpl) ListBalanceDailyStatement(ctx context.Context, in *ListBalanceDailyStatementRequest) (*ListBalanceDailyStatementResponse, error) {
	out := new(ListBalanceDailyStatementResponse)
	err := c.cc.Invoke(ctx, "GET", "/api/dataservice/v1/statements-daily", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDailyOrderFileRequest 查询日订单文件请求
type GetDailyOrderFileRequest struct {
	// 订单查询日期, 格式：yyyy-MM-dd
	OrderDate string `json:"order_date,omitempty"`
}

// GetDailyOrderFileResponse 查询日订单文件返回
type GetDailyOrderFileResponse struct {
	// 下载地址
	OrderDownloadURL string `json:"order_download_url,omitempty"`
}

// GetDailyBillFileV2Request 查询日流水文件请求
type GetDailyBillFileV2Request struct {
	// 所需获取的日流水日期，格式：yyyy-MM-dd
	BillDate string `json:"bill_date,omitempty"`
}

// GetDailyBillFileV2Response 查询日流水文件返回
type GetDailyBillFileV2Response struct {
	// 下载地址
	BillDownloadURL string `json:"bill_download_url,omitempty"`
}

// ListDealerRechargeRecordV2Request 平台企业预付业务服务费记录请求
type ListDealerRechargeRecordV2Request struct {
	// 开始时间，格式：yyyy-MM-dd
	BeginAt string `json:"begin_at,omitempty"`
	// 结束时间，格式：yyyy-MM-dd
	EndAt string `json:"end_at,omitempty"`
}

// ListDealerRechargeRecordV2Response 平台企业预付业务服务费记录返回
type ListDealerRechargeRecordV2Response struct {
	// 预付业务服务费记录
	Data []*RechargeRecordInfo `json:"data,omitempty"`
}

// RechargeRecordInfo 预付业务服务费记录信息
type RechargeRecordInfo struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 预付业务服务费记录 ID
	RechargeID int64 `json:"recharge_id,omitempty"`
	// 预付业务服务费
	Amount float64 `json:"amount,omitempty"`
	// 实际到账金额
	ActualAmount float64 `json:"actual_amount,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 资金用途
	RechargeChannel string `json:"recharge_channel,omitempty"`
	// 预付业务服务费备注
	Remark string `json:"remark,omitempty"`
	// 平台企业付款银行账号
	RechargeAccountNo string `json:"recharge_account_no,omitempty"`
}

// ListDailyOrderRequest 查询日订单数据请求
type ListDailyOrderRequest struct {
	// 订单查询日期, 格式：yyyy-MM-dd格式：yyyy-MM-dd
	OrderDate string `json:"order_date,omitempty"`
	// 偏移量
	Offset int32 `json:"offset,omitempty"`
	// 长度
	Length int32 `json:"length,omitempty"`
	// 支付路径名，银行卡（默认）、支付宝、微信
	Channel string `json:"channel,omitempty"`
	// 如果为 encryption，则对返回的 data 进行加密
	DataType string `json:"data_type,omitempty"`
}

// ListDailyOrderResponse 查询日订单数据返回
type ListDailyOrderResponse struct {
	// 总数目
	TotalNum int32 `json:"total_num,omitempty"`
	// 条目信息
	List []*DealerOrderInfo `json:"list,omitempty"`
}

// DealerOrderInfo 平台企业支付订单信息
type DealerOrderInfo struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 订单流水号
	Ref string `json:"ref,omitempty"`
	// 批次ID
	BatchID string `json:"batch_id,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty"`
	// 收款账号
	CardNo string `json:"card_no,omitempty"`
	// 综合服务主体订单金额
	BrokerAmount string `json:"broker_amount,omitempty"`
	// 综合服务主体加成服务费
	BrokerFee string `json:"broker_fee,omitempty"`
	// 支付路径流水号
	Bill string `json:"bill,omitempty"`
	// 订单状态
	Status string `json:"status,omitempty"`
	// 订单状态码描述
	StatusMessage string `json:"status_message,omitempty"`
	// 订单详情
	StatusDetail string `json:"status_detail,omitempty"`
	// 订单详细状态码描述
	StatusDetailMessage string `json:"status_detail_message,omitempty"`
	// 短周期授信账单号
	StatementID string `json:"statement_id,omitempty"`
	// 服务费账单号
	FeeStatementID string `json:"fee_statement_id,omitempty"`
	// 余额账单号
	BalStatementID string `json:"bal_statement_id,omitempty"`
	// 支付路径
	Channel string `json:"channel,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 完成时间
	FinishedTime string `json:"finished_time,omitempty"`
}

// ListDailyBillRequest 查询日流水数据请求
type ListDailyBillRequest struct {
	// 流水查询日期
	BillDate string `json:"bill_date,omitempty"` // 流水查询日期
	// 偏移量
	Offset int32 `json:"offset,omitempty"`
	// 长度
	Length int32 `json:"length,omitempty"`
	// 如果为 encryption，则对返回的 data 进行加密
	DataType string `json:"data_type,omitempty"`
}

// ListDailyBillResponse 查询日流水数据返回
type ListDailyBillResponse struct {
	// 总条数
	TotalNum int32 `json:"total_num,omitempty"`
	// 条目信息
	List []*DealerBillInfo `json:"list,omitempty"`
}

// DealerBillInfo 流水详情
type DealerBillInfo struct {
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业订单号
	OrderID string `json:"order_id,omitempty"`
	// 资金流水号
	Ref string `json:"ref,omitempty"`
	// 综合服务主体名称
	BrokerProductName string `json:"broker_product_name,omitempty"`
	// 平台企业名称
	DealerProductName string `json:"dealer_product_name,omitempty"`
	// 业务订单流水号
	BizRef string `json:"biz_ref,omitempty"`
	// 账户类型
	AcctType string `json:"acct_type,omitempty"`
	// 入账金额
	Amount string `json:"amount,omitempty"`
	// 账户余额
	Balance string `json:"balance,omitempty"`
	// 业务分类
	BusinessCategory string `json:"business_category,omitempty"`
	// 业务类型
	BusinessType string `json:"business_type,omitempty"`
	// 收支类型
	ConsumptionType string `json:"consumption_type,omitempty"`
	// 入账时间
	CreatedAt string `json:"created_at,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
}

// GetDailyOrderFileV2Request 查询日订单文件（支付和退款订单）请求
type GetDailyOrderFileV2Request struct {
	// 订单查询日期, 格式：yyyy-MM-dd
	OrderDate string `json:"order_date,omitempty"`
}

// GetDailyOrderFileV2Response 查询日订单文件（支付和退款订单）返回
type GetDailyOrderFileV2Response struct {
	// 下载地址
	URL string `json:"url,omitempty"`
}

// ListBalanceDailyStatementRequest 查询余额日账单数据请求
type ListBalanceDailyStatementRequest struct {
	// 账单查询日期 格式：yyyy-MM-dd
	StatementDate string `json:"statement_date,omitempty"`
}

// ListBalanceDailyStatementResponse 查询余额日账单数据返回
type ListBalanceDailyStatementResponse struct {
	// 条目信息
	List []*StatementDetail `json:"list,omitempty"`
}

// StatementDetail 余额账单信息详情
type StatementDetail struct {
	// 账单 ID
	StatementID string `json:"statement_id,omitempty"`
	// 账单日期
	StatementDate string `json:"statement_date,omitempty"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id,omitempty"`
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 综合服务主体名称
	BrokerProductName string `json:"broker_product_name,omitempty"`
	// 平台企业名称
	DealerProductName string `json:"dealer_product_name,omitempty"`
	// 业务类型
	BizType string `json:"biz_type,omitempty"`
	// 账单金额
	TotalMoney string `json:"total_money,omitempty"`
	// 订单金额
	Amount string `json:"amount,omitempty"`
	// 退汇金额
	ReexAmount string `json:"reex_amount,omitempty"`
	// 加成服务费金额
	FeeAmount string `json:"fee_amount,omitempty"`
	// 加成服务费抵扣金额
	DeductRebateFeeAmount string `json:"deduct_rebate_fee_amount,omitempty"`
	// 冲补金额
	MoneyAdjust string `json:"money_adjust,omitempty"`
	// 账单状态
	Status string `json:"status,omitempty"`
	// 开票状态
	InvoiceStatus string `json:"invoice_status,omitempty"`
	// 项目 ID
	ProjectID string `json:"project_id,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty"`
}
