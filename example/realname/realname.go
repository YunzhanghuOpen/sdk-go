package realname

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// CollectRealNameInfoFaceExample 用户实名认证信息收集-人脸认证方式
func CollectRealNameInfoFaceExample(client api.RealNameService) {
	req := &api.CollectRealNameInfoRequest{
		BrokerID:             base.BrokerID,
		DealerID:             base.DealerID,
		RealName:             "张三",
		IDCard:               "110xxxxxxxxxxxxx16",
		RealnameResult:       1,
		RealnameTime:         "2025-09-09 19:19:19",
		RealnameType:         1,
		RealnameTraceID:      "1413536187796566016",
		RealnamePlatform:     "xxxxxxx公司",
		FaceImageCollectType: 1,
		FaceImage:            "https://www.example.com/file_name.png",
		FaceVerifyScore:      "89.12",
	}
	resp, err := client.CollectRealNameInfo(context.TODO(), req)
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
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// CollectRealNameInfoBankExample 用户实名认证信息收集-银行卡四要素认证方式
func CollectRealNameInfoBankExample(client api.RealNameService) {
	req := &api.CollectRealNameInfoRequest{
		BrokerID:         base.BrokerID,
		DealerID:         base.DealerID,
		RealName:         "张三",
		IDCard:           "110xxxxxxxxxxxxx16",
		RealnameResult:   1,
		RealnameTime:     "2025-09-09 19:19:19",
		RealnameType:     2,
		RealnameTraceID:  "1413536187796566016",
		RealnamePlatform: "xxxxxxx公司",
		BankNo:           "6127000xxxxxxx16",
		BankPhone:        "188xxx8888",
	}
	resp, err := client.CollectRealNameInfo(context.TODO(), req)
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
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// CollectRealNameInfoReviewerExample 用户实名认证信息收集-人工审核
func CollectRealNameInfoReviewerExample(client api.RealNameService) {
	req := &api.CollectRealNameInfoRequest{
		BrokerID:             base.BrokerID,
		DealerID:             base.DealerID,
		RealName:             "张三",
		IDCard:               "110xxxxxxxxxxxxx16",
		RealnameResult:       1,
		RealnameTime:         "2025-09-09 19:19:19",
		RealnameType:         3,
		RealnameTraceID:      "1413536187796566016",
		RealnamePlatform:     "xxxxxxx公司",
		FaceImageCollectType: 1,
		FaceImage:            "https://www.example.com/file_name.png",
		FaceVerifyScore:      "89.12",
		BankNo:               "6127000xxxxxxx16",
		BankPhone:            "188xxx8888",
		Reviewer:             "908xxx8888",
	}
	resp, err := client.CollectRealNameInfo(context.TODO(), req)
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
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// QueryRealNameInfoExample 用户实名认证信息查询
func QueryRealNameInfoExample(client api.RealNameService) {
	req := &api.QueryRealNameInfoRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		RealName: "张三",
		IDCard:   "110xxxxxxxxxxxxx16",
	}
	resp, err := client.QueryRealNameInfo(context.TODO(), req)
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
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.RealNameService){
		CollectRealNameInfoFaceExample,
		CollectRealNameInfoBankExample,
		CollectRealNameInfoReviewerExample,
		QueryRealNameInfoExample,
	} {
		example(client)
	}
}
