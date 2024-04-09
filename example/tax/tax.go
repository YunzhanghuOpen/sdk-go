package tax

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// GetTaxFileExample 下载个税扣缴明细表
func GetTaxFileExample(client api.Tax) {
	req := &api.GetTaxFileRequest{
		DealerID:  base.DealerID,
		EntID:     "accumulus_tj",
		YearMonth: "2022-02",
	}
	resp, err := client.GetTaxFile(context.TODO(), req)
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
	for _, file := range resp.FileInfo {
		pwd, err := base.DecodeZipPwd(file.Pwd) // 解密zip压缩文件密码
		if err != nil {
			return
		}
		fmt.Println(pwd)
	}
}

// GetUserCrossExample 查询纳税人是否为跨集团用户
func GetUserCrossExample(client api.Tax) {
	req := &api.GetUserCrossRequest{
		DealerID: base.DealerID,
		EntID:    "accumulus_tj",
		Year:     "2022",
		IDCard:   "11010519491231002X",
	}
	resp, err := client.GetUserCross(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		// 失败返回
		fmt.Println(e.Code, e.Message)
	}
	// 操作成功
	fmt.Println(resp)
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.Tax){
		GetTaxFileExample,
		GetUserCrossExample,
	} {
		example(client)
	}
}
