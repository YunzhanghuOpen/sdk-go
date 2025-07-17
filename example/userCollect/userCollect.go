package userCollect

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// GetUserCollectPhoneStatusExample 查询手机号码绑定状态
func GetUserCollectPhoneStatusExample(client api.UserCollectService) {
	req := &api.GetUserCollectPhoneStatusRequest{
		DealerID:        base.DealerID,
		BrokerID:        base.BrokerID,
		DealerUserID:    "userId1234567890",
		RealName:        "张三",
		IDCard:          "11010519491231002X",
		CertificateType: 0,
	}
	resp, err := client.GetUserCollectPhoneStatus(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		errorData, err := json.Marshal(e)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(errorData))
	}
	// 操作成功
	data, err := json.Marshal(resp)
	fmt.Println(string(data))
}

// GetUserCollectPhoneURLExample 获取收集手机号码页面
func GetUserCollectPhoneURLExample(client api.UserCollectService) {
	req := &api.GetUserCollectPhoneURLRequest{
		Token:       "testToken",
		Color:       "",
		URL:         "https://www.example.com",
		RedirectURL: "",
	}
	resp, err := client.GetUserCollectPhoneURL(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		errorData, err := json.Marshal(e)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(errorData))
	}
	// 操作成功
	data, err := json.Marshal(resp)
	fmt.Println(string(data))
}

// NotifyUploadUserSignExample 签约成功状态回调通知
func NotifyUserCollectPhoneExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/userCollectPhone", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyUserCollectPhoneRequest{}
			err := base.NotifyDecoder(mess, timestamp, data, sign, signType, &req)
			if err != nil {
				fmt.Println(err)
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
	for _, example := range []func(api.UserCollectService){
		GetUserCollectPhoneStatusExample,
		GetUserCollectPhoneURLExample,
	} {
		example(client)
	}
}
