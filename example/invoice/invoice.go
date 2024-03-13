package invoice

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// ApplyInvoiceExample 开票申请
func ApplyInvoiceExample(client api.Invoice) {
	req := &api.ApplyInvoiceRequest{
		DealerID:       base.DealerID,
		BrokerID:       base.BrokerID,
		InvoiceApplyID: "111",
		Amount:         "10000",
		InvoiceType:    "2",
	}
	resp, err := client.ApplyInvoice(context.TODO(), req)
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

// GetInvoiceAmountExample 查询可开票额度和开票信息
func GetInvoiceAmountExample(client api.Invoice) {
	req := &api.GetInvoiceAmountRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
	}
	resp, err := client.GetInvoiceAmount(context.TODO(), req)
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

// GetInvoiceFileExample 下载 PDF 版发票
func GetInvoiceFileExample(client api.Invoice) {
	req := &api.GetInvoiceFileRequest{
		InvoiceApplyID: "test1234567890",
	}
	resp, err := client.GetInvoiceFile(context.TODO(), req)
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

// GetInvoiceStatExample 查询平台企业已开具和待开具发票金额
func GetInvoiceStatExample(client api.Invoice) {
	req := &api.GetInvoiceStatRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		Year:     2022,
	}
	resp, err := client.GetInvoiceStat(context.TODO(), req)
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

// GetInvoiceStatusExample 查询开票申请状态
func GetInvoiceStatusExample(client api.Invoice) {
	req := &api.GetInvoiceStatusRequest{
		InvoiceApplyID: "test1234567890",
	}
	resp, err := client.GetInvoiceStatus(context.TODO(), req)
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

// GetInvoiceInformation 查询发票信息
func GetInvoiceInformationExample(client api.Invoice) {
	req := &api.GetInvoiceInformationRequest{
		InvoiceApplyID: "test1234567890",
		ApplicationID:  "111",
	}
	resp, err := client.GetInvoiceInformation(context.TODO(), req)
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

// SendReminderEmailExample 发送发票扫描件压缩包下载链接邮件
func SendReminderEmailExample(client api.Invoice) {
	req := &api.SendReminderEmailRequest{
		InvoiceApplyID: "test1234567890",
	}
	resp, err := client.SendReminderEmail(context.TODO(), req)
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

// Example 用例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.Invoice){
		ApplyInvoiceExample,
		GetInvoiceAmountExample,
		GetInvoiceFileExample,
		GetInvoiceStatExample,
		GetInvoiceStatusExample,
		GetInvoiceInformationExample,
		SendReminderEmailExample,
	} {
		example(client)
	}
}
