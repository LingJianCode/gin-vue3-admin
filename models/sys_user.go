package models

import "my-ops-admin/global"

type SysUser struct {
	global.OPS_MODEL
	Username string         `gorm:"index;comment:用户登录名"`
	Nickname string         `gorm:"default:系统用户;comment:用户昵称"`
	Gender   int8           `gorm:"comment:性别"`
	Password string         `gorm:"comment:用户登录密码"`
	DeptID   int            `gorm:"comment:用户登录密码"`
	Dept     *SysDepartment `gorm:"foreignKey:DeptID;references:ID;comment:关联部门"`
	Avatar   string         `gorm:"default:header.jpg;comment:用户头像"`
	Mobile   string         `gorm:"comment:用户手机号"`
	Status   int8           `gorm:"default:1;comment:状态(1-正常 0-禁用)"`
	Email    string         `gorm:"comment:用户邮箱"`
	Openid   string         `gorm:"comment:微信openid"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
