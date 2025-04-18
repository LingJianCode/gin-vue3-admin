package request

type CreateDepartment struct {
	// 部门编号
	Code string `json:"code,omitempty"`
	// 部门ID
	ID uint `json:"id,omitempty"`
	// 部门名称
	Name string `json:"name,omitempty"`
	// 父部门ID
	ParentID int64 `json:"parentId"`
	// 排序(数字越小排名越靠前)
	Sort int64 `json:"sort,omitempty"`
	// 状态(1:启用;0:禁用)
	Status int64 `json:"status,omitempty"`
}
