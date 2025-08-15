package models

import (
	"my-ops-admin/global"
)

const (
	SYS_MENU_TYPE_NULL    = 0 // 未知
	SYS_MENU_TYPE_MENU    = 1 // 菜单
	SYS_MENU_TYPE_CATALOG = 2 // 目录
	SYS_MENU_TYPE_EXTLINK = 3 // 外链
	SYS_MENU_TYPE_BUTTON  = 4 // 按钮
)

type SysMenu struct {
	global.BASE_MODEL
	AlwaysShow int64  `gorm:"comment:【目录】只有一个子路由是否始终显示" json:"alwaysShow"`
	Component  string `gorm:"comment:组件路径(vue页面完整路径，省略.vue后缀)" json:"component"`
	Icon       string `gorm:"comment:菜单图标" json:"icon"`
	KeepAlive  int64  `gorm:"comment:【菜单】是否开启页面缓存" json:"keepAlive"`
	Name       string `gorm:"uniqueIndex;comment:菜单名称" json:"name"`
	ParentID   uint   `gorm:"comment:父菜单ID" json:"parentId"`
	Perm       string `gorm:"comment:权限标识" json:"perm"`
	Redirect   string `gorm:"comment:跳转路径" json:"redirect"`
	RouteName  string `gorm:"comment:路由名称" json:"routeName"`
	RoutePath  string `gorm:"comment:路由路径" json:"routePath"`
	Sort       int64  `gorm:"comment:排序(数字越小排名越靠前)" json:"sort"`
	Type       int64  `gorm:"comment:菜单类型（1-菜单 2-目录 3-外链 4-按钮）" json:"type"`
	Visible    int64  `gorm:"显示状态(1:显示;0:隐藏)" json:"visible"`
	// Parent   *SysMenu `gorm:"foreignKey:ParentID;references:ID;comment:父菜单ID"`
	// 路由参数， 这种数组关联`Has Many`,需要在关联model中增加本model的ID，例如SysMenuID
	Params []SysMenuParameter `json:"params"`
	Roles  []SysRole          `json:"-" gorm:"many2many:sys_role_menu"`
}

// Param，键值对
type SysMenuParameter struct {
	global.BASE_MODEL
	SysMenuID uint   `gorm:"comment:菜单ID" json:"-"`
	Key       string `gorm:"comment:参数key" json:"key"`
	Value     string `gorm:"comment:参数value" json:"value"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

func (SysMenuParameter) TableName() string {
	return "sys_menu_param"
}
