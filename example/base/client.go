package base

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/core"
	"github.com/YunzhanghuOpen/sdk-go/crypto"
)

// 接口配置信息
var (
	BrokerID      = "testbroker"
	DealerID      = "testdealer"
	AppKey        = "appkey"
	Des3Key       = "3deskey"
	YunPrivateKey = privateKey
	YunPublicKey  = publicKey
)

// 云账户公钥
var publicKey string = "XXXXX"

// 平台企业私钥
var privateKey string = "XXXXX"

// NewClient 新建 client Core
func NewClient() *api.Client {
	rand.Seed(time.Now().UnixNano())
	// 注入 request-id 中间件
	// request-id：请求 ID，请求的唯一标识
	// 建议平台企业自定义 request-id，并记录在日志中，便于问题发现及排查
	// 如未自定义 request-id，将使用 SDK 中的 random 方法自动生成。注意：random 方法生成的 request-id 不能保证全局唯一，推荐自定义 request-id
	requestIDMiddle := func(next core.Handler) core.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return next(core.WithRequestID(ctx, fmt.Sprint(rand.Int63())), req)
		}
	}

	// 请求返回打印中间件
	logMiddle := func(next core.Handler) core.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			fmt.Println("middleware begin ", "requestID=", core.RequestID(ctx), " req=", req)
			resp, err := next(ctx, req)
			fmt.Println("middleware done  ", "requestID=", core.RequestID(ctx), " resp=", resp)
			return resp, err
		}
	}

	conf := &api.Config{
		Host:       api.SandboxHost, // 沙箱环境域名，正式上线时使用 api.ProductHost，个体工商户注册使用 api.AicProductHost
		DealerID:   DealerID,
		PrivateKey: YunPrivateKey,
		AppKey:     AppKey,
		Des3Key:    Des3Key,
	}
	c, err := api.New(conf,
		core.WithHttpClient(&http.Client{
			Timeout: 30 * time.Second, // 可自定义超时时间
		}),
		core.WithMiddleware(requestIDMiddle, logMiddle),
		core.EnDebug(),
	)
	if err != nil {
		panic(err)
	}
	return c
}

// NewRsaSignVerifier 新建 RSA 签名验签
func NewRsaSignVerifier() crypto.SignVerifier {
	v, err := crypto.NewRsaSignVerifier(YunPublicKey, AppKey)
	if err != nil {
		panic(err)
	}
	return v
}

// NewHmacSignVerifier 新建 HMAC 签名
func NewHmacSignVerifier() crypto.SignVerifier {
	v, err := crypto.NewHmacSignVerifier(AppKey)
	if err != nil {
		panic(err)
	}
	return v
}

// NewDes3Decoder 新建加密解密对象
func NewDes3Decoder() crypto.Decoder {
	return crypto.NewDes3Encoding(Des3Key)
}

func NotifyDecoder(mess, timestamp, data, sign, signType string, out interface{}) error {
	sv, ok := map[string]func() crypto.SignVerifier{
		"SHA256": NewHmacSignVerifier,
		"RSA":    NewRsaSignVerifier,
	}[strings.ToUpper(signType)]
	if !ok {
		return errors.New("wrong sign type")
	}

	ok, err := sv().Verify(mess, timestamp, data, sign)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("签名错误")
	}

	decoder := NewDes3Decoder()
	b, err := decoder.Decode([]byte(data))
	if err != nil {
		return err
	}
	req := &out
	fmt.Println(string(b))
	return json.Unmarshal(b, req)
}

// NewRsaDecoder 新建 RSA 解密接口
func NewRsaDecoder() (crypto.Decoder, error) {
	return crypto.NewRsaDecoder([]byte(privateKey))
}

// DecodeZipPwd 解密 ZIP 密码
func DecodeZipPwd(pwdStr string) (string, error) {
	d, err := NewRsaDecoder()
	if err != nil {
		return "", err
	}
	b, err := d.Decode([]byte(pwdStr))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
