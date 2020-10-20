package yunke

import (
	`github.com/json-iterator/go`
)

func toModel(obj interface{}) (model map[string]interface{}, err error) {
	var (
		json      = jsoniter.ConfigCompatibleWithStandardLibrary
		jsonBytes []byte
	)

	if jsonBytes, err = json.Marshal(obj); nil != err {
		return
	}

	err = json.Unmarshal(jsonBytes, &model)

	return
}
