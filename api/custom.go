package api

import "context"
import "strings"

// CustomService 通用请求函数
type CustomService interface {
	// DoRequest 执行请求
	// method 请求方式
	// urlStr 云账户接口地址，无需包含域名，示例：/api/payment/v1/order-bankpay
	// req 接口请求结构体，直接传入结构体即可，无需序列化，具体字段参见云账户接口文档
	// rsp 接口响应结构体，直接传入结构体即可，无需序列化，具体字段参见云账户接口文档
	DoRequest(ctx context.Context, method string, urlStr string, req interface{}) (rsp interface{}, err error)
}

type customServiceImpl struct {
	cc Invoker
}

// NewCustomService 创建客户端
func NewCustomService(cc Invoker) CustomService {
	return &customServiceImpl{cc}
}

func (c *customServiceImpl) DoRequest(ctx context.Context, method string, urlStr string, req interface{}) (rsp interface{}, err error) {
	rsp = new(interface{})
	err = c.cc.Invoke(ctx, strings.ToUpper(method), urlStr, false, req, &rsp)
	if err != nil {
		return nil, err
	}
	return
}
