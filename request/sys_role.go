package request

type CreateRoleReq struct {
	// 角色编码
	Code string `json:"code"`
	// 数据权限
	DataScope int64 `json:"dataScope"`
	// 角色名称
	Name string `json:"name"`
	// 排序
	Sort int64 `json:"sort"`
	// 角色状态(1-正常；0-停用)
	Status int64 `json:"status"`
}
