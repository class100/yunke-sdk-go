package yunke

const (
	// 文件类型
	// FileDirTypeProductFile 产品文件
	FileDirTypeProductFile DirType = "product"
	// FileDirTypeProductRelease 产品发布
	FileDirTypeProductRelease DirType = "product-release"
	// FileDirTypePublicDisk 公共云盘
	FileDirTypePublicDisk DirType = "public"
	// FileDirTypePrivateDisk 私有云盘
	FileDirTypePrivateDisk DirType = "private"
	// FileDirTypeFileResource 普通文件
	FileDirTypeFileResource DirType = "resource"
	// FileDirTypeSystemFile 系统文件文件
	FileDirTypeSystemFile DirType = "system"
	// FileDirTypeOrgRelease 版本发布文件
	FileDirTypeOrgRelease DirType = "org-release"
)

// 文件类型
type DirType string
