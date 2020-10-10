package yunke

import (
	`strconv`

	`github.com/storezhang/gox`
)

// GetLastVersion 取得某个客户端类型的最新版本
func (a *Admin) GetLastVersion(clientType ClientType) (client *BaseClient, err error) {
	client = new(BaseClient)
	err = a.request("clients/finals/{clientType}", gox.HttpMethodGet, nil, map[string]string{
		"clientType": strconv.Itoa(int(clientType)),
	}, client)

	return
}
