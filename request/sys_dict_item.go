package request

type DictItemPagenationInfo struct {
	// 关键字(字典名称)
	Keywords string `form:"keywords"  binding:"omitempty"`
	// 页码
	PageNum int `form:"pageNum" binding:"required,gte=1"`
	// 每页记录数
	PageSize int `form:"pageSize" binding:"required,gte=10"`
}
