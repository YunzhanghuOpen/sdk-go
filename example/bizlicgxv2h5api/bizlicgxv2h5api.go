package bizlicgxv2h5api

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// GxV2H5APIPreCollectBizlicMsgExample 工商实名信息录入
func GxV2H5APIPreCollectBizlicMsgExample(client api.BizlicGxV2H5APIService) {
	req := &api.GxV2H5APIPreCollectBizlicMsgRequest{
		DealerID:            base.DealerID,
		BrokerID:            base.BrokerID,
		DealerUserID:        "userId1234567890",
		PhoneNo:             "+86-188****8888",
		IDCard:              "11010519491231002X",
		RealName:            "张三",
		IDCardAddress:       "省级行政区名称区县级行政区名称具体住宿地址",
		IDCardAgency:        "区县公安局名称",
		IDCardNation:        "20",
		IDCardValidityStart: "2022-02-22",
		IDCardValidityEnd:   "2042-02-22",
	}
	resp, err := client.GxV2H5APIPreCollectBizlicMsg(context.TODO(), req)
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

// GxV2H5APIGetStartUrlExample 预启动
func GxV2H5APIGetStartUrlExample(client api.BizlicGxV2H5APIService) {
	req := &api.GxV2H5APIGetStartUrlRequest{
		DealerID:      base.DealerID,
		BrokerID:      base.BrokerID,
		DealerUserID:  "userId1234567890",
		ClientType:    2,
		NotifyURL:     "https://www.example.com",
		Color:         "#007AFF",
		ReturnURL:     "https://www.example.com",
		CustomerTitle: 1,
	}
	resp, err := client.GxV2H5APIGetStartUrl(context.TODO(), req)
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

// GxV2H5APIGetAicStatusExample 查询个体工商户状态
func GxV2H5APIGetAicStatusExample(client api.BizlicGxV2H5APIService) {
	req := &api.GxV2H5APIGetAicStatusRequest{
		DealerID:     base.DealerID,
		BrokerID:     base.BrokerID,
		OpenID:       "openId1234567890",
		RealName:     "张三",
		IDCard:       "11010519491231002X",
		DealerUserID: "userId1234567890",
	}
	resp, err := client.GxV2H5APIGetAicStatus(context.TODO(), req)
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
	for _, example := range []func(api.BizlicGxV2H5APIService){
		GxV2H5APIPreCollectBizlicMsgExample,
		GxV2H5APIGetStartUrlExample,
		GxV2H5APIGetAicStatusExample,
	} {
		example(client)
	}
}
