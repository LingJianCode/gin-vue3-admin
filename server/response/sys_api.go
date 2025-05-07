package response

type ApiOption struct {
	// 选项的值
	Value uint `json:"value"`
	// 选项的标签
	Label string `json:"label,omitempty"`
	// 子选项列表
	Children []*ApiOption `json:"children,omitempty"`
}
