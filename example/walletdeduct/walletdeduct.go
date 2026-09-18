package walletdeduct

import (
	"context"
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// CreateWalletDeductExample 申请钱包余额扣减
func CreateWalletDeductExample(client api.WalletDeductService) {
	req := &api.CreateWalletDeductRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		UserInfo: &api.WalletDeductUserInfo{
			RealName: "张三",
			IDCard:   "11010519491231002X",
			CardType: "idcard",
		},
		WalletID: "wallet_123456",
		OrderID:  "20200903001656212987",
		Scene:    "3",
		Amount:   "300.00",
		Remark:   "根据平台企业规则扣减",
		NotifyURL: "https://www.example.com/realtime/notify",
	}
	resp, err := client.CreateWalletDeduct(context.TODO(), req)
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

// QueryWalletDeductExample 查询钱包余额扣减申请结果
func QueryWalletDeductExample(client api.WalletDeductService) {
	req := &api.QueryWalletDeductRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		OrderID:  "20200903001656212987",
		Ref:      "176826728300002",
	}
	resp, err := client.QueryWalletDeduct(context.TODO(), req)
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

// CompleteWalletDeductExample 提交钱包余额扣减结果
func CompleteWalletDeductExample(client api.WalletDeductService) {
	req := &api.CompleteWalletDeductRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		OrderID:   "20200903001656212987",
		Ref:       "176826728300002",
		Status:    "1",
		TradeNo:   "202010150030000001",
		FinishedAt: "2020-10-15 00:30:00",
	}
	resp, err := client.CompleteWalletDeduct(context.TODO(), req)
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

// NotifyWalletDeductExample 钱包余额扣减申请结果回调通知
func NotifyWalletDeductExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/walletdeduct", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyWalletDeductRequest{}
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
	for _, example := range []func(api.WalletDeductService){
		CreateWalletDeductExample,
		QueryWalletDeductExample,
		CompleteWalletDeductExample,
	} {
		example(client)
	}
}
