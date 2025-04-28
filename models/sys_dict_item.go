package models

import "my-ops-admin/global"

type SysDictItem struct {
	global.OPS_MODEL
	// 字典编码
	DictCode string `gorm:"comment:字典编码" json:"dictCode,omitempty"`
	// 字典项标签
	Label string `gorm:"comment:字典项标签" json:"label,omitempty"`
	// 排序
	Sort int64 `gorm:"comment:排序" json:"sort,omitempty"`
	// 状态（0：禁用，1：启用）
	Status int64 `gorm:"comment:状态（0：禁用，1：启用）" json:"status,omitempty"`
	// 字典类型（用于显示样式）
	TagType string `gorm:"comment:字典类型（用于显示样式）" json:"tagType,omitempty"`
	// 字典项值
	Value string `gorm:"comment:字典项值" json:"value,omitempty"`
}
