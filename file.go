package yunke

type (
	// UploadCourseContentReq 获取上传url请求
	UploadFileReq struct {
		// FileId 编号
		FileId string `json:"fileId" validate:"omitempty,len=20"`
	}

	// FileUploadRsp 获取文件上传url响应
	FileUploadRsp struct {
		// Url 上传url
		Url string `json:"url"`
		// FileId 唯一编号id
		FileId string `json:"fileId"`
	}

	// DeleteFileReq 删除文件请求
	DeleteFileReq struct {
		// FileId 编号fileId
		FileId string `json:"fileId" validate:"omitempty,len=20"`
	}

	// GetDownloadReq 获取文件下载或者打开Url
	GetDownloadReq struct {
		// FileId 编号Id
		FileId string `json:"fileId" validate:"required,len=20"`
		// Type 下载类型
		// 1 立即下载
		// 2 打开
		Type int8 `default:"1" json:"type" validate:"required,oneof=1 2"`
		// Name 文件另存名字
		Name string `json:"name" validate:"omitempty,filename"`
	}

	// GetDownloadRsp 获取文件下载l响应
	GetDownloadRsp struct {
		// Url 上传url
		Url string `json:"url"`
	}
)
