package yunke

import (
	`encoding/json`

	`github.com/storezhang/gox`
	`github.com/storezhang/replace`
	`github.com/storezhang/transfer`
)

type (
	// AppConfig App配置
	AppConfig struct {
		// Name 应用名称
		Name string `json:"name,omitempty" validate:"required,max=10"`
		// Logo 应用图标
		Logo string `json:"logo,omitempty" validate:"omitempty,len=20"`
		// StartupLogo 启动图标
		StartupLogo string `json:"startupLogo,omitempty" validate:"omitempty,len=20"`
		// SplashLogo 闪屏
		SplashLogo string `json:"splashLogo,omitempty" validate:"omitempty,len=20"`
		// Packaged 是否已打包
		Packaged AppPackaged `json:"packaged,omitempty"`
	}

	// AppPackaged 移动端打包
	AppPackaged struct {
		gox.JSONInitialized

		// Android 安卓是否打包
		Android bool `json:"android"`
		// IOS IOS是否打包
		IOS bool `json:"ios"`
	}
)

func (ac AppConfig) IsInitialized() bool {
	return ac.Packaged.Initialized
}

func (ac AppConfig) InitSQL(table string, field string) (sql string, err error) {
	paths := make([]string, 0, 1)

	if !ac.Packaged.Initialized {
		paths = append(paths, "packaged")
	}
	sql, err = gox.MySQLJsonInit(
		table, field,
		"product", ProductApp,
		ac.Packaged.InitializeField(), true,
		paths...,
	)

	return
}

func (ac *AppConfig) StructToMap() (model map[string]interface{}, err error) {
	return gox.StructToMap(ac)
}

func (ac *AppConfig) MapToStruct(model map[string]interface{}) (err error) {
	return gox.MapToStruct(model, ac)
}

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

func (ac *AppConfig) Replaces(
	url string,
	packageName string,
	splash transfer.File,
) (replaces []replace.Replace, err error) {
	replaces = []replace.Replace{
		// 替换闪屏图片
		replace.NewFileReplace(DefaultAppSplashFilename, splash),
		// 替换包名
		replace.NewStringContentReplace(DefaultAndroidManifestFilename, "com.class100.yunke.dev", packageName),
		// 替换通信地址
		replace.NewJSONReplace(DefaultAppConfigFilename, replace.JSONReplaceElement{
			Path:  "server",
			Value: url,
		}),
	}

	return
}
