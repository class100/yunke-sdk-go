package yunke

import (
	`fmt`
	libUrl `net/url`
	`strings`
)

type (
	// OrgConfig 机构配置
	OrgConfig struct {
		// Url 机构地址
		Url string `json:"url" mapstructure:"url" validate:"required"`
		// Name 机构名称
		Name string `json:"name" mapstructure:"name" validate:"required"`
	}
)

func (oc OrgConfig) Model() (map[string]interface{}, error) {
	return toModel(oc)
}

func (oc *OrgConfig) Exist() bool {
	return "" != oc.Url && "" != oc.Name
}

func (oc *OrgConfig) Domain() (domain string, err error) {
	var url *libUrl.URL

	if url, err = libUrl.Parse(oc.Url); nil != err {
		return
	}

	domain = url.Hostname()

	return
}

func (oc *OrgConfig) PackageName() (name string, err error) {
	var domain string
	if domain, err = oc.Domain(); nil != err {
		return
	}

	hosts := strings.Split(domain, ".")
	// 包名必须包含多个单词，如果只有一个单词，默认增加一个com后缀
	if 1 == len(hosts) {
		hosts = append(hosts, "com")
	}

	// 反转，和包名规范一致
	for i, j := 0, len(hosts)-1; i < j; i, j = i+1, j-1 {
		hosts[i], hosts[j] = hosts[j], hosts[i]
	}
	name = strings.Join(hosts, ".")
	// 包名不能包含特殊字符，比如-、_、@、#等
	name = strings.ReplaceAll(name, "-", ".")
	name = strings.ReplaceAll(name, "_", ".")
	name = strings.ReplaceAll(name, "@", ".")
	name = strings.ReplaceAll(name, "#", ".")

	return
}

func (oc *OrgConfig) ConfigUpdateUrl(name ConfigName) (path string, err error) {
	return oc.ApiUrl(OrgApiConfigUpdateUrl, map[string]string{"name": string(name)}, ApiVersionDefault)
}

func (oc *OrgConfig) PackageNotifyUrl(client BaseClient) (path string, err error) {
	return oc.ApiUrl(OrgApiClientPackageNotifyUrl, map[string]string{"id": client.IdString()}, ApiVersionDefault)
}

func (oc *OrgConfig) ApiUrl(
	api string,
	pathParams map[string]string,
	version ApiVersion,
) (path string, err error) {
	return oc.getUrl(api, pathParams, UrlApiPrefix, version)
}

func (oc *OrgConfig) getUrl(
	path string,
	pathParams map[string]string,
	prefix string,
	version ApiVersion,
) (url string, err error) {
	var sb strings.Builder

	if _, err = sb.WriteString(oc.Url); nil != err {
		return
	}

	if "" != prefix {
		if _, err = sb.WriteString(fmt.Sprintf("/%s", prefix)); nil != err {
			return
		}
	}

	if ApiVersionDefault != version {
		if _, err = sb.WriteString(fmt.Sprintf("/%s", version)); nil != err {
			return
		}
	}
	sb.WriteString(fmt.Sprintf("/%s", path))

	url = sb.String()
	// 处理路径参数
	if 0 < len(pathParams) {
		for param, value := range pathParams {
			url = strings.Replace(url, fmt.Sprintf("{%s}", param), libUrl.PathEscape(value), -1)
		}
	}

	return
}
