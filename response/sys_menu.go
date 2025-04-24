package response

import "my-ops-admin/models"

type MenuTreeRes struct {
	models.SysMenu
	// 子菜单
	Children []*MenuTreeRes `gorm:"-" json:"children,omitempty"`
}

type MenuRouteRes struct {
	// 子路由列表
	Children []*MenuRouteRes `json:"children,omitempty"`
	// 组件路径
	Component string `json:"component,omitempty"`
	// 必须这样嵌入才能在数据库查询中映射，如果使用 Meta Meta 在查询回来是空
	Meta `json:"meta"`
	// 路由名称
	Name string `json:"name,omitempty"`
	// 路由路径
	Path string `json:"path,omitempty"`
	// 跳转链接
	Redirect string `json:"redirect,omitempty"`
}

// Meta，路由属性类型
type Meta struct {
	// 【目录】只有一个子路由是否始终显示
	AlwaysShow bool `json:"alwaysShow,omitempty"`
	// 是否隐藏(true-是 false-否)
	Hidden bool `json:"hidden,omitempty"`
	// ICON
	Icon string `json:"icon,omitempty"`
	// 【菜单】是否开启页面缓存
	KeepAlive bool `json:"keepAlive,omitempty"`
	// 路由参数
	Params []models.SysMenuParameter `json:"params"`
	// 路由title
	Title string `json:"title,omitempty"`
}
