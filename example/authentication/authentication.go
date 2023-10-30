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
		IDCard:   "110121202202222222",
	}
	resp, err := client.IDCardVerify(context.TODO(), req)
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

// GetBankCardInfo_Example 银行卡信息查询接口
func GetBankCardInfo_Example(client api.Authentication) {
	req := &api.GetBankCardInfoRequest{
		CardNo: "8888888888888888888",
	}
	resp, err := client.GetBankCardInfo(context.TODO(), req)
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

// BankCardThreeVerify_Example 银行卡三要素验证
func BankCardThreeVerify_Example(client api.Authentication) {
	req := &api.BankCardThreeVerifyRequest{
		RealName: "张三",
		IDCard:   "110121202202222222",
		CardNo:   "8888888888888888888",
	}
	resp, err := client.BankCardThreeVerify(context.TODO(), req)
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

// BankCardFourAuthConfirm_Example 银行卡四要素确认鉴权（上传短信验证码）
func BankCardFourAuthConfirm_Example(client api.Authentication) {
	req := &api.BankCardFourAuthConfirmRequest{
		RealName: "张三",
		IDCard:   "110121202202222222",
		CardNo:   "8888888888888888888",
		Mobile:   "188****8888",
		Captcha:  "011099",
		Ref:      "11111111",
	}
	resp, err := client.BankCardFourAuthConfirm(context.TODO(), req)
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

// BankCardFourAuthVerify_Example 银行卡四要素鉴权请求（下发短信验证码）
func BankCardFourAuthVerify_Example(client api.Authentication) {
	req := &api.BankCardFourAuthVerifyRequest{
		RealName: "张三",
		IDCard:   "110121202202222222",
		CardNo:   "8888888888888888888",
		Mobile:   "188****8888",
	}
	resp, err := client.BankCardFourAuthVerify(context.TODO(), req)
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

// BankCardFourVerify_Example 银行卡四要素验证
func BankCardFourVerify_Example(client api.Authentication) {
	req := &api.BankCardFourVerifyRequest{
		RealName: "张三",
		IDCard:   "110121202202222222",
		CardNo:   "8888888888888888888",
		Mobile:   "188****8888",
	}
	resp, err := client.BankCardFourVerify(context.TODO(), req)
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

// UserExemptedInfo_Example 上传免验证用户名单信息
func UserExemptedInfo_Example(client api.Authentication) {
	srcByte, err := ioutil.ReadFile("example/authentication/test.png")
	if err != nil {
		log.Fatal(err)
	}
	base64str := base64.StdEncoding.EncodeToString(srcByte)
	req := &api.UserExemptedInfoRequest{
		CardType:     "passport",
		IDCard:       "ABC1212011111",
		RealName:     "张三",
		CommentApply: "mark",
		BrokerID:     base.BrokerID,
		DealerID:     base.DealerID,
		UserImages:   []string{base64str, base64str},
		Country:      "CHN",
		Birthday:     "20190809",
		Gender:       "男",
		NotifyURL:    "https://www.example.com",
		Ref:          "111122222",
	}
	resp, err := client.UserExemptedInfo(context.TODO(), req)
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

// UserWhiteCheck_Example 查看免验证用户名单是否存在
func UserWhiteCheck_Example(client api.Authentication) {
	req := &api.UserWhiteCheckRequest{
		RealName: "张三",
		IDCard:   "110121202202222222",
	}
	resp, err := client.UserWhiteCheck(context.TODO(), req)
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
