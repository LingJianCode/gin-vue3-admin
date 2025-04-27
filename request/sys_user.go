package request

type CreateUser struct {
	// 用户头像
	Avatar string `json:"avatar,omitempty"`
	// 部门ID
	DeptID uint `json:"deptId,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 性别
	Gender int8 `json:"gender,omitempty"`
	// 手机号码
	Mobile string `json:"mobile,omitempty"`
	// 昵称
	Nickname string `json:"nickname"`
	// 微信openId
	OpenID string `json:"openId,omitempty"`
	// 角色ID集合
	RoleIDS []uint `json:"roleIds"`
	// 用户状态(1:正常;0:禁用)
	Status int8 `json:"status,omitempty"`
	// 用户名
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPagenationInfo struct {
	// 验证器文档：https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Baked_In_Validators_and_Tags
	// dateformat是自定义验证器，内置验证器已经有了时间格式验证
	// 创建时间范围
	CreateTime0 string `form:"createTime[0]" binding:"omitempty,dateformat" time_format:"2006-01-02"`
	// CreateTime1 string `form:"createTime[1]" binding:"dateformat" time_format:"2006-01-02"`
	CreateTime1 string `form:"createTime[1]" binding:"omitempty,datetime=2006-01-02"`

	// 部门ID
	DeptID string `form:"deptId"`
	// 排序方式（正序:ASC；反序:DESC）
	// Direction string `json:"direction,omitempty"`
	// 排序的字段
	// Field string `json:"field,omitempty"`
	// 关键字(用户名/昵称/手机号)
	Keywords string `form:"keywords"`
	// 页码
	PageNum int `form:"pageNum" binding:"required,gte=1"`
	// 每页记录数
	PageSize int `form:"pageSize" binding:"required,gte=10"`
	// 角色ID
	// RoleIDS string `json:"roleIds,omitempty"`
	// 用户状态
	Status string `form:"status"`
}
