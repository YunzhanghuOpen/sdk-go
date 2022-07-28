package apiusersign

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

/**
* 获取协议预览 URL
 */
func ApiUseSignContract_Example(client api.ApiUserSignService) {
	req := &api.ApiUseSignContractRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
	}
	resp, err := client.ApiUseSignContract(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是sdk内部错误或网络错误，请求未能连接到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
* 用户签约
 */
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
			// 可能是sdk内部错误或网络错误，请求未能连接到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
* 获取用户签约状态
 */
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
			// 可能是sdk内部错误或网络错误，请求未能连接到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
* 用户解约（测试账号专用接口）
 */
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
			// 可能是sdk内部错误或网络错误，请求未能连接到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

func Example() {
	client := base.NewClient()
	for _, example := range []func(api.ApiUserSignService){
		ApiUseSignContract_Example,
		ApiUserSign_Example,
		GetApiUserSignStatus_Example,
		ApiUserSignRelease_Example,
	} {
		example(client)
	}
}
