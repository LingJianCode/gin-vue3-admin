package request

type CreateDict struct {
	// 字典编码
	DictCode string `json:"dictCode,omitempty"`
	// 字典名称
	Name string `json:"name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 字典状态（1-启用，0-禁用）
	Status int64 `json:"status,omitempty"`
}

type DictPaginationInfo struct {
	// 关键字(字典名称)
	Keywords string `form:"keywords"  binding:"omitempty"`
	// 页码
	PageNum int `form:"pageNum" binding:"required,gte=1"`
	// 每页记录数
	PageSize int `form:"pageSize" binding:"required,gte=10"`
}
