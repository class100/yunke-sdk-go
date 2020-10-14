package yunke

import (
	`encoding/json`
	`fmt`

	`github.com/class100/nuwa-sdk-go`
)

type (
	// PCConfig PC端配置
	PCConfig struct {
		// 标题
		Title string `json:"title" validate:"required,max=10"`
		// 图标
		Logo string `json:"logo" validate:"omitempty,len=20"`
		// 启动图标
		StartupLogo string `json:"startupLogo" validate:"omitempty,len=20"`
	}
)

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

func (pc *PCConfig) Replaces(url string, name map[string]string) (frs []nuwa.Replace, err error) {
	// 替换服务器通信地址
	elements := []nuwa.JSONReplaceElement{{
		Path:  "domain",
		Value: url,
	}}
	for lang, value := range name {
		// 替换语言
		elements = append(elements, nuwa.JSONReplaceElement{
			Path:  fmt.Sprintf("name.%s", lang),
			Value: value,
		})
	}

	frs = []nuwa.Replace{
		nuwa.NewJSONReplace(DefaultPCConfigFileName, elements...),
	}

	return
}
