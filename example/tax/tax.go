package tax

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

/**
 * 下载个税扣缴明细表
 */
func GetTaxFile_Example(client api.Tax) {
	req := &api.GetTaxFileRequest{
		DealerID:  "填写自己的",
		EntID:     "填写自己的",   // 商户签约主体，其中天津：accumulus_tj，甘肃: accumulus_gs（必填）
		YearMonth: "2022-02", // 所属期（必填）, 注意格式 yyyy-mm
	}
	resp, err := client.GetTaxFile(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		fmt.Println(e.Code, e.Message)
		if e.Code == "6201" {
			/*
				个税扣缴明细表不存在
				说明个税扣缴明还未生成
				需稍后再重试
			*/
			fmt.Println(e.Code, e.Message)
			return

		} else {
			/*
				其它错误详见文档
				可参考 e.message 中的错误信息
			*/

			fmt.Println(e.Code, e.Message)
		}
		return
	}

	fmt.Println(resp)
	for _, file := range resp.FileInfo {
		pwd, err := base.DecodeZipPwd(file.Pwd) // 解密zip压缩文件密码
		if err != nil {
			return
		}
		fmt.Println(pwd)
		// TODO: 下载文件

		// TODO: 解压文件
	}
}

/**
* 查询纳税人是否为跨集团用户
 */
func GetUserCross_Example(client api.Tax) {
	req := &api.GetUserCrossRequest{
		DealerID: base.DealerID,
		EntID:    "填写自己的", // 商户签约主体，其中天津：accumulus_tj，甘肃: accumulus_gs（必填）
		Year:     "2022",  // 用户报税所在年份(必填)
		IDCard:   "填写自己的",
	}
	resp, err := client.GetUserCross(context.TODO(), req)
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
	for _, example := range []func(api.Tax){
		GetTaxFile_Example,
		GetUserCross_Example,
	} {
		example(client)
	}
}
