package response

import "my-ops-admin/models"

type DictPagenationRes struct {
	List  []models.SysDict `json:"list,omitempty"`
	Total int64            `json:"total,omitempty"`
}
