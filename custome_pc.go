package yunke

import (
	`encoding/json`
	`fmt`

	`github.com/storezhang/gox`
	`github.com/storezhang/replace`
)

type (
	// PCConfig PC端配置
	PCConfig struct {
		// Title 标题
		Title string `json:"title,omitempty" validate:"required,max=10"`
		// Logo 图标
		Logo string `json:"logo,omitempty" validate:"omitempty,len=20"`
		// StartupLogo 启动图标
		StartupLogo string `json:"startupLogo,omitempty" validate:"omitempty,len=20"`
		// Packaged 是否已打包
		Packaged PCPackaged `json:"package,omitempty"`
	}

	// PCPackaged PC端打包状态
	PCPackaged struct {
		gox.JSONInitialized

		// Windows Windows是否打包
		Windows bool `json:"windows"`
		// Mac Mac是否打包
		Mac bool `json:"mac"`
	}
)

func (pc PCConfig) IsInitialized() bool {
	return pc.Packaged.Initialized
}

func (pc PCConfig) InitSQL(table string, field string) (sql string, err error) {
	paths := make([]string, 0, 1)

	if !pc.Packaged.Initialized {
		paths = append(paths, "packaged")
	}
	sql, err = gox.MySQLJsonInit(table, field, pc.Packaged.InitializeField(), true, paths...)

	return
}

func (pc *PCConfig) StructToMap() (model map[string]interface{}, err error) {
	return gox.StructToMap(pc)
}

func (pc *PCConfig) MapToStruct(model map[string]interface{}) (err error) {
	return gox.MapToStruct(model, pc)
}

func (pc PCConfig) String() string {
	jsonBytes, _ := json.MarshalIndent(pc, "", "    ")

	return string(jsonBytes)
}

func (pc *PCConfig) isDiff(other PCConfig) (diff bool) {
	if "" == other.Title && "" == other.StartupLogo && "" == other.Logo {
		diff = true

		return
	}

	if pc.Title != other.Title {
		diff = true

		return
	}

	if pc.Logo != other.Logo {
		diff = true

		return
	}

	if pc.StartupLogo != other.StartupLogo {
		diff = true

		return
	}

	return
}

func (pc *PCConfig) Replaces(url string, name map[string]string) (replaces []replace.Replace, err error) {
	// 替换服务器通信地址
	elements := []replace.JSONReplaceElement{{
		Path:  "domain",
		Value: url,
	}}
	for lang, value := range name {
		// 替换语言
		elements = append(elements, replace.JSONReplaceElement{
			Path:  fmt.Sprintf("name.%s", lang),
			Value: value,
		})
	}

	replaces = []replace.Replace{
		replace.NewJSONReplace(DefaultPCConfigFileName, elements...),
	}

	return
}
