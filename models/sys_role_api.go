package models

type SysRoleApi struct {
	SysRoleID uint `gorm:"column:sys_role_id"`
	SysApiID  uint `gorm:"column:sys_api_id"`
}

func (SysRoleApi) TableName() string {
	return "sys_role_api"
}
