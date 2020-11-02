package yunke

import (
	`encoding/json`
	`fmt`
	`github.com/go-resty/resty/v2`
	`net/http`
)

type course struct{}

func newCourse() *course {
	return &course{}
}

func (hdl *course) Add(req *AddCourseReq, host string) (course *Course, err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/courses", host)
	if resp, err = NewResty().SetBody(req).Post(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &course); nil != err {
		return
	}

	return
}

func (hdl *course) Delete(id int64, host string) (err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/courses/%v", host, id)
	if resp, err = NewResty().Delete(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	return
}

func (hdl *course) Update(id int64, req map[string]interface{}, host string) (course Course, err error) {
	if 0 == len(req) {
		return
	}

	var resp *resty.Response
	url := fmt.Sprintf("%s/api/class330/courses/%v", host, id)

	if resp, err = NewResty().SetBody(req).Put(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &course); nil != err {
		return
	}

	return
}
