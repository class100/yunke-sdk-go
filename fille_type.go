package yunke

const (
	// 文件类型
	// 产品文件
	FileDirTypeProductFile DirType = "product"
	// 产品发布
	FileDirTypeProductRelease DirType = "product-release"
	// 公共云盘
	FileDirTypePublicDisk DirType = "public"
	// 私有云盘
	FileDirTypePrivateDisk DirType = "private"
	// 普通文件
	FileDirTypeFileResource DirType = "resource"
	// 系统文件文件
	FileDirTypeSystemFile DirType = "system"
	// 课程资源
	FileDirTypeCourse = "course"
	// 版本发布文件
	FileDirTypeOrgRelease DirType = "org-release"
)

// 文件类型
type DirType string
