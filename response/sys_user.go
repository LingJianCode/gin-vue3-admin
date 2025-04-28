package response

type CurrentUser struct {
	// 头像地址
	Avatar string `json:"avatar"`
	// 用户昵称
	Nickname string `json:"nickname"`
	// 用户权限标识集合
	Perms []string `json:"perms"`
	// 用户角色编码集合
	Roles []string `json:"roles"`
	// 用户ID
	UserID uint `json:"userId"`
	// 用户名
	Username string `json:"username"`
}

type UserListPage struct {
	List  []User `json:"list,omitempty"`
	Total int64  `json:"total,omitempty"`
}

type User struct {
	// 用户头像地址
	Avatar string `json:"avatar,omitempty"`
	// 创建时间
	CreateTime string `json:"createTime,omitempty"`
	// 部门名称
	DeptName string `json:"deptName,omitempty"`
	// 用户邮箱
	Email string `json:"email,omitempty"`
	// 性别
	Gender int8 `json:"gender,omitempty"`
	// 用户ID
	ID uint `json:"id,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 用户昵称
	Nickname string `json:"nickname,omitempty"`
	// 角色名称，多个使用英文逗号(,)分割
	RoleNames string `json:"roleNames,omitempty"`
	// 用户状态(1:启用;0:禁用)
	Status int8 `json:"status,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
}

type UserForm struct {
	// 用户头像
	Avatar string `json:"avatar,omitempty"`
	// 部门ID
	DeptID uint `json:"deptId,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 性别
	Gender int8 `json:"gender,omitempty"`
	// 用户ID
	ID uint `json:"id,omitempty"`
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
}
