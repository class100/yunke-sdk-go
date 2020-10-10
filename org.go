package yunke

import (
	`encoding/json`

	`github.com/storezhang/gox`
)

const (
	// 已创建
	OrgStatusCreated OrgStatus = 1
	// 部署中
	OrgStatusInDeployment OrgStatus = 2
	// 服务中
	OrgStatusInService OrgStatus = 3
	// 已禁用
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

func (o Org) String() string {
	jsonBytes, _ := json.MarshalIndent(o, "", "    ")

	return string(jsonBytes)
}

func (op *OrgPaging) SortFieldName() string {
	return op.SortField
}
