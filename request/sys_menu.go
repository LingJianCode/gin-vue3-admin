package request

import "my-ops-admin/models"

// MenuForm，菜单表单对象
type CreateMenuReq struct {
	// 【目录】只有一个子路由是否始终显示
	AlwaysShow int64 `json:"alwaysShow,omitempty"`
	// 组件路径(vue页面完整路径，省略.vue后缀)
	Component string `json:"component,omitempty"`
	// 菜单图标
	Icon string `json:"icon,omitempty"`
	// 菜单ID
	ID int64 `json:"id,omitempty"`
	// 【菜单】是否开启页面缓存
	KeepAlive int64 `json:"keepAlive,omitempty"`
	// 菜单名称
	Name string `json:"name,omitempty"`
	// 路由参数
	Params []models.SysMenuParameter `json:"params,omitempty"`
	// 父菜单ID
	ParentID uint `json:"parentId,omitempty"`
	// 权限标识
	Perm string `json:"perm,omitempty"`
	// 跳转路径
	Redirect string `json:"redirect,omitempty"`
	// 路由名称
	RouteName string `json:"routeName,omitempty"`
	// 路由路径
	RoutePath string `json:"routePath,omitempty"`
	// 排序(数字越小排名越靠前)
	Sort int64 `json:"sort,omitempty"`
	// 菜单类型（1-菜单 2-目录 3-外链 4-按钮）
	Type int64 `json:"type,omitempty"`
	// 显示状态(1:显示;0:隐藏)
	Visible int64 `json:"visible,omitempty"`
}
