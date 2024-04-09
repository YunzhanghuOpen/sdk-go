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

// IDCardVerifyExample 身份证实名验证
func IDCardVerifyExample(client api.Authentication) {
	req := &api.IDCardVerifyRequest{
		RealName: "张三",
		IDCard:   "11010519491231002X",
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

// GetBankCardInfoExample 银行卡信息查询接口
func GetBankCardInfoExample(client api.Authentication) {
	req := &api.GetBankCardInfoRequest{
		CardNo: "6228888888888888888",
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

// BankCardThreeVerifyExample 银行卡三要素验证
func BankCardThreeVerifyExample(client api.Authentication) {
	req := &api.BankCardThreeVerifyRequest{
		RealName: "张三",
		IDCard:   "11010519491231002X",
		CardNo:   "6228888888888888888",
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

// BankCardFourAuthConfirmExample 银行卡四要素确认鉴权（上传短信验证码）
func BankCardFourAuthConfirmExample(client api.Authentication) {
	req := &api.BankCardFourAuthConfirmRequest{
		RealName: "张三",
		IDCard:   "11010519491231002X",
		CardNo:   "6228888888888888888",
		Mobile:   "188****8888",
		Captcha:  "011099",
		Ref:      "rx0g4iiLWoWA==",
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

// BankCardFourAuthVerifyExample 银行卡四要素鉴权请求（下发短信验证码）
func BankCardFourAuthVerifyExample(client api.Authentication) {
	req := &api.BankCardFourAuthVerifyRequest{
		RealName: "张三",
		IDCard:   "11010519491231002X",
		CardNo:   "6228888888888888888",
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

// BankCardFourVerifyExample 银行卡四要素验证
func BankCardFourVerifyExample(client api.Authentication) {
	req := &api.BankCardFourVerifyRequest{
		RealName: "张三",
		IDCard:   "11010519491231002X",
		CardNo:   "6228888888888888888",
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

// UserExemptedInfoExample 上传免验证用户名单信息
func UserExemptedInfoExample(client api.Authentication) {
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
		Ref:          "1234qwer",
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

// UserWhiteCheckExample 查看免验证用户名单是否存在
func UserWhiteCheckExample(client api.Authentication) {
	req := &api.UserWhiteCheckRequest{
		RealName: "张三",
		IDCard:   "11010519491231002X",
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
		BankCardFourAuthConfirmExample,
		BankCardFourAuthVerifyExample,
		BankCardFourVerifyExample,
		BankCardThreeVerifyExample,
		GetBankCardInfoExample,
		IDCardVerifyExample,
		UserExemptedInfoExample,
		UserWhiteCheckExample,
	} {
		example(client)
	}
}
