package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/YunzhanghuOpen/sdk-go/errorx"
)

// Invoke 实现invoker接口
func (o *Core) Invoke(ctx context.Context, method string, urlStr string, respEncrypted bool, req interface{}, resp interface{}) error {
	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		if RequestID(ctx) == "" {
			// 必须注入 requestID，方便调用链追踪，问题排查
			return nil, errors.New("please use function WithRequestID inject request_id")
		}
		err := o.invoke(ctx, method, urlStr, respEncrypted, req, resp)
		return resp, err
	}

	if len(o.middlewares) > 0 {
		h = Chain(o.middlewares...)(h)
	}
	_, err := h(ctx, req)
	return err
}

func (o *Core) invoke(ctx context.Context, method string, urlStr string, respEncrypted bool, req interface{}, resp interface{}) error {
	if respEncrypted {
		return o.doForEncryptResp(ctx, method, urlStr, req, resp)
	}
	return o.do(ctx, method, urlStr, req, resp)
}

func (o *Core) doForEncryptResp(ctx context.Context, method string, urlStr string, req interface{}, resp interface{}) error {
	b, err := o.req(ctx, method, urlStr, req)
	if err != nil {
		return err
	}

	if o.debug {
		o.logger.Logf(ctx, "req=%+v, resp=%s", req, string(b))
	}

	r := struct {
		Code      string `json:"code"`       // 响应码
		Message   string `json:"message"`    // 响应信息
		RequestID string `json:"request_id"` // 请求ID
		Data      string `json:"data"`       // 响应体
	}{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		o.logger.Logf(ctx, "json.Unmarshal failed, err=%v, body=%s", err, string(b))
		return err
	}

	if r.Code != errorx.OK {
		return errorx.New(r.Code, r.Message)
	}

	b, err = o.encoder.Decode([]byte(r.Data))
	if err != nil {
		o.logger.Logf(ctx, "Decode failed, err=%v, body=%s", err, r.Data)
		return err
	}

	err = json.Unmarshal(b, resp)
	if err != nil {
		o.logger.Logf(ctx, "json.Unmarshal failed, err=%v, body=%s", err, string(b))
		return err
	}
	return nil
}

func (o *Core) do(ctx context.Context, method string, urlStr string, req interface{}, resp interface{}) error {
	b, err := o.req(ctx, method, urlStr, req)
	if err != nil {
		return err
	}

	if o.debug {
		o.logger.Logf(ctx, "req=%+v, resp=%s", req, string(b))
	}

	r := struct {
		Code      string          `json:"code"`       // 响应码
		Message   string          `json:"message"`    // 响应信息
		RequestID string          `json:"request_id"` // 请求ID
		Data      json.RawMessage `json:"data"`       // 响应体
	}{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		o.logger.Logf(ctx, "json.Unmarshal failed, err=%v, body=%s", err, string(b))
		return err
	}

	if r.Code != errorx.OK {
		return errorx.New(r.Code, r.Message)
	}

	if len(r.Data) == 0 {
		// 返回可能没有 data 字段
		return nil
	}

	err = json.Unmarshal(r.Data, resp)
	if err != nil {
		o.logger.Logf(ctx, "json.Unmarshal failed, err=%v, data=%s", err, string(r.Data))
		return err
	}
	return nil
}

func (o *Core) encodeData(ctx context.Context, data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		o.logger.Logf(ctx, "json.Marshal failed, err=%v", err)
		return nil, err
	}

	b, err = o.encoder.Encode(b)
	if err != nil {
		o.logger.Logf(ctx, "Encode failed, err=%v", err)
		return nil, err
	}
	return b, nil
}

func (o *Core) req(ctx context.Context, method string, urlStr string, req interface{}) ([]byte, error) {
	r, err := o.buildReq(ctx, method, urlStr, req)
	if err != nil {
		return nil, err
	}

	reply, err := o.httpClient.Do(r)
	if err != nil {
		o.logger.Logf(ctx, "http.Do failed, err=%v", err)
		return nil, err
	}
	defer reply.Body.Close()

	if reply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong http status_code=%s", reply.Status)
	}
	return ioutil.ReadAll(reply.Body) // 为了打印body，使用ioutil.ReadAll
}

func (o *Core) buildReq(ctx context.Context, method string, urlStr string, req interface{}) (*http.Request, error) {
	b, err := o.encodeData(ctx, req)
	if err != nil {
		return nil, err
	}

	data := string(b)
	mess := fmt.Sprint(time.Now().Nanosecond())
	timestamp := fmt.Sprint(time.Now().Unix())
	sign, err := o.signer.Sign(mess, timestamp, data)
	if err != nil {
		o.logger.Logf(ctx, "@buildReq Sign failed, err=%v", err)
		return nil, err
	}

	m := map[string]string{
		"data":      data,
		"mess":      mess,
		"timestamp": timestamp,
		"sign":      sign,
		"sign_type": o.signType,
	}

	values := url.Values{}
	for k, v := range m {
		values.Set(k, v)
	}

	body := values.Encode()
	urlStr = o.host + urlStr
	if method == http.MethodGet {
		urlStr = fmt.Sprintf("%s?%s", urlStr, values.Encode())
		body = ""
	}

	requestID := RequestID(ctx)
	header := map[string]string{
		"dealer-id":    o.dealerID,
		"request-id":   requestID,
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   userAgent(),
	}

	if o.debug {
		o.logger.Logf(ctx, "request-id=%s, url=%s, data=%+v, header=%+v", requestID, urlStr, m, header)
	}

	r, err := http.NewRequestWithContext(ctx, method, urlStr, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		r.Header.Set(k, v)

	}
	return r, nil
}
