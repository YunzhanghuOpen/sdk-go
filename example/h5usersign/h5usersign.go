package h5usersign

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// H5UserPresign_Example H5 预申请签约
func H5UserPresign_Example(client api.H5UserSignService) {
	req := &api.H5UserPresignRequest{
		DealerID:        base.DealerID,
		BrokerID:        base.BrokerID,
		RealName:        "张三",
		IDCard:          "121201111111111111111",
		CertificateType: 0,
	}
	resp, err := client.H5UserPresign(context.TODO(), req)
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

// H5UserSign_Example H5 签约
func H5UserSign_Example(client api.H5UserSignService) {
	req := &api.H5UserSignRequest{
		Token: base.DealerID,
		URL:   "http://www.and.com/cabck",
	}
	resp, err := client.H5UserSign(context.TODO(), req)
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

// GeH5UserSignStatus_Example 获取用户签约状态
func GeH5UserSignStatus_Example(client api.H5UserSignService) {
	req := &api.GetH5UserSignStatusRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "121201111111111111111",
	}
	resp, err := client.GetH5UserSignStatus(context.TODO(), req)
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

// H5UserRelease_Example H5 对接测试解约接口
func H5UserRelease_Example(client api.H5UserSignService) {
	req := &api.H5UserReleaseRequest{
		DealerID:        base.DealerID,
		BrokerID:        base.BrokerID,
		RealName:        "张三",
		IDCard:          "121201111111111111111",
		CertificateType: 0,
	}
	resp, err := client.H5UserRelease(context.TODO(), req)
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
	for _, example := range []func(api.H5UserSignService){
		H5UserPresign_Example,
		H5UserSign_Example,
		GeH5UserSignStatus_Example,
		H5UserRelease_Example,
	} {
		example(client)
	}
}
