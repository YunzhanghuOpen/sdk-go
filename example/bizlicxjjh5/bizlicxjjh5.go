package bizlicxjjh5

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// H5GetStartUrl_Example 预启动
func H5GetStartUrl_Example(client api.BizlicXjjH5Service) {
	req := &api.H5GetStartUrlRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		DealerUserID: "123456_01",
		ClientType:   1,
		NotifyURL:    "http://www.abcdef.com/api/notify",
		ReturnURL:    "",
	}
	resp, err := client.H5GetStartUrl(context.TODO(), req)
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

// H5EcoCityAicStatus_Example 查询个体工商户状态
func H5EcoCityAicStatus_Example(client api.BizlicXjjH5Service) {
	req := &api.H5EcoCityAicStatusRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		DealerUserID: "123456_01",
	}
	resp, err := client.H5EcoCityAicStatus(context.TODO(), req)
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
	for _, example := range []func(api.BizlicXjjH5Service){
		H5GetStartUrl_Example,
		H5EcoCityAicStatus_Example,
	} {
		example(client)
	}
}
