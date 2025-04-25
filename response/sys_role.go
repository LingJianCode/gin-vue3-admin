package response

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
