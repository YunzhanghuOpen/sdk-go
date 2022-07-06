# 云账户 SDK For Golang
## 概览

### 简介
欢迎使用 云账户 SDK For Golang，这里向您介绍如何快速使用云账户 SDK For Golang 。
云账户 SDK For Golang  包含了请求的封装、加解密、签名验签等功能。
如果您在使用 云账户 SDK For Golang 的过程中遇到任何问题，欢迎在当前 GitHub 提交 Issues 或发送邮件至技术支持 <techsupport@yunzhanghu.com> 。

### 环境要求
- go1.16+

### 获取配置
#### 获取 dealer_id、broker_id、3DES Key、App Key

根据开户邮件中的账号登录[云账户综合服务平台](https://service.yunzhanghu.com)

![获取配置信息](.doc/keyconfig.png)

#### 生成密钥

- 方式一：使用 OpenSSL 生成 RSA 公私钥

```
① ⽣成私钥 private_key.pem

Openssl-> genrsa -out private_key.pem 2048 # 建议密钥⻓度⾄少为2048

② ⽣成公钥⽂件 pubkey.pem

Openssl-> rsa -in private_key.pem -pubout -out pubkey.pem

```

- 方式二：使用工具生成

请联系云账户技术支持获取 RSA 密钥生成工具

#### 上传平台企业公钥

登录云账户综合服务平台，在业务中心 -> 业务管理 -> 对接信息，点击页面右上角的编辑，配置平台企业公钥。
![配置平台企业公钥信息](.doc/publickeyconfig.png)


### 快速开始

#### 安装（使用go module）

```
go get github.com/YunzhanghuOpen/sdk-go
```


#### 快速使用

##### 功能列表

- [实时下单接口](api/payment.go)
- [数据接口](api/dataservice.go)
- [用户信息验证接口](api/authentication.go)
- [发票接口](api/invoice.go)
- [个税接口](api/tax.go)
- [异步回调](example/payment/payment.go)

##### 示例
```golang
package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/YunzhanghuOpen/sdk-go/api"
	"github.com/YunzhanghuOpen/sdk-go/core"
	"github.com/YunzhanghuOpen/sdk-go/errorx"
)

func main() {
	// 请求返回打印中间件
	logMiddle := func(next core.Handler) core.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			fmt.Println("middleware begin ", "requestID=", core.RequestID(ctx), " req=", req)
			resp, err := next(ctx, req)
			fmt.Println("middleware done  ", "requestID=", core.RequestID(ctx), " resp=", resp)
			return resp, err
		}
	}

	// requestID 注入中间件
	requestIDMiddle := func(next core.Handler) core.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return next(core.WithRequestID(ctx, fmt.Sprint(rand.Int63())), req)
		}
	}

	_ = requestIDMiddle
	dealerID := "需填写自己的商户ID"
	privateKey := "在云账户配置的公钥对应的私钥"
	appKey := "云账户对接信息中申请的 appKey"
	des3Key := "云账户对接信息中申请的 3DES Key"
	conf := &api.Config{
		Host:       api.ProductHost, // 正式环境域名，开发自测时可使用沙箱环境地址 api.SandboxHost
		DealerID:   dealerID,
		PrivateKey: privateKey,
		AppKey:     appKey,
		Des3Key:    des3Key,
	}

	client, err := api.New(conf,
		core.WithHttpClient(&http.Client{}),
		core.WithMiddleware(logMiddle),
		core.EnDebug(), // 测试时可开启，上线时可关闭
	)
	if err != nil {
		return
	}

	brokerID := "填写与当前商户签约的综合服务主体ID"
	orderID := time.Now().Format("20050102150405") // 订单号应由发号器生成，保证全局唯一，此处简写
	req := &api.CreateBankpayOrderRequest{
		BrokerID:  brokerID,
		DealerID:  dealerID,
		OrderID:   orderID,
		RealName:  "张三",
		IDCard:    "填写自己的",
		CardNo:    "填写自己的",
		PayRemark: "银行卡打款",
		NotifyURL: "https://wwww.callback.com", // 需填写自己实现的回调接口
		Pay:       "99.99",
	}

	requestID := fmt.Sprint(rand.Int63()) // 此处必需注入 requestID，便于发现问题时，进行问题排查，也可以使用 requestIDMiddle 中间件实现
	resp, err := client.CreateBankpayOrder(core.WithRequestID(context.TODO(), requestID), req)
	if err != nil {
		e, ok := errorx.FromError(err) // 解析配置
		if !ok {
			// 说明可能为sdk内部错误或网络错误，请求未到服务器
			// 也可能是服务端请求超时，需原单号重试
			return
		}
		if e.Code == "2002" {
			/*
				订单号重复
				检查是否已存在该订单号
			*/

			// TODO 异常处理流程

		} else {
			/*
				下单异常
				其它错误详见文档
				可参考 e.message 中的错误信息
			*/
		}
		return
	}
	_ = resp // 处理返回
}
```
