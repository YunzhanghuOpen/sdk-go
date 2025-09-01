package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// CreateBankpayOrderExample 创建银行卡支付订单示例
func CreateBankpayOrderExample(client api.Payment) {
	req := &api.CreateBankpayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "11010519491231002X",
		CardNo:    "6228888888888888888",
		PhoneNo:   "188****8888",
		PayRemark: "银行卡支付",
		NotifyURL: "https://www.example.com",
		Pay:       "99.99",
	}
	resp, err := client.CreateBankpayOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		if e.Code == "2002" {
			// 订单号重复
			// 检查是否已存在该订单号
		} else {
			// 其它错误码详见接口文档附录中响应码列表
			fmt.Println(e.Code, e.Message)
		}
		return
	}
	// 操作成功
	fmt.Println(resp)

}

// CreateAlipayOrderExample 创建支付宝支付订单示例
func CreateAlipayOrderExample(client api.Payment) {
	req := &api.CreateAlipayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "11010519491231002X",
		CardNo:    "188****8888",
		PhoneNo:   "188****8888",
		PayRemark: "支付宝支付",
		NotifyURL: "https://www.example.com",
		Pay:       "99.99",
	}
	resp, err := client.CreateAlipayOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		if e.Code == "2002" {
			// 订单号重复
			// 检查是否已存在该订单号
		} else {
			// 其它错误码详见接口文档附录中响应码列表
			fmt.Println(e.Code, e.Message)
		}
		return
	}
	// 操作成功
	fmt.Println(resp)

}

// CreateWxpayOrderExample 创建微信支付订单示例
func CreateWxpayOrderExample(client api.Payment) {
	req := &api.CreateWxpayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "11010519491231002X",
		Openid:    "o4GgauInH_RCEdvrrNGrntXDuXXX",
		PhoneNo:   "188****8888",
		PayRemark: "微信支付",
		NotifyURL: "https://www.example.com",
		Pay:       "99.99",
	}
	resp, err := client.CreateWxpayOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		if e.Code == "2002" {
			// 订单号重复
			// 检查是否已存在该订单号
		} else {
			// 其它错误码详见接口文档附录中响应码列表
			fmt.Println(e.Code, e.Message)
		}
		return
	}
	// 操作成功
	fmt.Println(resp)

}

// GetOrderExample 订单查询
func GetOrderExample(client api.Payment) {
	//  注意:
	// 	建议下单请求结束后30秒再进行订单查询操作。
	// 	若响应码code = 2018，表示订单不存在，则可使用原订单号重新下单，切记不可更换其他单号重试，否则将存在资损风险。
	// 	若返回响应码 code != 2018 ，均表示异常状态，则需继续查单，不能进行下单重试，否则将存在有资损风险。
	req := &api.GetOrderRequest{
		OrderID:  "416739461477437492",
		Channel:  "微信",
		DataType: "encryption",
	}
	resp, err := client.GetOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		if e.Code == "2018" {
			// 订单不存在
			return
		}
		return
	}
	// 操作成功
	fmt.Println(resp)
}

// GetDealerVARechargeAccountExample 查询平台企业汇款信息
func GetDealerVARechargeAccountExample(client api.Payment) {
	req := &api.GetDealerVARechargeAccountRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
	}
	resp, err := client.GetDealerVARechargeAccount(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// ListAccountExample 查询平台企业余额
func ListAccountExample(client api.Payment) {
	req := &api.ListAccountRequest{
		DealerID: base.DealerID,
	}
	resp, err := client.ListAccount(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// GetEleReceiptFileExample 查询电子回单
func GetEleReceiptFileExample(client api.Payment) {
	req := &api.GetEleReceiptFileRequest{
		OrderID: "123456",
		Ref:     "",
	}
	resp, err := client.GetEleReceiptFile(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// CancelOrderExample 取消待支付订单
func CancelOrderExample(client api.Payment) {
	req := &api.CancelOrderRequest{
		DealerID: base.DealerID,
		OrderID:  "123456",
		Channel:  "",
		Ref:      "",
	}
	resp, err := client.CancelOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// RetryOrderExample 重试挂起状态订单
func RetryOrderExample(client api.Payment) {
	req := &api.RetryOrderRequest{
		DealerID: base.DealerID,
		OrderID:  "202009010016562012987",
		Ref:      "176826728298982",
		Channel:  "bankpay",
	}
	resp, err := client.RetryOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// CheckUserAmountExample 用户结算金额校验
func CheckUserAmountExample(client api.Payment) {
	req := &api.CheckUserAmountRequest{
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "11010519491231002X",
		Amount:   "10000.00",
	}
	resp, err := client.CheckUserAmount(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// CreateBatchOrderExample 批次下单
func CreateBatchOrderExample(client api.Payment) {
	orderList := []*api.BatchOrderInfo{
		&api.BatchOrderInfo{OrderID: "202210220001", RealName: "张三", IDCard: "440524188001010014", Pay: "1", CardNo: "6228888888888888887", PhoneNo: "188****8888"},
		&api.BatchOrderInfo{OrderID: "202210220002", RealName: "李四", IDCard: "11010519491231002X", Pay: "1", CardNo: "6228888888888888888", PhoneNo: "188****8888"},
	}
	req := &api.CreateBatchOrderRequest{
		BrokerID:   base.BrokerID,
		DealerID:   base.DealerID,
		BatchID:    "2022102200000001",
		Channel:    "银行卡",
		TotalCount: "2",
		TotalPay:   "2",
		OrderList:  orderList,
	}
	resp, err := client.CreateBatchOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		if e.Code == "2001" {
			// 批次号重复
			// 检查是否已上传过该批次号
		} else {
			// 其它错误码详见接口文档附录中响应码列表
			fmt.Println(e.Code, e.Message)
		}
		return
	}
	// 操作成功
	fmt.Println(resp)
}

// ConfirmBatchOrderExample 批次确认
func ConfirmBatchOrderExample(client api.Payment) {
	req := &api.ConfirmBatchOrderRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		BatchID:  "2022102200000001",
		Channel:  "银行卡",
	}
	resp, err := client.ConfirmBatchOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		return
	}
	// 操作成功
	fmt.Println(resp)
}

// QueryBatchOrderExample 查询批次订单信息
func QueryBatchOrderExample(client api.Payment) {
	req := &api.QueryBatchOrderRequest{
		BatchID:  "2022102200000001",
		DealerID: base.DealerID,
	}
	resp, err := client.QueryBatchOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		return
	}
	// 操作成功
	fmt.Println(resp)
}

// CancelBatchOrderExample 批次撤销
func CancelBatchOrderExample(client api.Payment) {
	req := &api.CancelBatchOrderRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		BatchID:  "2022102200000001",
	}
	resp, err := client.CancelBatchOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		return
	}
	// 操作成功
	fmt.Println(resp)
}

// GetOrderLxlwExample 查询劳务模式单笔订单信息
func GetOrderLxlwExample(client api.Payment) {
	req := &api.GetOrderLxlwRequest{
		OrderID:  "140490814193549",
		Channel:  "银行卡",
		DataType: "",
	}
	resp, err := client.GetOrderLxlw(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
		return
	}
	// 操作成功
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// NotifyOrderExample 订单回调样例
func NotifyOrderExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/order", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")

			req := api.NotifyOrderRequestV2{}
			err := base.NotifyDecoder(mess, timestamp, data, sign, signType, &req)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}
	}))
}

// NotifyOrderLxlwExample 连续劳务订单回调样例
func NotifyOrderLxlwExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/order", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyOrderLxlwRequest{}
			err := base.NotifyDecoder(mess, timestamp, data, sign, signType, &req)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}
	}))
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.Payment){
		CreateBankpayOrderExample,
		CreateAlipayOrderExample,
		CreateWxpayOrderExample,
		GetOrderExample,
		GetDealerVARechargeAccountExample,
		ListAccountExample,
		GetEleReceiptFileExample,
		CancelOrderExample,
		RetryOrderExample,
		CheckUserAmountExample,
		CreateBatchOrderExample,
		ConfirmBatchOrderExample,
		QueryBatchOrderExample,
		GetOrderLxlwExample,
	} {
		example(client)
	}
}
