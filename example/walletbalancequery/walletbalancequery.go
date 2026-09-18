package walletbalancequery

import (
	"context"
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// QueryWalletBalanceExample 查询钱包余额
func QueryWalletBalanceExample(client api.WalletBalanceQueryService) {
	req := &api.QueryWalletBalanceRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		UserInfo: &api.WalletBalanceQueryUserInfo{
			RealName: "张三",
			IDCard:   "11010519491231002X",
			CardType: "idcard",
		},
		WalletID: "wallet_123456",
	}
	resp, err := client.QueryWalletBalance(context.TODO(), req)
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

// NotifyWalletBalanceChangeExample 钱包余额变更结果回调通知
func NotifyWalletBalanceChangeExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/walletbalance", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyWalletBalanceChangeRequest{}
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
	for _, example := range []func(api.WalletBalanceQueryService){
		QueryWalletBalanceExample,
	} {
		example(client)
	}
}
