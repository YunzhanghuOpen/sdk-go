package invoice

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

/**
* 开票申请
 */
func ApplyInvoice_Example(client api.Invoice) {
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
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
*查询可开票额度和开票信息
 */
func GetInvoiceAmount_Example(client api.Invoice) {
	req := &api.GetInvoiceAmountRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
	}
	resp, err := client.GetInvoiceAmount(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
* 下载发票 PDF
 */
func GetInvoiceFile_Example(client api.Invoice) {
	req := &api.GetInvoiceFileRequest{
		InvoiceApplyID: "111",
	}
	resp, err := client.GetInvoiceFile(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
* 查询平台企业已开具和待开具发票金额
 */
func GetInvoiceStat_Example(client api.Invoice) {
	req := &api.GetInvoiceStatRequest{
		DealerID: base.DealerID,
		BrokerID: base.BrokerID,
		Year:     2022,
	}
	resp, err := client.GetInvoiceStat(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
* 查询开票申请状态
 */
func GetInvoiceStatus_Example(client api.Invoice) {
	req := &api.GetInvoiceStatusRequest{
		InvoiceApplyID: "111",
	}
	resp, err := client.GetInvoiceStatus(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

/**
* 发送发票扫描件压缩包下载链接邮件
 */
func SendReminderEmail_Example(client api.Invoice) {
	req := &api.SendReminderEmailRequest{
		InvoiceApplyID: "111",
	}
	resp, err := client.SendReminderEmail(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

func Example() {
	client := base.NewClient()
	for _, example := range []func(api.Invoice){
		ApplyInvoice_Example,
		GetInvoiceAmount_Example,
		GetInvoiceFile_Example,
		GetInvoiceStat_Example,
		GetInvoiceStatus_Example,
		SendReminderEmail_Example,
	} {
		example(client)
	}
}
