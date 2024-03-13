package apiusersign

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// ApiUseSignContractExample 获取协议预览 URL
func ApiUserSignContractExample(client api.ApiUserSignService) {
	req := &api.ApiUserSignContractRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
	}
	resp, err := client.ApiUserSignContract(context.TODO(), req)
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

// ApiUserSignExample 用户签约
func ApiUserSignExample(client api.ApiUserSignService) {
	req := &api.ApiUserSignRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "110121202202222222",
		CardType: "idcard",
	}
	resp, err := client.ApiUserSign(context.TODO(), req)
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

// GetApiUserSignStatusExample 获取用户签约状态
func GetApiUserSignStatusExample(client api.ApiUserSignService) {
	req := &api.GetApiUserSignStatusRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "110121202202222222",
	}
	resp, err := client.GetApiUserSignStatus(context.TODO(), req)
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

// ApiUserSignReleaseExample 用户解约（测试账号专用接口）
func ApiUserSignReleaseExample(client api.ApiUserSignService) {
	req := &api.ApiUserSignReleaseRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "110121202202222222",
		CardType: "idcard",
	}
	resp, err := client.ApiUserSignRelease(context.TODO(), req)
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
	for _, example := range []func(api.ApiUserSignService){
		ApiUserSignContractExample,
		ApiUserSignExample,
		GetApiUserSignStatusExample,
		ApiUserSignReleaseExample,
	} {
		example(client)
	}
}
