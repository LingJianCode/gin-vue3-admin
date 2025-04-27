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
