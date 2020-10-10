package yunke

import (
	`fmt`
	`net/http`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

const (
	// 版本号
	// ApiVersionDefault 无版本，默认
	ApiVersionDefault ApiVersion = "default"
	// ApiVersionV1 版本1
	ApiVersionV1 ApiVersion = "v1"

	// UrlApiPrefix Api前缀
	UrlApiPrefix string = "api"
)

const (
	// ApiClientPackageNotifyUrl 客户端打包通知地址
	ApiClientPackageNotifyUrl string = "clients/{id}/packages/notifies"
)

type (
	// Admin 云视课堂
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

	// ApiVersion 版本
	ApiVersion string
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

	if authToken, err = token(a.SigningMethod, a.Secret); nil != err {
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
