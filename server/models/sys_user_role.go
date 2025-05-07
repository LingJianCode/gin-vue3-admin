package models

type SysUserRole struct {
	SysUserID uint `gorm:"column:sys_user_id"`
	SysRoleID uint `gorm:"column:sys_role_id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
