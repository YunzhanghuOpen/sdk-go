package walletwithdraw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// CreateWalletWithdrawExample 发起钱包余额提现
func CreateWalletWithdrawExample(client api.WalletWithdrawService) {
	req := &api.CreateWalletWithdrawRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		UserInfo: &api.WalletWithdrawUserInfo{
			RealName: "张三",
			IDCard:   "11010519491231002X",
			CardType: "idcard",
		},
		WalletID:  "wallet_123456",
		OrderID:   "20200903001656212989",
		WxAppID:   "wx1234567890abcdef",
		Amount:    "300.00",
		Channel:   "wxpay",
		Account:   "13800000000",
		Remark:    "10月提现",
		NotifyURL: "https://www.example.com/withdraw/notify",
	}
	resp, err := client.CreateWalletWithdraw(context.TODO(), req)
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

// QueryWalletWithdrawExample 查询钱包余额提现结果
func QueryWalletWithdrawExample(client api.WalletWithdrawService) {
	req := &api.QueryWalletWithdrawRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		Channel:  "wxpay",
		OrderID:  "20200903001656212989",
		Ref:      "176826728300003",
	}
	resp, err := client.QueryWalletWithdraw(context.TODO(), req)
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

// CancelWalletWithdrawExample 取消挂起的钱包余额提现订单
func CancelWalletWithdrawExample(client api.WalletWithdrawService) {
	req := &api.CancelWalletWithdrawRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		OrderID:  "20200903001656212989",
		Ref:      "176826728300003",
	}
	resp, err := client.CancelWalletWithdraw(context.TODO(), req)
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

// RetryWalletWithdrawExample 重试挂起的钱包余额提现订单
func RetryWalletWithdrawExample(client api.WalletWithdrawService) {
	req := &api.RetryWalletWithdrawRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		OrderID:  "20200903001656212989",
		Ref:      "176826728300003",
	}
	resp, err := client.RetryWalletWithdraw(context.TODO(), req)
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

// GetWalletWithdrawReceiptFileExample 查询钱包余额提现电子回单
func GetWalletWithdrawReceiptFileExample(client api.WalletWithdrawService) {
	req := &api.GetWalletWithdrawReceiptFileRequest{
		BrokerID:    base.BrokerID,
		DealerID:    base.DealerID,
		OrderID:     "20200903001656212989",
		Ref:         "176826728300003",
		ReceiptType: "付款回单",
	}
	resp, err := client.GetWalletWithdrawReceiptFile(context.TODO(), req)
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

// NotifyWalletWithdrawExample 通知钱包余额提现结果回调通知
func NotifyWalletWithdrawExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/walletwithdraw", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyWalletWithdrawRequest{}
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
	for _, example := range []func(api.WalletWithdrawService){
		CreateWalletWithdrawExample,
		QueryWalletWithdrawExample,
		CancelWalletWithdrawExample,
		RetryWalletWithdrawExample,
		GetWalletWithdrawReceiptFileExample,
	} {
		example(client)
	}
}
