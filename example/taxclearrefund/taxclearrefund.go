package taxclearrefund

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// NotifyClearTaxDoneExample 连续劳务税费清缴完成通知
func NotifyClearTaxDoneExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/order", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyClearTaxDoneRequest{}
			err := base.NotifyDecoder(mess, timestamp, data, sign, signType, &req)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}
	}))
}

// GetClearTaxInfoExample 查询税费清缴完成结果
func GetClearTaxInfoExample(client api.TaxClearRefundService) {
	req := &api.GetClearTaxInfoRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		TaxMonth: "2025-10",
	}
	resp, err := client.GetClearTaxInfo(context.TODO(), req)
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

// GetClearTaxFileExample 查询税费清缴明细文件
func GetClearTaxFileExample(client api.TaxClearRefundService) {
	req := &api.GetClearTaxFileRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		TaxMonth: "2025-10",
		BatchID:  "10313232135454132",
	}
	resp, err := client.GetClearTaxFile(context.TODO(), req)
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

// NotifyRefundTaxDoneExample 连续劳务税费退补完成通知
func NotifyRefundTaxDoneExample() {
	// 除本实现方式外，还可采用其他 http 请求框架实现
	http.HandleFunc("notify/order", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() == nil {
			data := r.PostForm.Get("data")
			timestamp := r.PostForm.Get("timestamp")
			mess := r.PostForm.Get("mess")
			sign := r.PostForm.Get("sign")
			signType := r.PostForm.Get("sign_type")
			req := api.NotifyRefundTaxDoneRequest{}
			err := base.NotifyDecoder(mess, timestamp, data, sign, signType, &req)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}
	}))
}

// GetRefundTaxInfoExample 查询税费退补完成结果
func GetRefundTaxInfoExample(client api.TaxClearRefundService) {
	req := &api.GetRefundTaxInfoRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		TaxMonth: "2025-10",
		BatchID:  "10313232135454132",
	}
	resp, err := client.GetRefundTaxInfo(context.TODO(), req)
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

// GetRefundTaxLaborInfoExample 查询税费退补涉及劳动者
func GetRefundTaxLaborInfoExample(client api.TaxClearRefundService) {
	req := &api.GetRefundTaxLaborInfoRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		BatchID:  "10313232135454132",
		TaxMonth: "2025-10",
		Offset:   0,
		Length:   500,
	}
	resp, err := client.GetRefundTaxLaborInfo(context.TODO(), req)
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

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.TaxClearRefundService){
		GetClearTaxInfoExample,
		GetClearTaxFileExample,
		GetRefundTaxInfoExample,
		GetRefundTaxLaborInfoExample,
	} {
		example(client)
	}
}
