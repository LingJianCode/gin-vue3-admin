package response

type CurrentUser struct {
	// 头像地址
	Avatar string `json:"avatar,omitempty"`
	// 用户昵称
	Nickname string `json:"nickname,omitempty"`
	// 用户权限标识集合
	Perms []string `json:"perms,omitempty"`
	// 用户角色编码集合
	Roles []string `json:"roles,omitempty"`
	// 用户ID
	UserID uint `json:"userId,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
}
