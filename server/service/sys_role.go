package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	mycasbin "my-ops-admin/pkg/my_casbin"
	"my-ops-admin/request"
	"my-ops-admin/response"

	"gorm.io/gorm"
)

var RoleServiceApp = new(SysRoleService)

type SysRoleService struct{}

func (a *SysRoleService) CreateRole(r models.SysRole) error {
	// 这里感觉写的有问题，应该所有错误都返回
	if !errors.Is(global.OPS_DB.Where("code = ?", r.Code).First(&models.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("角色已存在")
	}
	return global.OPS_DB.Create(&r).Error
}

func (a *SysRoleService) GetRoleOptions() ([]response.RoleOption, error) {
	var roleList []models.SysRole
	err := global.OPS_DB.Order("sort").Find(&roleList).Error
	if err != nil {
		return nil, err
	}
	var roleOptions []response.RoleOption
	for _, v := range roleList {
		roleOptions = append(roleOptions, response.RoleOption{
			Label: v.Name,
			Value: v.ID,
		})
	}
	return roleOptions, nil
}

func (a *SysRoleService) GetRolePagination(rpi request.RolePaginationInfo) (rp response.RolePaginationRes, err error) {
	db := global.OPS_DB.Model(&models.SysRole{})
	limit := rpi.PageSize
	offset := rpi.PageSize * (rpi.PageNum - 1)
	if rpi.Keywords != "" {
		db = db.Or("name LIKE ?", "%"+rpi.Keywords+"%")
		db = db.Or("code LIKE ?", "%"+rpi.Keywords+"%")
	}
	err = db.Count(&rp.Total).Error
	if err != nil {
		return
	}
	err = db.Order("sort").Limit(limit).Offset(offset).Find(&rp.List).Error
	return
}
func (a *SysRoleService) GetRoleForm(id uint) (r models.SysRole, err error) {
	err = global.OPS_DB.First(&r, id).Error
	return
}

func (a *SysRoleService) UpdateRole(id uint, r models.SysRole) error {
	var or models.SysRole
	if errors.Is(global.OPS_DB.First(&or, id).Error, gorm.ErrRecordNotFound) {
		return errors.New("角色不存在")
	}
	or.Code = r.Code
	or.Name = r.Name
	or.DataScope = r.DataScope
	or.Sort = r.Sort
	or.Status = r.Status
	return global.OPS_DB.Save(&or).Error
}

func (a *SysRoleService) GetRoleMenus(roleId uint) ([]uint, error) {
	var role models.SysRole
	if errors.Is(global.OPS_DB.Order("sort").Preload("Menus").First(&role, roleId).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("角色不存在")
	}
	// 为空则返回空数组，而不是nil；使用var manuList []uint时,如果为空则会返回nil。
	menuList := []uint{}
	for _, v := range role.Menus {
		menuList = append(menuList, v.ID)
	}
	return menuList, nil
}

// AssignMenuToRole 这里应该用事务处理
//
//	@param uint
//	@param []uint
//	@return error
//	@author lingjian
//	@date 2025-04-30 11:18:44
func (a *SysRoleService) AssignMenuToRole(roleId uint, menuIds []uint) error {
	return global.OPS_DB.Transaction(func(tx *gorm.DB) error {
		var role models.SysRole
		if errors.Is(tx.Order("sort").Preload("Menus").First(&role, roleId).Error, gorm.ErrRecordNotFound) {
			return errors.New("角色不存在")
		}
		err := tx.Where("sys_role_id = ?", roleId).Delete(&models.SysRoleMenu{}).Error
		if err != nil {
			return err
		}
		var menus []models.SysMenu
		for _, v := range menuIds {
			menus = append(menus, models.SysMenu{
				OPS_MODEL: global.OPS_MODEL{ID: v},
			})
		}
		role.Menus = menus
		return tx.Save(&role).Error
	})
}

func (a *SysRoleService) GetRoleApis(roleId uint) ([]uint, error) {
	var role models.SysRole
	if errors.Is(global.OPS_DB.Order("sort").Preload("Apis").First(&role, roleId).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("角色不存在")
	}
	// 为空则返回空数组，而不是nil；使用var manuList []uint时,如果为空则会返回nil。
	apiList := []uint{}
	for _, v := range role.Apis {
		apiList = append(apiList, v.ID)
	}
	return apiList, nil
}

// AssignApiToRole 这里应该用事务处理，并对casbin权限进行添加
//
//	@param uint
//	@param []uint
//	@return error
//	@author lingjian
//	@date 2025-04-30 11:18:34
func (a *SysRoleService) AssignApiToRole(roleId uint, apiIds []uint) error {
	return global.OPS_DB.Transaction(func(tx *gorm.DB) error {
		var role models.SysRole
		if errors.Is(tx.Order("sort").Preload("Apis").First(&role, roleId).Error, gorm.ErrRecordNotFound) {
			return errors.New("角色不存在")
		}
		err := tx.Where("sys_role_id = ?", roleId).Delete(&models.SysRoleApi{}).Error
		if err != nil {
			return err
		}
		_, err = mycasbin.Casbin.DeleteRolePolicy(role.ID)
		if err != nil {
			return err
		}
		if len(apiIds) == 0 {
			return nil
		}
		var apis []models.SysApi
		for _, v := range apiIds {
			if v != 0 {
				apis = append(apis, models.SysApi{
					OPS_MODEL: global.OPS_MODEL{ID: v},
				})
			}
		}
		role.Apis = apis
		err = tx.Save(&role).Error
		if err != nil {
			return err
		}
		// casbin
		err = tx.Order("sort").Preload("Apis").First(&role, roleId).Error
		if err != nil {
			return err
		}
		var apiUriList, apiMethodList []string
		for _, v := range role.Apis {
			apiUriList = append(apiUriList, v.Uri)
			apiMethodList = append(apiMethodList, v.Method)
		}
		_, err = mycasbin.Casbin.AddRolePolicies(roleId, apiUriList, apiMethodList)
		return err
	})
}
