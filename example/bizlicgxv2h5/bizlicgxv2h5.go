package bizlicgxv2h5

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// GxV2H5GetStartUrlExample 预启动
func GxV2H5GetStartUrlExample(client api.BizlicGxV2H5Service) {
	req := &api.GxV2H5GetStartUrlRequest{
		DealerID:      base.DealerID,
		BrokerID:      base.BrokerID,
		DealerUserID:  "userId1234567890",
		ClientType:    1,
		NotifyURL:     "https://www.example.com",
		Color:         "#007AFF",
		ReturnURL:     "https://www.example.com",
		CustomerTitle: 1,
	}
	resp, err := client.GxV2H5GetStartUrl(context.TODO(), req)
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

// GxV2H5GetAicStatusExample 查询个体工商户状态
func GxV2H5GetAicStatusExample(client api.BizlicGxV2H5Service) {
	req := &api.GxV2H5GetAicStatusRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		OpenID:       "openId1234567890",
		RealName:     "张三",
		IDCard:       "11010519491231002X",
		DealerUserID: "userId1234567890",
	}
	resp, err := client.GxV2H5GetAicStatus(context.TODO(), req)
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
	for _, example := range []func(api.BizlicGxV2H5Service){
		GxV2H5GetStartUrlExample,
		GxV2H5GetAicStatusExample,
	} {
		example(client)
	}
}
