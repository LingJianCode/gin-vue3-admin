package models

import "my-ops-admin/global"

type SysRole struct {
	global.OPS_MODEL
	Code      string    `gorm:"uniqueIndex;comment:角色编码" json:"code"`
	DataScope int64     `gorm:"comment:数据权限" json:"dataScope"`
	Name      string    `gorm:"comment:角色名称" json:"name"`
	Sort      int64     `gorm:"comment:排序" json:"sort"`
	Status    int64     `gorm:"comment:角色状态(1-正常；0-停用)" json:"status"`
	Users     []SysUser `json:"-" gorm:"many2many:sys_user_role"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
