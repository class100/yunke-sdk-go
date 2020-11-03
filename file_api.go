package yunke

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type fileHandler struct{}

func newFileHandler() *fileHandler {
	return &fileHandler{}
}

func (hdl *fileHandler) GetUploadInfo(req *UploadFileReq, host string) (rsp *FileUploadRsp, err error) {
	var (
		resp *resty.Response
		url  = fmt.Sprintf("%s/api/open/files/uploads/infos", host)
	)

	body := map[string]string{
		"fileId": req.FileId,
	}

	if resp, err = NewResty().SetQueryParams(body).Get(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = json.Unmarshal(resp.Body(), &rsp); nil != err {
		return
	}

	return
}

func (hdl *fileHandler) GetDownloadInfo(req *GetDownloadReq, host string) (downloadUrl string, err error) {
	var (
		resp *resty.Response
		url  = fmt.Sprintf("%s/api/open/files/downloads/%v", host, req.FileId)
	)

	body := map[string]string{
		"type": fmt.Sprintf("%v", req.Type),
		"name": req.Name,
	}

	if resp, err = NewResty().SetQueryParams(body).Get(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}
	downloadUrl = string(resp.Body())

	return
}

func (hdl *fileHandler) Delete(req *DeleteFileReq, host string) (err error) {
	var (
		url  = fmt.Sprintf("%s/api/open/files/%v", host, req.FileId)
		resp *resty.Response
	)

	if resp, err = NewResty().Delete(url); nil != err {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = getErr(resp)
		return
	}

	return
}
