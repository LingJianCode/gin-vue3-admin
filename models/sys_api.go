package models

import "my-ops-admin/global"

type SysApi struct {
	global.OPS_MODEL
	Name   string    `gorm:"comment:api名称" json:"name"`
	Uri    string    `gorm:"comment:api路径" json:"uri"`
	Method string    `gorm:"comment:api方法:GET,POST,PUT,DELETE,PATCH" json:"method"`
	Group  string    `gorm:"comment:api分组名称" json:"group"`
	Roles  []SysRole `json:"-" gorm:"many2many:sys_role_api"`
}

func (SysApi) TableName() string {
	return "sys_api"
}
