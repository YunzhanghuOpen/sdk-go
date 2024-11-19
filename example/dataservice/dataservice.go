package dataservice

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// GetDailyBillFileV2Example 查询日流水文件
func GetDailyBillFileV2Example(client api.DataService) {
	req := &api.GetDailyBillFileV2Request{
		BillDate: "2022-01-01",
	}
	resp, err := client.GetDailyBillFileV2(context.TODO(), req)
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

// ListDailyOrderExample 查询日订单数据
func ListDailyOrderExample(client api.DataService) {
	req := &api.ListDailyOrderRequest{
		OrderDate: "2022-03-23",
		Offset:    1,
		Length:    20,
		Channel:   "银行卡",
	}
	resp, err := client.ListDailyOrder(context.TODO(), req)
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

// ListDailyOrderV2Example 查询日订单数据（支付和退款订单）
func ListDailyOrderV2Example(client api.DataService) {
	req := &api.ListDailyOrderV2Request{
		OrderDate: "2024-09-05",
		Offset:    0,
		Length:    100,
		Channel:   "alipay",
		DataType:  "",
	}
	resp, err := client.ListDailyOrderV2(context.TODO(), req)
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

// GetOrderDownloadsUrlExample 查询日订单文件
func GetOrderDownloadsUrlExample(client api.DataService) {
	req := &api.GetDailyOrderFileRequest{
		OrderDate: "2022-03-23",
	}
	resp, err := client.GetDailyOrderFile(context.TODO(), req)
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

// GetDailyOrderFileV2Example 查询日订单文件（支付和退款订单）
func GetDailyOrderFileV2Example(client api.DataService) {
	req := &api.GetDailyOrderFileV2Request{
		OrderDate: "2022-03-23",
	}
	resp, err := client.GetDailyOrderFileV2(context.TODO(), req)
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

// ListDailyBillExample 查询日流水数据
func ListDailyBillExample(client api.DataService) {
	req := &api.ListDailyBillRequest{
		BillDate: "2022-03-23",
		Offset:   1,
		Length:   100,
	}
	resp, err := client.ListDailyBill(context.TODO(), req)
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

// ListDealerRechargeRecordV2Example 查询平台企业预付业务服务费记录
func ListDealerRechargeRecordV2Example(client api.DataService) {
	req := &api.ListDealerRechargeRecordV2Request{
		BeginAt: "2022-01-01",
		EndAt:   "2022-04-01",
	}
	resp, err := client.ListDealerRechargeRecordV2(context.TODO(), req)
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

// ListBalanceDailyStatementExample 获取余额日账单
func ListBalanceDailyStatementExample(client api.DataService) {
	req := &api.ListBalanceDailyStatementRequest{
		StatementDate: "2022-03-23",
	}
	resp, err := client.ListBalanceDailyStatement(context.TODO(), req)
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
	for _, example := range []func(api.DataService){
		GetDailyBillFileV2Example,
		ListDailyOrderExample,
		ListDailyOrderV2Example,
		GetOrderDownloadsUrlExample,
		GetDailyOrderFileV2Example,
		ListDailyBillExample,
		ListDealerRechargeRecordV2Example,
		ListBalanceDailyStatementExample,
	} {
		example(client)
	}
}
