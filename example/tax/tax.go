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
		DealerID:  base.DealerID,
		EntID:     "accumulus_tj", // 平台企业签约主体，其中天津：accumulus_tj，甘肃: accumulus_gs（必填）
		YearMonth: "2022-02",      // 所属期（必填）, 注意格式 yyyy-mm
	}
	resp, err := client.GetTaxFile(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是sdk内部错误或网络错误，请求未能连接到服务器
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
				其它错误码详见接口文档附录中响应码列表
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
		EntID:    "accumulus_tj", // 天津：accumulus_tj，甘肃: accumulus_gs（必填）
		Year:     "2022",         // 用户报税所在年份(必填)
		IDCard:   "121201111111111111111",
	}
	resp, err := client.GetUserCross(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是sdk内部错误或网络错误，请求未能连接到服务器
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
