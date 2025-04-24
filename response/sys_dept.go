package response

import "my-ops-admin/models"

type DeptTreeRes struct {
	models.SysDept
	Children []*DeptTreeRes `gorm:"-"`
}
