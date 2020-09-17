package yunke

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storezhang/gox"
)

type (
	// Admin 云视课堂
	Admin struct {
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

// GetLastVersion 取得某个客户端类型的最新版本
func (a *Admin) GetLastVersion(clientType ClientType) (baseClient BaseClient, err error) {
	var (
		rsp       *resty.Response
		authToken string
	)

	if authToken, err = token(a.SigningMethod, a.Secret); nil != err {
		return
	}

	if rsp, err = NewResty().SetHeader("Authorization", fmt.Sprintf("%s %s", a.AuthScheme, authToken)).
		SetResult(&baseClient).
		Get(fmt.Sprintf("%s/api/clients/finals/%d", a.Url, clientType)); nil != err {
		log.WithFields(log.Fields{
			"url":        a.Url,
			"secret":     a.Secret,
			"clientType": clientType,
			"body":       RestyStringBody(rsp),
			"error":      err,
		}).Error("取得最新版本出错")

		return
	}

	if http.StatusOK != rsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(rsp.StatusCode()), "取得最新版本出错", rsp.String())

		return
	}
	log.WithFields(log.Fields{
		"url":        a.Url,
		"secret":     a.Secret,
		"clientType": clientType,
	}).Debug("取得最新版本成功")

	return
}
