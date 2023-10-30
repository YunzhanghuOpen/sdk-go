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
		PhoneNo:             "188****8888",
		IDCard:              "110121202202222222",
		RealName:            "张三",
		IDCardAddress:       "省级行政区名称区县级行政区名称具体住宿地址",
		IDCardAgency:        "区县公安局名称",
		IDCardNation:        "20",
		IDCardValidityStart: "2018-01-21",
		IDCardValidityEnd:   "2027-01-12",
	}
	resp, err := client.H5PreCollectBizlicMsg(context.TODO(), req)
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

// H5APIGetStartUrl_Example 预启动
func H5APIGetStartUrl_Example(client api.BizlicXjjH5APIService) {
	req := &api.H5APIGetStartUrlRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		DealerUserID: "123456_01",
		ClientType:   2,
		NotifyURL:    "https://www.example.com",
		ReturnURL:    "https://www.example.com",
	}
	resp, err := client.H5APIGetStartUrl(context.TODO(), req)
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
	for _, example := range []func(api.BizlicXjjH5APIService){
		H5PreCollectBizlicMsg_Example,
		H5APIGetStartUrl_Example,
		H5APIEcoCityAicStatus_Example,
	} {
		example(client)
	}
}
