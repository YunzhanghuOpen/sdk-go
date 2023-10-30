package payment

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// CreateBankpayOrder_Example 创建银行卡支付订单示例
func CreateBankpayOrder_Example(client api.Payment) {
	req := &api.CreateBankpayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "110121202202222222",
		CardNo:    "8888888888888888888",
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

// CreateAlipayOrder_Example 创建支付宝支付订单示例
func CreateAlipayOrder_Example(client api.Payment) {
	req := &api.CreateAlipayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "110121202202222222",
		CardNo:    "username@example.com",
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

// CreateWxpayOrder_Example 创建微信支付订单示例
func CreateWxpayOrder_Example(client api.Payment) {
	req := &api.CreateWxpayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "110121202202222222",
		Openid:    "wx11111111111111111111111",
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

// GetOrder_Example 订单查询
func GetOrder_Example(client api.Payment) {
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

// GetDealerVARechargeAccount_Example 查询平台企业汇款信息
func GetDealerVARechargeAccount_Example(client api.Payment) {
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

// ListAccount_Example 查询平台企业余额
func ListAccount_Example(client api.Payment) {
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

// GetEleReceiptFile_Example 查询电子回单
func GetEleReceiptFile_Example(client api.Payment) {
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

// CancelOrder_Example 取消待支付订单
func CancelOrder_Example(client api.Payment) {
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

// CreateBatchOrder_Example 批次下单
func CreateBatchOrder_Example(client api.Payment) {
	orderList := []*api.BatchOrderInfo{
		&api.BatchOrderInfo{OrderID: "202210220001", RealName: "张三", IDCard: "110121202202222221", Pay: "1", CardNo: "8888888888888888881"},
		&api.BatchOrderInfo{OrderID: "202210220002", RealName: "赵四", IDCard: "110121202202222222", Pay: "1", CardNo: "8888888888888888882"},
		&api.BatchOrderInfo{OrderID: "202210220003", RealName: "小白", IDCard: "110121202202222223", Pay: "1", CardNo: "8888888888888888883"},
	}
	req := &api.CreateBatchOrderRequest{
		BrokerID:   base.BrokerID,
		DealerID:   base.DealerID,
		BatchID:    "2022102200000001",
		Channel:    "银行卡",
		TotalCount: "3",
		TotalPay:   "3",
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

// ConfirmBatchOrder_Example 批次确认
func ConfirmBatchOrder_Example(client api.Payment) {
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

// CancelBatchOrder_Example 批次撤销
func CancelBatchOrder_Example(client api.Payment) {
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

// NotifyOrder_Example 订单回调样例
func NotifyOrder_Example() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/order", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")

			req := api.NotifyOrderRequestV2{}.Data
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
		CreateBankpayOrder_Example,
		CreateAlipayOrder_Example,
		CreateWxpayOrder_Example,
		GetOrder_Example,
		GetDealerVARechargeAccount_Example,
		ListAccount_Example,
		GetEleReceiptFile_Example,
		CancelOrder_Example,
		CreateBatchOrder_Example,
		ConfirmBatchOrder_Example,
	} {
		example(client)
	}
}
