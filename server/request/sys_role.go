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

type RolePaginationInfo struct {
	// 结束日期
	// EndDate string `json:"endDate,omitempty"`
	// 关键字(角色名称/角色编码)
	Keywords string `form:"keywords,omitempty"`
	// 页码
	PageNum int `form:"pageNum"  binding:"required,gte=1"`
	// 每页记录数
	PageSize int `form:"pageSize"  binding:"required,gte=10"`
	// 开始日期
	// StartDate string `json:"startDate,omitempty"`
}
