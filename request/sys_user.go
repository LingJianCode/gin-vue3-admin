package request

type CreateUser struct {
	// 用户头像
	Avatar string `json:"avatar,omitempty"`
	// 部门ID
	DeptID int `json:"deptId,omitempty"`
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
	RoleIDS []int64 `json:"roleIds"`
	// 用户状态(1:正常;0:禁用)
	Status int8 `json:"status,omitempty"`
	// 用户名
	Username string `json:"username"`
	Password string `json:"password"`
}
