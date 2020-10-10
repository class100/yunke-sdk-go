package yunke

import (
	`fmt`
	libUrl `net/url`
	`strconv`
	`strings`
	`time`

	`github.com/storezhang/gox`
)

const (
	DefaultMaxOrgCacheDuration = 15 * time.Minute
)

var (
	cachedOrg     *Org  = nil
	orgCachedTime int64 = 0
)

// GetOrg 取得机构的信息
func (a *Admin) GetOrg() (org *Org, err error) {
	if nil == cachedOrg || DefaultMaxOrgCacheDuration < time.Now().Sub(time.Unix(orgCachedTime, 0)) {
		org = new(Org)
		err = a.request("orgs/{id}", gox.HttpMethodGet, nil, map[string]string{
			"id": strconv.FormatInt(a.Id, 10),
		}, org)

		cachedOrg = org
	}
	org = cachedOrg

	return
}

func (a *Admin) GetOrgDomain() (domain string, err error) {
	var (
		org *Org
		url *libUrl.URL
	)

	if org, err = a.GetOrg(); nil != err {
		return
	}
	if url, err = libUrl.Parse(org.Url); nil != err {
		return
	}

	domain = url.Hostname()

	return
}

func (a *Admin) GetOrgUrl() (url string, err error) {
	var org *Org

	if org, err = a.GetOrg(); nil != err {
		return
	}
	url = org.Url

	return
}

func (a *Admin) GetOrgName() (name string, err error) {
	var org *Org

	if org, err = a.GetOrg(); nil != err {
		return
	}
	name = org.Name

	return
}

func (a *Admin) GetOrgApiUrl(api string, pathParams map[string]string, version ApiVersion) (path string, err error) {
	return a.getOrgUrl(api, pathParams, UrlApiPrefix, version)
}

func (a *Admin) GetOrgPackageName() (packageName string, err error) {
	var domain string
	if domain, err = a.GetOrgDomain(); nil != err {
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
	packageName = strings.Join(hosts, ".")
	// 包名不能包含特殊字符，比如-、_、@、#等
	packageName = strings.ReplaceAll(packageName, "-", ".")
	packageName = strings.ReplaceAll(packageName, "_", ".")
	packageName = strings.ReplaceAll(packageName, "@", ".")
	packageName = strings.ReplaceAll(packageName, "#", ".")

	return
}

func (a *Admin) getOrgUrl(
	path string,
	pathParams map[string]string,
	prefix string,
	version ApiVersion,
) (url string, err error) {
	var (
		org *Org
		sb  strings.Builder
	)

	if org, err = a.GetOrg(); nil != err {
		return
	}

	if _, err = sb.WriteString(org.Url); nil != err {
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
	sb.WriteString(path)

	url = sb.String()
	// 处理路径参数
	if 0 < len(pathParams) {
		for param, value := range pathParams {
			url = strings.Replace(url, fmt.Sprintf("{%s}", param), libUrl.PathEscape(value), -1)
		}
	}

	return
}
