package yunke

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type userHandler struct{}

func newUserHandler() *userHandler {
	return &userHandler{}
}

func (hdl *userHandler) Add(req *AddUserReq, host string) (course *User, err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/open/users", host)
	if resp, err = NewResty().SetBody(req).Post(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusCreated {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &course); nil != err {
		return
	}

	return
}

func (hdl *userHandler) Delete(id int64, host string) (err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/open/users/teachers/%v", host, id)
	if resp, err = NewResty().Delete(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusNoContent {
		err = getErr(resp)
		return
	}

	return
}

func (hdl *userHandler) Update(id int64, req map[string]interface{}, host string) (user *User, err error) {
	if 0 == len(req) {
		return
	}

	var resp *resty.Response
	url := fmt.Sprintf("%s/api/open/users/%v", host, id)

	if resp, err = NewResty().SetBody(req).Put(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &user); nil != err {
		return
	}

	return
}
