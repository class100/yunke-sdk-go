package yunke

import (
	`fmt`
	`net/http`

	`github.com/dgrijalva/jwt-go`
	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

type (
	// Admin 云视课堂管理
	Admin struct {
		// 产品编号
		Id int64
		// 地址
		Url string
		// 通信秘钥
		Secret string
		// 加密方法
		SigningMethod string
		// 授权码，在Header里面
		AuthScheme string
	}
)

func (a *Admin) request(
	url string,
	method gox.HttpMethod,
	params interface{}, pathParams map[string]string,
	rsp interface{},
) (err error) {
	var (
		adminRsp           *resty.Response
		authToken          string
		expectedStatusCode int
	)

	if authToken, err = token(jwt.GetSigningMethod(a.SigningMethod), a.Secret); nil != err {
		return
	}

	req := NewResty().SetResult(rsp).SetHeader(gox.HeaderAuthorization, fmt.Sprintf("%s %s", a.AuthScheme, authToken))
	// 注入路径参数
	if 0 != len(pathParams) {
		req = req.SetPathParams(pathParams)
	}

	// 修正请求地址为全路径
	url = fmt.Sprintf("%s/api/%s", a.Url, url)

	switch method {
	case gox.HttpMethodGet:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetQueryParams(params.(map[string]string))
		}
		adminRsp, err = req.Get(url)
	case gox.HttpMethodPost:
		expectedStatusCode = http.StatusCreated

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Post(url)
	case gox.HttpMethodPut:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Put(url)
	case gox.HttpMethodDelete:
		expectedStatusCode = http.StatusNoContent

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Delete(url)
	}
	if nil != err {
		return
	}

	if nil == adminRsp {
		err = gox.NewCodeError(gox.ErrorCode(adminRsp.StatusCode()), "无返回数据", RestyStringBody(adminRsp))

		return
	}

	// 检查状态码
	if expectedStatusCode != adminRsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(adminRsp.StatusCode()), "请求服务器不符合预期", RestyStringBody(adminRsp))
	}

	return
}
