package response

type DictItem struct {
	// 字典项标签
	Label string `json:"label,omitempty"`
	// 标签类型
	TagType string `json:"tagType,omitempty"`
	// 字典项值
	Value string `json:"value,omitempty"`
}
