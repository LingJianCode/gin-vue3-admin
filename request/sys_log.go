package request

type LogPaginationInfo struct {
	// 关键字(字典名称)
	UserId int64 `form:"userId"`
	// 页码
	PageNum int `form:"pageNum" binding:"required,gte=1"`
	// 每页记录数
	PageSize int `form:"pageSize" binding:"required,gte=10"`
}
