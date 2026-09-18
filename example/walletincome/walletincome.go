package walletincome

import (
	"context"
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// CreateWalletIncomeExample 发起钱包余额入账
func CreateWalletIncomeExample(client api.WalletIncomeService) {
	req := &api.CreateWalletIncomeRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		UserInfo: &api.WalletIncomeUserInfo{
			RealName: "张三",
			IDCard:   "11010519491231002X",
			CardType: "idcard",
			PhoneNo:  "13800000000",
		},
		WalletID: "wallet_123456",
		PlatformInfo: &api.WalletIncomePlatformInfo{
			PlatformName: "xxx平台",
			UserID:       "123456",
			UserNickname: "张三",
		},
		OrderID:   "20200903001656212987",
		Amount:    "300.00",
		EarnedAt:  "2020-09-01 10:00:00",
		Remark:    "9月直播收入",
		NotifyURL: "https://www.example.com/income/notify",
	}
	resp, err := client.CreateWalletIncome(context.TODO(), req)
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

// QueryWalletIncomeExample 查询钱包余额入账结果
func QueryWalletIncomeExample(client api.WalletIncomeService) {
	req := &api.QueryWalletIncomeRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		OrderID:  "20200903001656212987",
		Ref:      "176826728300001",
	}
	resp, err := client.QueryWalletIncome(context.TODO(), req)
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

// CancelWalletIncomeExample 取消钱包收入计税订单
func CancelWalletIncomeExample(client api.WalletIncomeService) {
	req := &api.CancelWalletIncomeRequest{
		BrokerID:     base.BrokerID,
		DealerID:     base.DealerID,
		OrderID:      "20200903001656212987",
		Ref:          "176826728300001",
		CancelOrderID: "20200903001656212988",
	}
	resp, err := client.CancelWalletIncome(context.TODO(), req)
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

// RetryWalletIncomeExample 重试挂起的计税订单
func RetryWalletIncomeExample(client api.WalletIncomeService) {
	req := &api.RetryWalletIncomeRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		OrderID:  "20200903001656212987",
		Ref:      "176826728300001",
	}
	resp, err := client.RetryWalletIncome(context.TODO(), req)
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

// NotifyWalletIncomeExample 钱包余额入账结果回调通知
func NotifyWalletIncomeExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/walletincome", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyWalletIncomeRequest{}
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
	for _, example := range []func(api.WalletIncomeService){
		CreateWalletIncomeExample,
		QueryWalletIncomeExample,
		CancelWalletIncomeExample,
		RetryWalletIncomeExample,
	} {
		example(client)
	}
}
