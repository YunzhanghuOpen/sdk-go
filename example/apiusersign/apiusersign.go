package apiusersign

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// ApiUseSignContract_Example 获取协议预览 URL
func ApiUserSignContract_Example(client api.ApiUserSignService) {
	req := &api.ApiUserSignContractRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
	}
	resp, err := client.ApiUserSignContract(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// ApiUserSign_Example 用户签约
func ApiUserSign_Example(client api.ApiUserSignService) {
	req := &api.ApiUserSignRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "121201111111111111111",
		CardType: "idcard",
	}
	resp, err := client.ApiUserSign(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// GetApiUserSignStatus_Example 获取用户签约状态
func GetApiUserSignStatus_Example(client api.ApiUserSignService) {
	req := &api.GetApiUserSignStatusRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "121201111111111111111",
	}
	resp, err := client.GetApiUserSignStatus(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// ApiUserSignRelease_Example 用户解约（测试账号专用接口）
func ApiUserSignRelease_Example(client api.ApiUserSignService) {
	req := &api.ApiUserSignReleaseRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "121201111111111111111",
		CardType: "idcard",
	}
	resp, err := client.ApiUserSignRelease(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.ApiUserSignService){
		ApiUserSignContract_Example,
		ApiUserSign_Example,
		GetApiUserSignStatus_Example,
		ApiUserSignRelease_Example,
	} {
		example(client)
	}
}
