package calculatelabor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// LaborCaculatorExample 连续劳务税费试算（计算器）
func LaborCaculatorExample(client api.CalculateLaborService) {
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

// CalcTax 订单税费试算
func CalcTaxExample(client api.CalculateLaborService) {
	req := &api.CalcTaxRequest{
		BrokerID:              base.BrokerID,
		DealerID:              base.DealerID,
		RealName:              "张三",
		IDCard:                "11010519491231002X",
		Pay:                   "99",
		TaxType:               "before_tax",
		BeforeTaxAmountType:   "max",
		IncludeRecoveryAmount: 1,
		IncludeUserServiceFee: 1,
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

// CalculationYearH5Url 连续劳务年度税费测算-H5
func CalculationYearH5UrlExample(client api.CalculateLaborService) {
	req := &api.CalculationYearH5UrlRequest{
		DealerID:   base.DealerID,
		BrokerID:   base.BrokerID,
		Color:      "#FF3D3D",
		NavbarHide: 0,
		Title:      "云账户",
	}
	resp, err := client.CalculationYearH5Url(context.TODO(), req)
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

// CalculationH5Url 连续劳务单笔结算税费测算-H5
func CalculationH5UrlExample(client api.CalculateLaborService) {
	req := &api.CalculationH5UrlRequest{
		DealerID:   base.DealerID,
		BrokerID:   base.BrokerID,
		RealName:   "张三",
		IDCard:     "11010519491231002X",
		Color:      "#FF3D3D",
		NavbarHide: 0,
		Title:      "云账户",
	}
	resp, err := client.CalculationH5Url(context.TODO(), req)
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
	for _, example := range []func(api.CalculateLaborService){
		LaborCaculatorExample,
		CalcTaxExample,
		CalculationYearH5UrlExample,
		CalculationH5UrlExample,
	} {
		example(client)
	}
}
