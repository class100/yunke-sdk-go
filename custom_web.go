package yunke

import (
	`encoding/json`

	`github.com/storezhang/gox`
)

type (
	// WebConfig 产品页面配置
	WebConfig struct {
		// Title 标题
		Title string `json:"title,omitempty" validate:"required,max=10"`
		// Dsp 描述
		Dsp string `json:"dsp,omitempty" validate:"omitempty,max=50"`
		// Logo 图标
		Logo string `json:"logo,omitempty" validate:"omitempty,len=20"`
		// Favicon 标签页图标
		Favicon string `json:"favicon,omitempty" validate:"omitempty,len=20"`

		// Copyright 版权信息
		Copyright string `json:"copyright,omitempty" validate:"omitempty,max=50"`
		// BeiAnHao 备案信息
		BeiAnHao string `json:"beiAnHao,omitempty" validate:"omitempty,max=50"`
		// BeiAnLink 备案链接
		BeiAnLink string `json:"beiAnLink,omitempty" validate:"omitempty,url"`
		// Contact 联系方式
		Contact string `json:"contact,omitempty" validate:"omitempty,max=100"`
		// DownloadPublicityMap 宣传图
		DownloadPublicityMap string `json:"downloadPublicityMap,omitempty" validate:"omitempty,len=20"`
	}
)

func (wc *WebConfig) StructToMap() (model map[string]interface{}, err error) {
	return gox.StructToMap(wc)
}

func (wc *WebConfig) MapToStruct(model map[string]interface{}) (err error) {
	return gox.MapToStruct(model, wc)
}

func (wc WebConfig) String() string {
	jsonBytes, _ := json.MarshalIndent(wc, "", "    ")

	return string(jsonBytes)
}
