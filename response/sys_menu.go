package response

// MenuVO，菜单视图对象
type GetMenuRes struct {
	// 子菜单
	Children []GetMenuRes `json:"children,omitempty"`
	// 组件路径
	Component string `json:"component,omitempty"`
	// ICON
	Icon string `json:"icon,omitempty"`
	// 菜单ID
	ID uint `json:"id,omitempty"`
	// 菜单名称
	Name string `json:"name,omitempty"`
	// 父菜单ID
	ParentID uint `json:"parentId,omitempty"`
	// 按钮权限标识
	Perm string `json:"perm,omitempty"`
	// 跳转路径
	Redirect string `json:"redirect,omitempty"`
	// 路由名称
	RouteName string `json:"routeName,omitempty"`
	// 路由路径
	RoutePath string `json:"routePath,omitempty"`
	// 菜单排序(数字越小排名越靠前)
	Sort int64 `json:"sort,omitempty"`
	// 菜单类型
	Type int64 `json:"type,omitempty"`
	// 菜单是否可见(1:显示;0:隐藏)
	Visible int64 `json:"visible,omitempty"`
}
