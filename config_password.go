package yunke

import (
	`github.com/storezhang/gox`
)

type (
	// PasswordConfig 角色默认密码配置
	PasswordConfig struct {
		// 密码消耗
		Cost int `json:"cost" validate:"omitempty,min=4,max=31"`
		// 默认密码
		Default DefaultPasswordConfig `json:"default"`
	}

	// DefaultPasswordConfig 默认密码配置
	DefaultPasswordConfig struct {
		gox.JSONInitialized

		// 讲师、助教默认密码
		Lecturer string `json:"lecturer" validate:"required,min=8,max=30"`
		// 学生默认密码
		Student string `json:"student" validate:"required,min=8,max=30"`
		// 后台除了管理员的默认密码
		Backend string `json:"backend" validate:"required,min=8,max=30"`
	}
)

func (pc PasswordConfig) IsInitialized() bool {
	return pc.Default.Initialized
}

func (pc PasswordConfig) InitSQL(table string, field string) (sql string, err error) {
	paths := make([]string, 0, 1)

	if !pc.Default.Initialized {
		paths = append(paths, "default")
	}
	sql, err = gox.MySQLJsonInit(table, field, pc.Default.InitializeField(), pc.Default.IsInitialized(), paths...)

	return
}

func (pc *PasswordConfig) StructToMap() (model map[string]interface{}, err error) {
	return gox.StructToMap(pc)
}

func (pc *PasswordConfig) MapToStruct(model map[string]interface{}) (err error) {
	return gox.MapToStruct(model, pc)
}
