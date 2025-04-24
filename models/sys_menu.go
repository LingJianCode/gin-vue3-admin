package models

import (
	"my-ops-admin/global"
)

type SysMenu struct {
	global.OPS_MODEL
	AlwaysShow int64  `gorm:"comment:【目录】只有一个子路由是否始终显示" json:"alwaysShow,omitempty"`
	Component  string `gorm:"comment:组件路径(vue页面完整路径，省略.vue后缀)" json:"component,omitempty"`
	Icon       string `gorm:"comment:菜单图标" json:"icon,omitempty"`
	KeepAlive  int64  `gorm:"comment:【菜单】是否开启页面缓存" json:"keepAlive,omitempty"`
	Name       string `gorm:"uniqueIndex;comment:菜单名称" json:"name,omitempty"`
	ParentID   uint   `gorm:"comment:父菜单ID" json:"parentId,omitempty"`
	Perm       string `gorm:"comment:权限标识" json:"perm,omitempty"`
	Redirect   string `gorm:"comment:跳转路径" json:"redirect,omitempty"`
	RouteName  string `gorm:"comment:路由名称" json:"routeName,omitempty"`
	RoutePath  string `gorm:"comment:路由路径" json:"routePath,omitempty"`
	Sort       int64  `gorm:"comment:排序(数字越小排名越靠前)" json:"sort,omitempty"`
	Type       int64  `gorm:"comment:菜单类型（1-菜单 2-目录 3-外链 4-按钮）" json:"type,omitempty"`
	Visible    int64  `gorm:"显示状态(1:显示;0:隐藏)" json:"visible,omitempty"`
	// Parent   *SysMenu `gorm:"foreignKey:ParentID;references:ID;comment:父菜单ID"`
	// 路由参数， 这种数组关联`Has Many`,需要在关联model中增加本model的ID，例如SysMenuID
	Params []SysMenuParameter
}

// Param，键值对
type SysMenuParameter struct {
	global.OPS_MODEL
	SysMenuID uint   `gorm:"comment:菜单ID" json:"-"`
	Key       string `gorm:"comment:参数key" json:"key,omitempty"`
	Value     string `gorm:"comment:参数value" json:"value,omitempty"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

func (SysMenuParameter) TableName() string {
	return "sys_menu_param"
}
