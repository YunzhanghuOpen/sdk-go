package h5usersign

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// H5UserPresignExample H5 预申请签约
func H5UserPresignExample(client api.H5UserSignService) {
	req := &api.H5UserPresignRequest{
		DealerID:        base.DealerID,
		BrokerID:        base.BrokerID,
		RealName:        "张三",
		IDCard:          "11010519491231002X",
		CertificateType: 0,
		CollectPhoneNo: 0,
	}
	resp, err := client.H5UserPresign(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// H5UserSignExample H5 签约
func H5UserSignExample(client api.H5UserSignService) {
	req := &api.H5UserSignRequest{
		Token: "X06603X195",
		URL:   "https://www.example.com",
	}
	resp, err := client.H5UserSign(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// GeH5UserSignStatusExample 获取用户签约状态
func GeH5UserSignStatusExample(client api.H5UserSignService) {
	req := &api.GetH5UserSignStatusRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		RealName: "张三",
		IDCard:   "11010519491231002X",
	}
	resp, err := client.GetH5UserSignStatus(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// H5UserReleaseExample H5 对接测试解约接口
func H5UserReleaseExample(client api.H5UserSignService) {
	req := &api.H5UserReleaseRequest{
		DealerID:        base.DealerID,
		BrokerID:        base.BrokerID,
		RealName:        "张三",
		IDCard:          "11010519491231002X",
		CertificateType: 0,
	}
	resp, err := client.H5UserRelease(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.H5UserSignService){
		H5UserPresignExample,
		H5UserSignExample,
		GeH5UserSignStatusExample,
		H5UserReleaseExample,
	} {
		example(client)
	}
}
