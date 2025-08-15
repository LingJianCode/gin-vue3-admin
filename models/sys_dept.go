package models

import (
	"my-ops-admin/global"
)

type SysDept struct {
	global.BASE_MODEL
	Name     string `gorm:"comment:部门名称" json:"name"`
	Code     string `gorm:"uniqueIndex;comment:部门编号" json:"code"`
	ParentID uint   `gorm:"comment:父节点id"   json:"parentId"`
	Sort     int64  `gorm:"comment:显示顺序"   json:"sort"`
	Status   int8   `gorm:"comment:状态(1-正常 0-禁用)"  json:"status"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}
