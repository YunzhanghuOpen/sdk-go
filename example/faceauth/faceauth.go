package faceauth

import (
	"context"
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// ApplyFaceAuthExample 申请人脸识别实名核验
func ApplyFaceAuthExample(client api.FaceAuthService) {
	req := &api.ApplyFaceAuthRequest{
		BrokerID:       base.BrokerID,
		DealerID:       base.DealerID,
		VerificationID: "verificationExampleId123456",
		RealName:       "张三",
		IDCard:         "11010519491231002X",
		CallbackURL:    "https://www.example.com/callback",
		RedirectURL:    "https://www.example.com/success",
		Color:          "#FF0000",
	}
	resp, err := client.ApplyFaceAuth(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		fmt.Println(e.Code, e.Message)
		return
	}
	// 操作成功
	fmt.Println(resp)
}

// GetFaceAuthResultExample 查询人脸识别实名核验结果
func GetFaceAuthResultExample(client api.FaceAuthService) {
	req := &api.GetFaceAuthResultRequest{
		BrokerID:       base.BrokerID,
		DealerID:       base.DealerID,
		RecordID:       "recoreExampleId123456",
		VerificationID: "verificationExampleId123456",
	}
	resp, err := client.GetFaceAuthResult(context.TODO(), req)
	if err != nil {
		e, ok := errorx.FromError(err)
		if !ok {
			// 发生异常
			fmt.Println(err)
			return
		}
		fmt.Println(e.Code, e.Message)
		return
	}
	// 操作成功
	fmt.Println(resp)
}

// Example 样例
func Example() {
	client := base.NewClient()
	for _, example := range []func(api.FaceAuthService){
		ApplyFaceAuthExample,
		GetFaceAuthResultExample,
	} {
		example(client)
	}
}
