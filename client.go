package yunke

import (
	"encoding/json"
	`fmt`
	"strconv"

	"github.com/storezhang/gox"
)

const (
	// 客户端标识
	// ClientTypeWindows Windows客户端
	ClientTypeWindows ClientType = 1
	// ClientTypeMac Mac客户端
	ClientTypeMac ClientType = 2
	// ClientTypeAndroid 安卓客户端
	ClientTypeAndroid ClientType = 3

	// 更新类型
	// UpdateTypeTip 提示更新
	UpdateTypeTip UpdateType = 1
	// UpdateTypeSilent 静默更新
	UpdateTypeSilent UpdateType = 2
	// UpdateTypeForce 强制更新
	UpdateTypeForce UpdateType = 3

	// 更新文件类型
	// UpdateFileTypePatch 补丁包
	UpdateFileTypePatch UpdateFileType = 1
	// UpdateFileTypeInstaller 安装包
	UpdateFileTypeInstaller UpdateFileType = 2

	// 状态
	// ClientStatusPublished 已发布
	ClientStatusPublished ClientStatus = 10
	// ClientStatusNew 新创建
	ClientStatusNew ClientStatus = 11
	// ClientStatusPackaging 打包中
	ClientStatusPackaging ClientStatus = 21
	// ClientStatusPackaged 打包已完成
	ClientStatusPackaged ClientStatus = 22
	// ClientStatusPackageError 打包错误
	ClientStatusPackageError ClientStatus = 23
	// ClientStatusRepackaging 发布后再重新打包
	ClientStatusRepackaging ClientStatus = 24
)

type (
	// ClientType 客户端标识
	ClientType int8

	// UpdateType 更新类型
	UpdateType int8

	// UpdateFileType 更新文件类型
	UpdateFileType int8

	// ClientStatus 客户端状态
	ClientStatus int8

	// BaseClient 客户端版本
	BaseClient struct {
		gox.BaseStruct `xorm:"extends"`

		// ClientType 客户端
		ClientType ClientType `xorm:"tinyint default(1)" json:"clientType"`
		// UpdateType 更新类型
		UpdateType UpdateType `xorm:"tinyint default(1)" json:"updateType"`
		// FileType 文件类型
		FileType UpdateFileType `xorm:"tinyint default(1)" json:"fileType"`
		// Version 版本号
		Version string `xorm:"varchar(16) default('1.0.0')" json:"version"`
		// File 文件编号
		File string `xorm:"char(20) default('')" json:"file"`
		// Status 状态
		Status ClientStatus `xorm:"tinyint default(1)" json:"status"`
		// UpdateInfo 版本说明
		UpdateInfo string `xorm:"varchar(1024) default('')" json:"updateInfo"`
	}

	// BaseClientPaging 分页查询
	BaseClientPaging struct {
		gox.Paging

		// ClientType 客户端
		ClientType ClientType `json:"client" validate:"omitempty,oneof=1 2 3"`
		// SortField 排序字段
		SortField string `default:"updated_at" json:"sortField" validate:"oneof=id created_at updated_at version update_info"`
	}
)

func ParseClientType(ct string) (clientType ClientType, err error) {
	var iType int

	if iType, err = strconv.Atoi(ct); nil != err {
		return
	}
	clientType = ClientType(int8(iType))

	return
}

func (c BaseClient) Filename(name string) (filename string) {
	var ext string

	switch c.ClientType {
	case ClientTypeWindows:
		ext = ".exe"
	case ClientTypeMac:
		ext = ".dmg"
	case ClientTypeAndroid:
		ext = ".apk"
	default:
		ext = ".exe"
	}
	if UpdateFileTypePatch == c.FileType {
		ext = ".zip"
	}
	filename = fmt.Sprintf("%s-%s%s", name, c.Version, ext)

	return
}

func (c BaseClient) String() string {
	jsonBytes, _ := json.MarshalIndent(c, "", "    ")

	return string(jsonBytes)
}

func (cp *BaseClientPaging) SortFieldName() string {
	return cp.SortField
}
