package yunke

import "github.com/storezhang/gox"

type (
	// LectureContent 课程信息
	Lecture struct {
		gox.BaseStruct `xorm:"extends"`

		// CourseId 课程编号
		CourseId int64 `xorm:"bigint notnull default(0)" json:"courseId,string"`
		// Type 讲次类型
		// 0：章节
		// 1：讲次
		Type int8 `default:"0" xorm:"tinyint notnull default(0)" json:"type"`
		// ParentId 讲次的章节编号
		ParentId int64 `xorm:"bigint notnull default(0)" json:"parentId,string"`
		// Sequence 讲次顺序
		Sequence int8 `default:"0" xorm:"tinyint notnull default(0)" json:"sequence"`
		// Name 课程名称
		Name string `xorm:"varchar(32) notnull default('')" json:"name"`
		// Profile 课程简介
		Profile string `xorm:"varchar(255) notnull default('')" json:"profile"`
	}

	// LectureContent 讲次内容
	LectureContent struct {
		gox.BaseStruct `xorm:"extends"`

		// LectureId 课程Id
		LectureId int64 `xorm:"bigint notnull default(0)" json:"lectureId,string"`
		// FileId 展示图片文件编号Id
		FileId string `xorm:"char(20) notnull default(0)" json:"fileId"`
		// FileName 名称
		FileName string `xorm:"varchar(32) notnull default('')" json:"fileName"`
		// 内容类型：0-讲次视频，1-讲次资料
		Type int8 `default:"0" xorm:"tinyint notnull default(0)" json:"type"`
	}

	// AddLectureReq 添加讲次请求
	AddLectureReq struct {
		// 课程Id
		CourseId int64 `json:"courseId,string" validate:"required"`
		// Type 讲次类型
		// 0：章节
		// 1：讲次
		Type int8 `default:"0" json:"type" validate:"omitempty,oneof=0 1"`
		// ParentId 章节Id
		ParentId int64 `json:"parentId,string"`
		// Sequence 顺序
		Sequence int8 `default:"0" json:"sequence"`
		// Name 名称
		Name string `json:"name" validate:"required"`
		// Profile 简介
		Profile string `json:"profile" validate:"omitempty,max=255"`
		// LectureContents 讲次内容
		LectureContents []*LectureContent `json:"lectureContents" validate:"omitempty"`
	}

	// UpdateLectureReq 更新讲次请求
	UpdateLectureReq struct {
		gox.BaseStruct

		// CourseId 课程Id
		CourseId int64 `json:"courseId,string"`
		// Type 讲次类型
		// 0：章节
		// 1：讲次
		Type int8 `json:"Type"`
		// ParentId 父级Id
		ParentId int64 `json:"parentId,string"`
		// Sequence 顺序
		Sequence int8 `json:"sequence"`
		// Name 名称
		Name string `json:"name"`
		// Profile 简介
		Profile string `json:"profile" validate:"omitempty,max=255"`
		// LectureContents 讲次内容
		LectureContents []*LectureContent `json:"lectureContents"`
	}

	// SwitchItem 交换项
	SwitchItem struct {
		gox.IdStruct

		// ParentId 课程Id
		ParentId int64 `json:"parentId,string"`
		// Sequence 顺序
		Sequence int8 `json:"sequence" validate:"required"`
	}

	// SwitchSequenceReq 交换讲次请求
	SwitchSequenceReq struct {
		SwitchItems []*SwitchItem `json:"switchItems"`
	}

	// LectureInfo 章节子节点信息
	LectureInfo struct {
		*Lecture

		// Contents 内容
		Contents []*LectureContent `json:"contents"`
	}

	// ChapterInfo 章节信息
	ChapterInfo struct {
		*Lecture

		Lectures []*LectureInfo `json:"lectures"`
	}
)

func (alr *AddLectureReq) Model() *Lecture {
	return &Lecture{
		CourseId: alr.CourseId,
		Type:     alr.Type,
		ParentId: alr.ParentId,
		Sequence: alr.Sequence,
		Name:     alr.Name,
		Profile:  alr.Profile,
	}
}

func (ulr *UpdateLectureReq) Model() *Lecture {
	return &Lecture{
		BaseStruct: gox.BaseStruct{
			Id: ulr.Id,
		},
		CourseId: ulr.CourseId,
		Type:     ulr.Type,
		ParentId: ulr.ParentId,
		Sequence: ulr.Sequence,
		Name:     ulr.Name,
		Profile:  ulr.Profile,
	}
}

func (slr *SwitchSequenceReq) Models() (items []*Lecture) {
	items = make([]*Lecture, 0, len(slr.SwitchItems))
	for _, item := range slr.SwitchItems {
		lecture := &Lecture{
			BaseStruct: gox.BaseStruct{
				Id: item.Id,
			},
			ParentId: item.ParentId,
			Sequence: item.Sequence,
		}
		items = append(items, lecture)
	}

	return
}
