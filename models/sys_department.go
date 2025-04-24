package models

import (
	"my-ops-admin/global"
)

type SysDepartment struct {
	global.OPS_MODEL
	Name     string `gorm:"comment:部门名称"`
	Code     string `gorm:"comment:部门编号"`
	ParentId int64  `gorm:"comment:父节点id"`
	// Parent   *SysDepartment `gorm:"foreignKey:ParentId;references:ID"`
	TreePath string `gorm:"comment:父节点id路径"`
	Sort     int64  `gorm:"comment:显示顺序"`
	Status   int8   `gorm:"comment:状态(1-正常 0-禁用)"`
}

func (SysDepartment) TableName() string {
	return "sys_dept"
}
