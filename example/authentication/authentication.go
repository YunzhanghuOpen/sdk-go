package authentication

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// IDCardVerify_Example 身份证实名验证
func IDCardVerify_Example(client api.Authentication) {
	req := &api.IDCardVerifyRequest{
		RealName: "张三",
		IDCard:   "120000000000000000",
	}
	resp, err := client.IDCardVerify(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// GetBankCardInfo_Example 银行卡信息查询接口
func GetBankCardInfo_Example(client api.Authentication) {
	req := &api.GetBankCardInfoRequest{
		CardNo: "1111111111111111111111111",
	}
	resp, err := client.GetBankCardInfo(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// BankCardThreeVerify_Example 银行卡三要素验证
func BankCardThreeVerify_Example(client api.Authentication) {
	req := &api.BankCardThreeVerifyRequest{
		RealName: "张三",
		IDCard:   "121201111111111111111",
		CardNo:   "1111111111111111111111111",
	}
	resp, err := client.BankCardThreeVerify(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// BankCardFourAuthConfirm_Example 银行卡四要素确认鉴权（上传短信验证码）
func BankCardFourAuthConfirm_Example(client api.Authentication) {
	req := &api.BankCardFourAuthConfirmRequest{
		RealName: "张三",
		IDCard:   "121201111111111111111",
		CardNo:   "1111111111111111111111111",
		Mobile:   "15800001111",
		Captcha:  "011099",
		Ref:      "11111111",
	}
	resp, err := client.BankCardFourAuthConfirm(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// BankCardFourAuthVerify_Example 银行卡四要素鉴权请求（下发短信验证码）
func BankCardFourAuthVerify_Example(client api.Authentication) {
	req := &api.BankCardFourAuthVerifyRequest{
		RealName: "张三",
		IDCard:   "121201111111111111111",
		CardNo:   "1111111111111111111111111",
		Mobile:   "15811111111",
	}
	resp, err := client.BankCardFourAuthVerify(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// BankCardFourVerify_Example 银行卡四要素验证
func BankCardFourVerify_Example(client api.Authentication) {
	req := &api.BankCardFourVerifyRequest{
		RealName: "张三",
		IDCard:   "121201111111111111111",
		CardNo:   "1111111111111111111111111",
		Mobile:   "15811111111",
	}
	resp, err := client.BankCardFourVerify(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// UserExemptedInfo_Example 上传免验证用户名单信息
func UserExemptedInfo_Example(client api.Authentication) {
	srcByte, err := ioutil.ReadFile("example/authentication/test.png")
	if err != nil {
		log.Fatal(err)
	}
	base64str := base64.StdEncoding.EncodeToString(srcByte)

	req := &api.UserExemptedInfoRequest{
		//  证件类型码
		CardType: "passport",
		//  证件号码
		IDCard: "ABC1212011111",
		//  姓名
		RealName: "张三",
		//  申请备注
		CommentApply: "mark",
		//  综合服务主体 ID
		BrokerID: base.BrokerID,
		//  平台企业 ID
		DealerID: base.DealerID,
		//  人员信息图片
		UserImages: []string{base64str, base64str},
		//  国别（地区）代码
		Country: "CHN",
		//  出生日期
		Birthday: "20190809",
		//  性别
		Gender: "男",
		//  回调地址
		NotifyURL: "http://www.callback.com",
		//  请求流水号
		Ref: "111122222",
	}
	resp, err := client.UserExemptedInfo(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// UserWhiteCheck_Example 查看免验证用户名单是否存在
func UserWhiteCheck_Example(client api.Authentication) {
	req := &api.UserWhiteCheckRequest{
		RealName: "张三",
		IDCard:   "121201111111111111111",
	}
	resp, err := client.UserWhiteCheck(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 可能是 SDK 内部处理产生错误(如字符集问题、网络不通等)，请求未能到达服务器
			// 也可能是服务端请求超时，需要稍后重试
			return
		}
		fmt.Println(e.Code, e.Message)
	}
	fmt.Println(resp)
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.Authentication){
		BankCardFourAuthConfirm_Example,
		BankCardFourAuthVerify_Example,
		BankCardFourVerify_Example,
		BankCardThreeVerify_Example,
		GetBankCardInfo_Example,
		IDCardVerify_Example,
		UserExemptedInfo_Example,
		UserWhiteCheck_Example,
	} {
		example(client)
	}
}
