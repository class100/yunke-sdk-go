package yunke

import (
	`encoding/json`

	`github.com/storezhang/gox`
)

const (
	// DefaultOrgPackageName 默认包名
	DefaultOrgPackageName string = "com.class100.yunke"
	// DefaultLocalOrgPackageName 本地调试包名
	DefaultLocalOrgPackageName string = "com.localhost"
	// DefaultOrgUrl 默认连接地址
	DefaultOrgUrl string = "https://yunke.class100.com"
	// DefaultOrgDomain 默认域名
	DefaultOrgDomain string = "yunke.class100.com"
	// DefaultOrgName 默认机构名称
	DefaultOrgName string = "云视课堂"
	// DefaultLocalOrgName 本地调试机构名称
	DefaultLocalOrgName string = "本地开发环境"
)

const (
	// ConfigNamePwd 密码配置
	ConfigNamePwd ConfigName = "pwd"
	// ConfigNameUpgrade 升级配置
	ConfigNameUpgrade ConfigName = "upgrade"
	// ConfigNameOrg 机构配置
	ConfigNameOrg ConfigName = "org"
)

type (
	// ConfigName 配置名称
	ConfigName string

	// Config 配置
	Config struct {
		// Name 配置名称
		Name ConfigName `xorm:"pk varchar(16) notnull default('')" json:"name"`
		// Config 配置
		Config map[string]interface{} `xorm:"varchar(1024) notnull default('')" json:"config"`
		// UpdatedAt 更新时间
		UpdatedAt gox.Timestamp `xorm:"updated default('2020-06-11 09:55:52')" json:"updatedAt"`
	}
)

func (c Config) String() string {
	jsonBytes, _ := json.MarshalIndent(c, "", "    ")

	return string(jsonBytes)
}
