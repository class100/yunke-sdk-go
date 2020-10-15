package yunke

import (
	`github.com/storezhang/gox`
)

// AddClient 添加客户端
func (o *Org) AddClient(client BaseClient, version ApiVersion) (c *BaseClient, err error) {
	c = new(BaseClient)
	orgClient := &struct {
		BaseClient

		// 原始未打包的文件编号
		OriginalFile string `json:"originalFile"`
	}{
		BaseClient:   client,
		OriginalFile: client.File,
	}
	// 清空原来的文件（防止提交的数据和机构数据定义有冲突）
	orgClient.File = ""

	err = o.requestApi(OrgApiClientAddUrl, gox.HttpMethodPost, orgClient, nil, version, c)

	return
}
