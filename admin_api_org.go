package yunke

import (
	`strconv`

	`github.com/storezhang/gox`
)

// GetOrg 取得机构的信息
func (a *Admin) GetOrg() (org *Org, err error) {
		org = new(Org)
		err = a.request("orgs/{id}", gox.HttpMethodGet, nil, map[string]string{
			"id": strconv.FormatInt(a.Id, 10),
		}, org)

	return
}
