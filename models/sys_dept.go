package models

import (
	"my-ops-admin/global"
)

type SysDept struct {
	global.OPS_MODEL
	Name     string `gorm:"comment:部门名称"`
	Code     string `gorm:"uniqueIndex;comment:部门编号"`
	ParentID uint   `gorm:"comment:父节点id"`
	Sort     int64  `gorm:"comment:显示顺序"`
	Status   int8   `gorm:"comment:状态(1-正常 0-禁用)"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}
