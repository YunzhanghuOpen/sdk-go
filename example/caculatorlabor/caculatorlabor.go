package caculatorlabor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// LaborCaculatorExample 连续劳务税费试算（计算器）
func LaborCaculatorExample(client api.CaculatorLaborService) {
	monthSettlementList := []*api.MonthSettlement{
		{
			Month:             1,
			MonthPreTaxAmount: "10.00",
		},
		{
			Month:             2,
			MonthPreTaxAmount: "10.00",
		},
	}

	req := &api.LaborCaculatorRequest{
		BrokerID:            base.BrokerID,
		DealerID:            base.DealerID,
		MonthSettlementList: monthSettlementList,
	}
	resp, err := client.LaborCaculator(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		errorData, err := json.Marshal(e)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(errorData))
	}
	// 操作成功
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// CalcTax 订单税费计算
func CalcTaxExample(client api.CaculatorLaborService) {
	req := &api.CalcTaxRequest{
		BrokerID: base.BrokerID,
		DealerID: base.DealerID,
		RealName: "张三",
		IDCard:   "11010519491231002X",
		Pay:      "99",
	}
	resp, err := client.CalcTax(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		errorData, err := json.Marshal(e)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(errorData))
	}
	// 操作成功
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.CaculatorLaborService){
		LaborCaculatorExample,
		CalcTaxExample,
	} {
		example(client)
	}
}
