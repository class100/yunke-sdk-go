package yunke

import (
	`encoding/json`

	`github.com/class100/nuwa-sdk-go`
)

type (
	// AppConfig App配置
	AppConfig struct {
		// 应用名称
		Name string `json:"name" validate:"required,max=10"`
		// 应用图标
		Logo string `json:"logo" validate:"omitempty,len=20"`
		// 启动图标
		StartupLogo string `json:"startupLogo" validate:"omitempty,len=20"`
		// 闪屏
		SplashLogo string `json:"splashLogo" validate:"omitempty,len=20"`
	}
)

func (ac AppConfig) String() string {
	jsonBytes, _ := json.MarshalIndent(ac, "", "    ")

	return string(jsonBytes)
}

func (ac *AppConfig) isDiff(newConfig AppConfig) (diff bool) {
	if "" == newConfig.Name && "" == newConfig.StartupLogo {
		diff = true

		return
	}

	if ac.Name != newConfig.Name {
		diff = true

		return
	}

	if ac.StartupLogo != newConfig.StartupLogo {
		diff = true

		return
	}

	return
}

func (ac *AppConfig) Replaces(url string, packageName string, splash nuwa.File) (frs []nuwa.Replace, err error) {
	frs = []nuwa.Replace{
		// 替换闪屏图片
		nuwa.NewFileReplace(DefaultAppSplashFileName, splash),
		// 替换包名
		nuwa.NewStringContentReplace(DefaultAndroidManifestFileName, "com.class100.yunke.dev", packageName),
		// 替换通信地址
		nuwa.NewJSONReplace(DefaultAppConfigFileName, nuwa.JSONReplaceElement{
			Path:  "server",
			Value: url,
		}),
	}

	return
}
