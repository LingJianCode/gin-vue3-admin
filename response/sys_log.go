package response

import "my-ops-admin/models"

type LogPaginationRes struct {
	List  []models.SysLog `json:"list,omitempty"`
	Total int64           `json:"total,omitempty"`
}
