package response

import "my-ops-admin/models"

type DeptTreeRes struct {
	models.SysDept
	Children []*DeptTreeRes `gorm:"-" json:"children,omitempty"`
}

type DeptOption struct {
	// 子选项列表
	Children []*DeptOption `json:"children,omitempty"`
	// 选项的标签
	Label string `json:"label"`
	// 标签类型
	Tag string `json:"-"`
	// 选项的值
	Value uint `json:"value"`
}
