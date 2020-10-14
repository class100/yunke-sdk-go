package yunke

import (
	`github.com/storezhang/gox`
)

// UpdateConfig 更新配置
func (o *Org) UpdateConfig(name ConfigName, value interface{}, version ApiVersion) (config *Config, err error) {
	config = new(Config)
	err = o.requestApi(OrgApiConfigUpdateUrl, gox.HttpMethodPut, value, map[string]string{
		"name": string(name),
	}, version, config)

	return
}
