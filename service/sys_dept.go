package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/response"
	"reflect"

	"gorm.io/gorm"
)

// CreateDept
//
//	@param models.SysDept
//	@return error
//	@author lingjian
//	@date 2025-04-24 17:56:15
func CreateDept(d models.SysDept) error {
	var sd models.SysDept
	if !errors.Is(global.OPS_DB.Where("code = ?", d.Code).First(&sd).Error, gorm.ErrRecordNotFound) {
		return errors.New("部门编码已注册")
	}
	return global.OPS_DB.Create(&d).Error
}

// GetDeptTree
//
//	@return []*response.DeptTreeRes
//	@return error
//	@author lingjian
//	@date 2025-04-24 17:56:10
func GetDeptTree() (deptTree []*response.DeptTreeRes, err error) {
	var deptList []*response.DeptTreeRes
	err = global.OPS_DB.Find(&deptList).Error
	if err != nil {
		return
	}
	deptTree = buildDeptTree(deptList, 0)
	return
}

// buildDeptTree 考虑泛型?
//
//	@param []*response.DeptTreeRes
//	@param uint
//	@return []*response.DeptTreeRes
//	@author lingjian
//	@date 2025-04-24 17:53:12
func buildDeptTree(deptTree []*response.DeptTreeRes, parentId uint) []*response.DeptTreeRes {
	var nodes []*response.DeptTreeRes
	if reflect.ValueOf(deptTree).IsValid() {
		for _, v := range deptTree {
			if v.ParentID == parentId {
				v.Children = append(v.Children, buildDeptTree(deptTree, v.ID)...)
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}

func GetDeptOptionsTree() (deptOptions []*response.DeptOption, err error) {
	var deptList []*models.SysDept
	err = global.OPS_DB.Find(&deptList).Error
	if err != nil {
		return
	}
	deptOptions = buildDeptOptionsTree(deptList, 0)
	return
}

func buildDeptOptionsTree(deptList []*models.SysDept, parentId uint) []*response.DeptOption {
	var nodes []*response.DeptOption
	if reflect.ValueOf(deptList).IsValid() {
		for _, v := range deptList {
			if v.ParentID == parentId {
				deptOption := response.DeptOption{
					Label: v.Name,
					Value: v.ID,
				}
				deptOption.Children = buildDeptOptionsTree(deptList, v.ID)
				nodes = append(nodes, &deptOption)
			}
		}
	}
	return nodes
}
