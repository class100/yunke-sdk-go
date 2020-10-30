package yunke

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type lecture struct{}

func newLecture() *lecture {
	return &lecture{}
}

func (hdl *lecture) Add(req *AddLectureReq, host string) (lecture *LectureInfo, err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/lectures", host)
	if resp, err = NewResty().SetBody(req).Post(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusCreated {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &lecture); nil != err {
		return
	}

	return
}

func (hdl *lecture) Delete(id int64, host string) (err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/lectures/%v", host, id)
	if resp, err = NewResty().Delete(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusNoContent {
		err = getErr(resp)
		return
	}

	return
}

func (hdl *lecture) Update(id int64, req map[string]interface{}, host string) (lecture *LectureInfo, err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/lectures/%v", host, id)
	if resp, err = NewResty().SetBody(req).Put(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	lecture = new(LectureInfo)
	if err = json.Unmarshal(resp.Body(), &lecture); nil != err {
		return
	}

	return
}

func (hdl *lecture) Get(id int64, host string) (lecture *LectureInfo, err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/lectures/%v", host, id)
	if resp, err = NewResty().Get(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &lecture); nil != err {
		return
	}

	return
}

func (hdl *lecture) Gets(courseId int64, host string) (chapters []*ChapterInfo, err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/lectures/courses/%v", host, courseId)
	if resp, err = NewResty().Get(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &chapters); nil != err {
		return
	}

	return
}

func (hdl *lecture) SwitchSequence(req *SwitchSequenceReq, host string) (lectures []*LectureInfo, err error) {
	var (
		resp *resty.Response
	)

	url := fmt.Sprintf("%s/api/class330/lectures/switches", host)
	if resp, err = NewResty().SetBody(req).Post(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &lectures); nil != err {
		return
	}

	return
}
