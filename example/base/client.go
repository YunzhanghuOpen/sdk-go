package base

import (
	"context"
	_ "embed"
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

//go:embed testdata/mch_public_key.txt
var publicKey string

//go:embed testdata/mch_private_key.txt
var privateKey string

// NewCore 新建 client Core
func NewClient() *api.Client {
	rand.Seed(time.Now().UnixNano())
	// 注入requestID中间件
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
		Host:       api.SandboxHost, // 沙箱环境域名，正式上线时使用 api.ProductHost
		DealerID:   DealerID,
		PrivateKey: YunPrivateKey,
		AppKey:     AppKey,
		Des3Key:    Des3Key,
	}
	c, err := api.New(conf,
		core.WithHttpClient(&http.Client{}),
		core.WithMiddleware(requestIDMiddle, logMiddle),
		core.EnDebug(),
	)
	if err != nil {
		panic(err)
	}
	return c
}

// NewRsaSignVerifier 新建 RSA 签名验签
func NewRsaSignVerifier(dealerID string) crypto.SignVerifier {
	v, err := crypto.NewRsaSignVerifier(YunPublicKey, AppKey)
	if err != nil {
		panic(err)
	}
	return v
}

// NewHmacSignVerifier 新建 RSA 签名
func NewHmacSignVerifier(dealerID string) crypto.SignVerifier {
	v, err := crypto.NewHmacSignVerifier(AppKey)
	if err != nil {
		panic(err)
	}
	return v
}

// NewEncoder 新建加密解密对象
func NewDes3Decoder(dealerID string) crypto.Decoder {
	return crypto.NewDes3Encoding(Des3Key)
}

func NotifyDecoder(dealerID, mess, timestamp, data, sign, signType string, out interface{}) error {
	sv, ok := map[string]crypto.SignVerifier{
		"SHA256": NewHmacSignVerifier(dealerID),
		"RSA":    NewRsaSignVerifier(dealerID),
	}[strings.ToUpper(signType)]
	if !ok {
		return errors.New("wrong sign type")
	}

	ok, err := sv.Verify(mess, timestamp, data, sign)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("签名错误")
	}

	decoder := NewDes3Decoder(dealerID)
	b, err := decoder.Decode([]byte(data))
	if err != nil {
		return err
	}

	req := &struct {
		NotifyID   string      `json:"notify_id,omitempty"`
		NotifyTime string      `json:"notify_time,omitempty"`
		Data       interface{} `json:"data,omitempty"`
	}{
		Data: out,
	}
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
