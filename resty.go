package yunke

import (
	`crypto/tls`
	`encoding/json`
	`fmt`

	`github.com/go-resty/resty/v2`
)

type (
	response struct {
		ErrorCode int         `json:"errorCode"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
	}
)

// NewResty Resty客户端
func NewResty() *resty.Request {
	return resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).R()
}

// RestyStringBody 字符串形式的结果
func RestyStringBody(rsp *resty.Response) string {
	body := ""
	if nil != rsp {
		body = rsp.String()
	}

	return body
}

func getErr(resp *resty.Response) (err error) {
	var rsp *response
	if err = json.Unmarshal(resp.Body(), &rsp); nil != err {
		return
	}
	err = fmt.Errorf(rsp.Message)

	return
}
