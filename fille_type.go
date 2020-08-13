package yunke

const (
	// 文件类型
	// 之所以要以10为间隔，是怕中途会半圆相类似的类型
	// 产品文件
	FileDirTypeProductFile DirType = "product"
	// 产品发布
	FileDirTypeProductRelease DirType = "release"
	// 公共云盘
	FileDirTypePublicDisk DirType = "public"
	// 私有云盘
	FileDirTypePrivateDisk DirType = "private"
	// 普通文件
	FileDirTypeFileResource DirType = "resource"
)

// 文件类型
type DirType string
