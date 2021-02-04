package internet

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type HttpReq struct {
	Header  map[string]string
	Body    interface{}
	URL     string
	Method  string
	Timeout int
}

// 发送http请求
func SendHttpRequest(req *HttpReq, respData interface{}) (err error) {
	client := resty.New()

	httpReq := client.R().
		SetHeaders(req.Header).
		SetBody(req.Body).
		SetResult(&respData)

	var resp *resty.Response
	switch req.Method {
	case "GET":
		resp, err = httpReq.Get(req.URL)
	default:
		resp, err = httpReq.Post(req.URL)
	}
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("调取接口失败")
	}
	if resp.StatusCode() != 200 {
		return errors.New("调取第三方接口失败")
	}

	return nil
}
