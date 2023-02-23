package bizlicxjjh5api

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// H5PreCollectBizlicMsg_Example 工商实名信息录入
func H5PreCollectBizlicMsg_Example(client api.BizlicXjjH5APIService) {
	req := &api.H5PreCollectBizlicMsgRequest{
		DealerID:            base.DealerID,
		BrokerID:            base.BrokerID,
		DealerUserID:        "123456_01",
		PhoneNo:             "13412345678",
		IDCard:              "120000000000000000",
		RealName:            "张三",
		IDCardAddress:       "张三",
		IDCardAgency:        "张三",
		IDCardNation:        "张三",
		IDCardValidityStart: "2018-01-21",
		IDCardValidityEnd:   "2027-01-12",
	}
	resp, err := client.H5PreCollectBizlicMsg(context.TODO(), req)
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

// H5APIGetStartUrl_Example 预启动
func H5APIGetStartUrl_Example(client api.BizlicXjjH5APIService) {
	req := &api.H5APIGetStartUrlRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		DealerUserID: "123456_01",
		ClientType:   2,
		NotifyURL:    "http://www.abcdef.com/api/notify",
		ReturnURL:    "",
	}
	resp, err := client.H5APIGetStartUrl(context.TODO(), req)
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

// H5APIEcoCityAicStatus_Example 查询个体工商户状态
func H5APIEcoCityAicStatus_Example(client api.BizlicXjjH5APIService) {
	req := &api.H5APIEcoCityAicStatusRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		DealerUserID: "123456_01",
	}
	resp, err := client.H5APIEcoCityAicStatus(context.TODO(), req)
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
	for _, example := range []func(api.BizlicXjjH5APIService){
		H5PreCollectBizlicMsg_Example,
		H5APIGetStartUrl_Example,
		H5APIEcoCityAicStatus_Example,
	} {
		example(client)
	}
}
