package yunke

import `github.com/storezhang/gox`

const (
	// CourseTeachingTypeMiniClass 1：小班课
	CourseTypeMiniClass CourseTeachingType = 1
	// CourseTypeOneToOne 2：一对一
	CourseTeachingTypeOneToOne CourseTeachingType = 2
	// CourseTypeRecord 3：录播课
	CourseTeachingTypeRecord CourseTeachingType = 3
	// CourseTypePublic 4：公开课
	CourseTeachingTypePublic CourseTeachingType = 4

	// ClassModelOriginal 0：原生模式
	ClassModelOriginal ClassMode = 0
	// ClassModelCustom 1：定制的
	ClassModelCustom ClassMode = 1
)

type (
	// CourseTeachingType 课程类型
	CourseTeachingType int8

	// ClassMode 上课模式
	ClassMode int8

	// Course 课程
	Course struct {
		gox.BaseStruct `xorm:"extends"`

		// Name 课程名称
		Name string `xorm:"varchar(32) notnull default('')" json:"name" validate:"required,min=1,max=32"`
		// Type 课程类型
		// 1：小班课
		// 2：一对一
		// 3：录播课
		// 4：公开课
		Type CourseTeachingType `default:"1" xorm:"tinyint notnull default(1)" json:"type" validate:"required"`
		// CategoryId 课程分类
		CategoryId int64 `default:"1" xorm:"tinyint notnull default(1)" json:"categoryId,string" validate:"required"`
		// Cover 封面
		Cover string `xorm:"char(20) notnull default('')" json:"cover" validate:"omitempty,len=20"`
		// CreatorId 创建人
		CreatorId int64 `xorm:"bigint(20) notnull default(1)" json:"creatorId,string" validate:"required"`
		// MaxNum 教室容量
		MaxNum int `xorm:"int notnull default(1)" json:"maxNum" validate:"required"`
		// Info 介绍
		Info string `xorm:"text(10000) notnull default('')" json:"info" validate:"omitempty,max=10000"`
		// ResourcePath 课程资源关联路径
		// 例如/a/b/c
		ResourcePath string `xorm:"varchar(255)" json:"resourcePath"`
		// ClassMode 上课模式
		// 0：原生类型
		// 1：自定义类型
		ClassMode ClassMode `xorm:"tinyint default(1)" json:"classMode"`
	}

	// AddCourseReq 添加课程
	AddCourseReq struct {
		// Name 课程名称
		Name string `json:"name" validate:"required,without_special_symbol,min=2,max=30"`
		// Type 课程类型
		// 1 小班课
		// 2 一对一
		// 3 录播课
		// 4 公开课
		Type CourseTeachingType `json:"type" validate:"required_with=1 2 3 4"`
		// CategoryId 课程分类
		CategoryId int64 `json:"categoryId,string" validate:"required"`
		// Cover 封面
		Cover string `json:"cover" validate:"omitempty,len=20"`
		// CreatorId 创建人
		CreatorId int64 `json:"creatorId,string" validate:"required"`
		// MaxNum 教室容量
		MaxNum int `json:"maxNum" validate:"omitempty"`
		// Info 介绍
		Info string `json:"info" validate:"omitempty,max=10000"`
		// TeacherIds 讲师编号列表
		TeacherIds gox.Int64Slice `json:"teacherIds"`
		// ResourcePath 课程资源关联路径
		// 例如 /a/b/c
		ResourcePath string `json:"resourcePath" validate:"omitempty,startswith=/"`
		// ClassMode 上课模式
		// 0：原生类型
		// 1：自定义类型
		ClassMode ClassMode `default:"0" json:"classMode" validate:"omitempty,oneof=0 1"`
	}

	// UpdateCourseReq 更新课程请求
	UpdateCourseReq struct {
		gox.BaseStruct

		// Name 课程名称
		Name string `json:"name" validate:"required,without_special_symbol,min=2,max=30"`
		// Type 课程类型
		// 1 小班课
		// 2 一对一
		// 3 录播课
		// 4 公开课
		Type CourseTeachingType `json:"type"`
		// CategoryId 课程分类
		CategoryId int64 `json:"categoryId,string"`
		// Cover 封面
		Cover string `json:"cover" validate:"omitempty,len=20"`
		// MaxNum 教室容量
		MaxNum int `json:"maxNum"`
		// Info 介绍
		Info string `json:"info" validate:"omitempty,max=10000"`
		// TeacherIds 讲师编号列表
		TeacherIds gox.Int64Slice `json:"teacherIds"`
		// ResourcePath 课程资源关联路径
		// 例如/a/b/c
		ResourcePath string `json:"resourcePath" validate:"omitempty,startswith=/"`
		// ClassMode 上课模式
		// 0：原生类型
		// 1：自定义类型
		ClassMode ClassMode `default:"0" json:"classMode" validate:"omitempty,oneof=0 1"`
	}
)
