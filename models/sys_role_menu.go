package models

type SysRoleMenu struct {
	SysRoleID uint `gorm:"column:sys_role_id"`
	SysMenuID uint `gorm:"column:sys_menu_id"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
