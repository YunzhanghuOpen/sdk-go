package dataservice

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// GetDailyBillFileV2_Example 查询日流水文件
func GetDailyBillFileV2_Example(client api.DataService) {
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

// ListDailyOrder_Example 查询日订单数据
func ListDailyOrder_Example(client api.DataService) {
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

// GetOrderDownloadsUrl_Example 查询日订单文件
func GetOrderDownloadsUrl_Example(client api.DataService) {
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

// GetDailyOrderFileV2_Example 查询日订单文件（支付和退款订单）
func GetDailyOrderFileV2_Example(client api.DataService) {
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

// ListDailyBill_Example 查询日流水数据
func ListDailyBill_Example(client api.DataService) {
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

// ListDealerRechargeRecordV2_Example 查询平台企业预付业务服务费记录
func ListDealerRechargeRecordV2_Example(client api.DataService) {
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

// ListBalanceDailyStatement_Example 获取余额日账单
func ListBalanceDailyStatement_Example(client api.DataService) {
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
		GetDailyBillFileV2_Example,
		ListDailyOrder_Example,
		GetOrderDownloadsUrl_Example,
		GetDailyOrderFileV2_Example,
		ListDailyBill_Example,
		ListDealerRechargeRecordV2_Example,
		ListBalanceDailyStatement_Example,
	} {
		example(client)
	}
}
