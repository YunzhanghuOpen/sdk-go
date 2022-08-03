package dataservice

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"fmt"
	"net/http"

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
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		fmt.Println(e.Code, e.Message)
		if e.Code == "8300" {
			// 获取日流水链接失败
			// 说明日流水文件还未生成，需稍后再重试

		} else {
			// 其它错误码详见接口文档附录中响应码列表
			fmt.Println(e.Code, e.Message)
		}
		return
	}
	fmt.Println(resp)

	res, err := http.Get(resp.BillDownloadURL)
	if err != nil {
		// 可能为网络错误
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	r, err := gzip.NewReader(res.Body)
	if err != nil {
		// 可能为网络错误
		fmt.Println(err)
		return
	}

	cr := csv.NewReader(r)
	lines, err := cr.ReadAll()
	if err != nil {
		// 说明文件存在问题
		fmt.Println(err)
		return
	}

	fmt.Println(lines)
}

//ListDailyOrder_Example 查询日订单数据
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
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			return
		}
		fmt.Println(e.Code, e.Message)
	}

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
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		fmt.Println(e.Code, e.Message)
		if e.Code == "8300" {
			// 获取日订单链接失败, 说明日订单文件还未生成
			// 需稍后再重试
		} else {
			// 其它错误码详见接口文档附录中响应码列表
			fmt.Println(e.Code, e.Message)
		}
		return
	}
	fmt.Println(resp)

	res, err := http.Get(resp.OrderDownloadURL)
	if err != nil {
		// 可能为网络错误
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	r, err := gzip.NewReader(res.Body)
	if err != nil {
		// 可能为网络错误
		fmt.Println(err)
		return
	}

	cr := csv.NewReader(r)
	lines, err := cr.ReadAll()
	if err != nil {
		// 说明文件存在问题
		fmt.Println(err)
		return
	}

	fmt.Println(lines)
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
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		fmt.Println(e.Code, e.Message)
		if e.Code == "8300" {
			// 获取日订单链接失败, 说明日订单文件还未生成
			// 需稍后再重试
		} else {
			// 其它错误码详见接口文档附录中响应码列表
			fmt.Println(e.Code, e.Message)
		}
		return
	}
	fmt.Println(resp)

	res, err := http.Get(resp.URL)
	if err != nil {
		// 可能为网络错误
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	r, err := gzip.NewReader(res.Body)
	if err != nil {
		// 可能为网络错误
		fmt.Println(err)
		return
	}

	cr := csv.NewReader(r)
	lines, err := cr.ReadAll()
	if err != nil {
		// 说明文件存在问题
		fmt.Println(err)
		return
	}

	fmt.Println(lines)
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
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			return
		}
		fmt.Println(e.Code, e.Message)
	}

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
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			return
		}
		fmt.Println(e.Code, e.Message)
	}

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
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			return
		}
		fmt.Println(e.Code, e.Message)
	}
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
