package response

import "my-ops-admin/models"

type RoleOption struct {
	// 子选项列表
	Children []*RoleOption `json:"children,omitempty"`
	// 选项的标签
	Label string `json:"label"`
	// 标签类型
	Tag string `json:"-"`
	// 选项的值
	Value uint `json:"value"`
}

type RolePage struct {
	List  []models.SysRole `json:"list,omitempty"`
	Total int64            `json:"total,omitempty"`
}
