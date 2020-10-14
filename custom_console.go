package yunke

import (
	`encoding/json`
)

type (
	// ConsoleConfig 管理后台配置
	ConsoleConfig struct {
		// 标题
		Title string `json:"title" validate:"required,max=10"`
		// 描述
		Dsp string `json:"dsp" validate:"omitempty,max=50"`
		// 后台导航栏Logo
		NavBarLogo string `json:"navBarLogo" validate:"omitempty,len=20"`
		// 后台登录页Logo
		LoginLogo string `json:"loginLogo" validate:"omitempty,len=20"`
		// 后台登录页背景图
		BackgroundImage string `json:"backgroundImage" validate:"omitempty,len=20"`
		// 标签页图标
		Favicon string `json:"favicon" validate:"omitempty,len=20"`
	}
)

func (cc ConsoleConfig) String() string {
	jsonBytes, _ := json.MarshalIndent(cc, "", "    ")

	return string(jsonBytes)
}
