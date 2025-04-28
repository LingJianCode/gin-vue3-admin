package response

import "my-ops-admin/models"

type DictItem struct {
	// 字典项标签
	Label string `json:"label,omitempty"`
	// 标签类型
	TagType string `json:"tagType,omitempty"`
	// 字典项值
	Value string `json:"value,omitempty"`
}

type DictItemPaginationRes struct {
	List  []models.SysDictItem `json:"list,omitempty"`
	Total int64                `json:"total,omitempty"`
}
