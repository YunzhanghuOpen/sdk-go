package uploadusersign

import (
	"context"
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// UploadUserSignExample 用户签约信息上传
func UploadUserSignExample(client api.UploadUserSignService) {
	req := &api.UploadUserSignRequest{
		BrokerID:  base.BrokerID,
		DealerID:  base.DealerID,
		RealName:  "张三",
		IDCard:    "110121202202222222",
		Phone:     "188****8888",
		IsAbroad:  false,
		NotifyURL: "https://www.example.com",
	}
	resp, err := client.UploadUserSign(context.TODO(), req)
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

// GetUploadUserSignStatusExample 获取用户签约状态
func GetUploadUserSignStatusExample(client api.UploadUserSignService) {
	req := &api.GetUploadUserSignStatusRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		RealName: "张三",
		IDCard:   "110121202202222222",
	}
	resp, err := client.GetUploadUserSignStatus(context.TODO(), req)
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

// NotifyUploadUserSignExample 签约成功状态回调通知
func NotifyUploadUserSignExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/uploadusersign", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")

			req := api.NotifyUploadUserSignRequest{}
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
	for _, example := range []func(api.UploadUserSignService){
		UploadUserSignExample,
		GetUploadUserSignStatusExample,
	} {
		example(client)
	}
}
