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
		IDCard:    "填写自己的",
		CardNo:    "1111111111111111111111111",
		PhoneNo:   "13333333333",
		PayRemark: "银行卡打款",
		NotifyURL: "https://wwww.callback.com",
		Pay:       "99.99",
	}
	resp, err := client.CreateBankpayOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		fmt.Println(e.Code, e.Message)
		if e.Code == "2002" {
			/*
				订单号重复
				检查是否已存在该订单号
			*/

			// TODO 异常处理流程

		} else {
			/*
				下单异常
				其它错误详见文档
				可参考 e.message 中的错误信息
			*/

			fmt.Println(e.Code, e.Message)
		}
		return
	}

	// 没有 err 说明下单成功
	fmt.Println(resp)

}

// CreateAlipayOrder_Example 创建支付宝支付订单示例
func CreateAlipayOrder_Example(client api.Payment) {
	req := &api.CreateAlipayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "填写自己的",
		CardNo:    "1111111111111111111111111",
		PhoneNo:   "13333333333",
		PayRemark: "银行卡打款",
		NotifyURL: "https://wwww.callback.com",
		Pay:       "99.99",
	}
	resp, err := client.CreateAlipayOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		fmt.Println(e.Code, e.Message)
		if e.Code == "2002" {
			/*
				订单号重复
				检查是否已存在该订单号
			*/

			// TODO 异常处理流程

		} else {
			/*
				下单异常
				其它错误详见文档
				可参考 e.message 中的错误信息
			*/

			fmt.Println(e.Code, e.Message)
		}
		return
	}

	// 没有 err 说明下单成功
	fmt.Println(resp)

}

// CreateWxpayOrder_Example 创建支付宝支付订单示例
func CreateWxpayOrder_Example(client api.Payment) {
	req := &api.CreateWxpayOrderRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   time.Now().Format("20050102150405"),
		RealName:  "张三",
		IDCard:    "填写自己的",
		Openid:    "wx11111111111111111111111",
		PhoneNo:   "13333333333",
		PayRemark: "银行卡打款",
		NotifyURL: "https://wwww.callback.com",
		Pay:       "99.99",
	}
	resp, err := client.CreateWxpayOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		fmt.Println(e.Code, e.Message)
		if e.Code == "2002" {
			/*
				订单号重复
				检查是否已存在该订单号
			*/

			// TODO 异常处理流程

		} else {
			/*
				下单异常
				其它错误详见文档
				可参考 e.message 中的错误信息
			*/

			fmt.Println(e.Code, e.Message)
		}
		return
	}

	// 没有 err 说明下单成功
	fmt.Println(resp)

}

// GetOrder_Example 订单查询
func GetOrder_Example(client api.Payment) {
	/*
		注意:
			建议下单请求结束后30秒再进行订单查询操作，若查询接口返回code = 2018（订单不存在），则可以用原订单号重新下单, 切记不可换单号重试，否则有资损风险。
			响应码 code != 2018 时，均为异常状态，需要继续查单，不能进行下单重试，否则有资损风险。
	*/
	req := &api.GetOrderRequest{
		OrderID:  "416739461477437492",
		Channel:  "微信",
		DataType: "encryption",
	}
	resp, err := client.GetOrder(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}

		if e.Code == "2018" {
			// 说明订单不存在, 下单可以使用原单号重试
			// TODO
			// TODO:
			return
		}
		return
	}
	fmt.Println(resp)
}

// GetDealerVARechargeAccount_Example 查询平台企业 VA 账户信息
func GetDealerVARechargeAccount_Example(client api.Payment) {
	req := &api.GetDealerVARechargeAccountRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
	}
	resp, err := client.GetDealerVARechargeAccount(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
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
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
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
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
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
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// NotifyOrder_Example 订单回调样例
func NotifyOrder_Example() {
	// 可以采用其他框架
	http.HandleFunc("notify/order", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := r.PostForm.Get("data")
		timestamp := r.PostForm.Get("timestamp")
		mess := r.PostForm.Get("mess")
		sign := r.PostForm.Get("sign")
		signType := r.PostForm.Get("sign_type")
		dealerID := r.Header.Get("dealer-id")

		req := api.NotifyOrderRequest{}
		err := base.NotifyDecoder(dealerID, timestamp, data, mess, sign, signType, &req)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		// TODO: 处理 req信息 业务流程
	}))
}

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
	} {
		example(client)
	}
}
