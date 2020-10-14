package yunke

import (
	`encoding/json`
	`fmt`
	`net/http`

	`github.com/dgrijalva/jwt-go`
	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

const (
	// OrgStatusCreated 已创建
	OrgStatusCreated OrgStatus = 1
	// OrgStatusInDeployment 部署中
	OrgStatusInDeployment OrgStatus = 2
	// OrgStatusInService 服务中
	OrgStatusInService OrgStatus = 3
	// OrgStatusDisabled 已禁用
	OrgStatusDisabled OrgStatus = 4
)

type (
	OrgStatus int8

	// Org 机构
	Org struct {
		gox.BaseStruct `xorm:"extends"`

		// 机构名称
		Name string `xorm:"varchar(64) notnull default('')" json:"name"`
		// 域名
		Url string `xorm:"varchar(64) notnull default('')" json:"url"`
		// 数据库配置
		DatabaseId int64 `xorm:"bigint notnull default(0)" json:"databaseId,string"`
		// 数据库名字
		Schema string `xorm:"varchar(32) notnull default('')" json:"schema"`
		// Redis配置
		RedisId int64 `xorm:"bigint notnull default(0)" json:"redisId,string"`
		// Redis Db序列号
		RedisDb uint8 `xorm:"tinyint notnull default(0)" json:"redisDb"`
		// 服务器密钥
		Secret string `xorm:"varchar(128) notnull default('')" json:"secret"`
		// 状态
		// 1 已创建
		// 2 部署中
		// 3 服务中
		// 4 已禁用
		Status OrgStatus `default:"1" xorm:"tinyint notnull index('idx_status') default(1)" json:"status"`
	}

	// OrgPaging 分页查询
	OrgPaging struct {
		gox.Paging

		// 状态
		Status int8 `default:"2" json:"status"`
		// 排序字段
		SortField string `default:"updated_at" json:"sortField" validate:"oneof=id created_at updated_at name domain"`
	}
)

func (o *Org) requestApi(
	path string,
	method gox.HttpMethod,
	params interface{}, pathParams map[string]string,
	version ApiVersion,
	rsp interface{},
) (err error) {
	return o.request(path, UrlApiPrefix, method, params, pathParams, version, rsp)
}

func (o *Org) request(
	path string,
	prefix string,
	method gox.HttpMethod,
	params interface{}, pathParams map[string]string,
	version ApiVersion,
	rsp interface{},
) (err error) {
	var (
		adminRsp           *resty.Response
		authToken          string
		expectedStatusCode int
	)

	if authToken, err = token(jwt.SigningMethodHS256, o.Secret); nil != err {
		return
	}

	req := NewResty().SetResult(rsp).SetHeader(gox.HeaderAuthorization, fmt.Sprintf("Bearer %s", authToken))
	// 注入路径参数
	if 0 != len(pathParams) {
		req = req.SetPathParams(pathParams)
	}

	// 修正请求地址为全路径
	orgConfig := OrgConfig{
		Url:  o.Url,
		Name: o.Name,
	}
	var url string
	if url, err = orgConfig.getUrl(path, pathParams, prefix, version); nil != err {
		return
	}

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

func (o Org) String() string {
	jsonBytes, _ := json.MarshalIndent(o, "", "    ")

	return string(jsonBytes)
}

func (op *OrgPaging) SortFieldName() string {
	return op.SortField
}
