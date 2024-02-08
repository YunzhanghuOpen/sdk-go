package bizlicgxh5api

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// GxH5APIPreCollectBizlicMsg_Example 工商实名信息录入
func GxH5APIPreCollectBizlicMsg_Example(client api.BizlicGxH5APIService) {
	req := &api.GxH5APIPreCollectBizlicMsgRequest{
		DealerID:            base.DealerID,
		BrokerID:            base.BrokerID,
		DealerUserID:        "123456_01",
		PhoneNo:             "188****8888",
		IDCard:              "110121202202222222",
		RealName:            "张三",
		IDCardAddress:       "省级行政区名称区县级行政区名称具体住宿地址",
		IDCardAgency:        "区县公安局名称",
		IDCardNation:        "20",
		IDCardValidityStart: "2018-01-21",
		IDCardValidityEnd:   "2027-01-12",
	}
	resp, err := client.GxH5APIPreCollectBizlicMsg(context.TODO(), req)
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

// GxH5APIGetStartUrl_Example 预启动
func GxH5APIGetStartUrl_Example(client api.BizlicGxH5APIService) {
	req := &api.GxH5APIGetStartUrlRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		DealerUserID: "123456_01",
		ClientType:   2,
		NotifyURL:    "https://www.example.com",
		ReturnURL:    "https://www.example.com",
	}
	resp, err := client.GxH5APIGetStartUrl(context.TODO(), req)
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

// GxH5APIEcoCityAicStatus_Example 查询个体工商户状态
func GxH5APIEcoCityAicStatus_Example(client api.BizlicGxH5APIService) {
	req := &api.GxH5APIEcoCityAicStatusRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		DealerUserID: "123456_01",
	}
	resp, err := client.GxH5APIEcoCityAicStatus(context.TODO(), req)
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
	for _, example := range []func(api.BizlicGxH5APIService){
		GxH5APIPreCollectBizlicMsg_Example,
		// GxH5APIGetStartUrl_Example,
		// GxH5APIEcoCityAicStatus_Example,
	} {
		example(client)
	}
}
