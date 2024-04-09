package custom

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// 自定义请求结构体（注意需要补充 json-tag）
type customRequest struct {
	// 平台企业订单号
	OrderID string `json:"order_id"`
	// 平台企业 ID
	DealerID string `json:"dealer_id"`
	// 综合服务主体 ID
	BrokerID string `json:"broker_id"`
	// 姓名
	RealName string `json:"real_name"`
	// 银行卡号
	CardNo string `json:"card_no"`
	// 身份证号码
	IDCard string `json:"id_card"`
	// 手机号
	PhoneNo string `json:"phone_no"`
	// 订单金额
	Pay string `json:"pay"`
	// 订单备注
	PayRemark string `json:"pay_remark"`
	// 回调地址
	NotifyURL string `json:"notify_url"`
}

// CustomServiceExample 通用请求函数
func CustomServiceExample(client api.CustomService) {
	// 构建业务请求参数
	req := &customRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   "202009010016562012987",
		RealName:  "张三",
		IDCard:    "11010519491231002X",
		CardNo:    "6228888888888888888",
		PhoneNo:   "188****8888",
		PayRemark: "",
		NotifyURL: "https://www.example.com",
		Pay:       "100.00",
	}

	// 请求方式
	method := "POST"
	// 接口地址，具体接口地址参见云账户接口文档
	url := "/api/payment/v1/order-bankpay"

	// 发起请求
	resp, err := client.DoRequest(context.TODO(), method, url, req)
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

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(service api.CustomService){
		CustomServiceExample,
	} {
		example(client)
	}
}
