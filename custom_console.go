package yunke

import (
	`encoding/json`

	`github.com/storezhang/gox`
)

type (
	// ConsoleConfig 管理后台配置
	ConsoleConfig struct {
		// Title 标题
		Title string `json:"title,omitempty" validate:"required,max=10"`
		// Dsp 描述
		Dsp string `json:"dsp,omitempty" validate:"omitempty,max=50"`
		// NavBarLogo 后台导航栏Logo
		NavBarLogo string `json:"navBarLogo,omitempty" validate:"omitempty,len=20"`
		// LoginLogo 后台登录页Logo
		LoginLogo string `json:"loginLogo,omitempty" validate:"omitempty,len=20"`
		// BackgroundImage 后台登录页背景图
		BackgroundImage string `json:"backgroundImage,omitempty" validate:"omitempty,len=20"`
		// Favicon 标签页图标
		Favicon string `json:"favicon,omitempty" validate:"omitempty,len=20"`
	}
)

func (cc *ConsoleConfig) StructToMap() (model map[string]interface{}, err error) {
	return gox.StructToMap(cc)
}

func (cc *ConsoleConfig) MapToStruct(model map[string]interface{}) (err error) {
	return gox.MapToStruct(model, cc)
}

func (cc ConsoleConfig) String() string {
	jsonBytes, _ := json.MarshalIndent(cc, "", "    ")

	return string(jsonBytes)
}
