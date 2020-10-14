package yunke

import (
	`github.com/storezhang/gox`
)

// AddClient 添加客户端
func (o *Org) AddClient(client interface{}, version ApiVersion) (c *BaseClient, err error) {
	c = new(BaseClient)
	err = o.requestApi(OrgApiClientAddUrl, gox.HttpMethodPost, client, nil, version, c)

	return
}
