package yunke

import (
	`encoding/json`
)

type (
	// WebConfig 产品页面配置
	WebConfig struct {
		// 标题
		Title string `json:"title" validate:"required,max=10"`
		// 描述
		Dsp string `json:"dsp" validate:"omitempty,max=50"`
		// 图标
		Logo string `json:"logo" validate:"omitempty,len=20"`
		// 标签页图标
		Favicon string `json:"favicon" validate:"omitempty,len=20"`

		// 版权信息
		Copyright string `json:"copyright" validate:"omitempty,max=50"`
		// 备案信息
		BeiAnHao string `json:"beiAnHao" validate:"omitempty,max=50"`
		// 备案链接
		BeiAnLink string `json:"beiAnLink" validate:"omitempty,url"`
		// 联系方式
		Contact string `json:"contact" validate:"omitempty,max=100"`
		// 宣传图
		DownloadPublicityMap string `json:"downloadPublicityMap" validate:"omitempty,len=20"`
	}
)

func (wc WebConfig) String() string {
	jsonBytes, _ := json.MarshalIndent(wc, "", "    ")

	return string(jsonBytes)
}
