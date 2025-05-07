package response

import "my-ops-admin/models"

type DictPaginationRes struct {
	List  []models.SysDict `json:"list,omitempty"`
	Total int64            `json:"total,omitempty"`
}
