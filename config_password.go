package yunke

type (
	// PasswordConfig 角色默认密码配置
	PasswordConfig struct {
		// 密码消耗
		Cost int `json:"cost" validate:"omitempty,min=4,max=31"`
		// 默认密码
		Default DefaultPasswordConfig `json:"default"`
	}

	// DefaultPasswordConfig 默认密码配置
	DefaultPasswordConfig struct {
		// 讲师、助教默认密码
		Lecturer string `json:"lecturer" validate:"required,min=8,max=30"`
		// 学生默认密码
		Student string `json:"student" validate:"required,min=8,max=30"`
		// 后台除了管理员的默认密码
		Backend string `json:"backend" validate:"required,min=8,max=30"`
	}
)
