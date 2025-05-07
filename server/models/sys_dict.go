package models

import "my-ops-admin/global"

type SysDict struct {
	global.OPS_MODEL
	// 字典编码
	DictCode string `gorm:"comment:字典编码" json:"dictCode,omitempty"`
	// 字典名称
	Name string `gorm:"comment:字典名称" json:"name,omitempty"`
	// 备注
	Remark string `gorm:"comment:备注" json:"remark,omitempty"`
	// 字典状态（1-启用，0-禁用）
	Status int64 `gorm:"comment:显示状态(1:显示;0:隐藏)" json:"status,omitempty"`
}
