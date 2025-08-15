package models

import "my-ops-admin/global"

// SysUser
//  @author lingjian
//  @date 2025-04-24 17:56:22
type SysUser struct {
	global.BASE_MODEL
	Username string    `gorm:"uniqueIndex;comment:用户登录名"`
	Nickname string    `gorm:"default:系统用户;comment:用户昵称"`
	Gender   int8      `gorm:"comment:性别"`
	Password string    `gorm:"comment:用户登录密码" json:"-" `
	DeptID   uint      `gorm:"comment:用户登录密码"`
	Dept     *SysDept  `gorm:"foreignKey:DeptID;references:ID;comment:关联部门"`
	Avatar   string    `gorm:"default:https://foruda.gitee.com/images/1723603502796844527/03cdca2a_716974.gif;comment:用户头像"`
	Mobile   string    `gorm:"comment:用户手机号"`
	Status   int8      `gorm:"default:1;comment:状态(1-正常 0-禁用)"`
	Email    string    `gorm:"comment:用户邮箱"`
	Openid   string    `gorm:"comment:微信openid"`
	Roles    []SysRole `gorm:"many2many:sys_user_role"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
