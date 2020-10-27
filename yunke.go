package yunke

const (
	// 版本号
	// ApiVersionDefault 无版本，默认
	ApiVersionDefault ApiVersion = "default"
	// ApiVersionV1 版本1
	ApiVersionV1 ApiVersion = "v1"

	// UrlApiPrefix Api前缀
	UrlApiPrefix string = "api"
)

const (
	// OrgApiClientPackageNotifyUrl 客户端打包通知地址
	OrgApiClientPackageNotifyUrl string = "clients/{id}/packages/notifies"
	// OrgApiClientAddUrl 添加客户端地址
	OrgApiClientAddUrl string = "clients"

	// OrgApiConfigUpdateUrl 更新配置地址
	OrgApiConfigUpdateUrl string = "configs/{name}"
)

type (
	// ApiVersion 版本
	ApiVersion string
)
