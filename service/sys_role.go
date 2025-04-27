package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/response"

	"gorm.io/gorm"
)

func CreateRole(r models.SysRole) error {
	if !errors.Is(global.OPS_DB.Where("code = ?", r.Code).First(&models.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.OPS_DB.Create(&r).Error
}

func GetRoleOptions() ([]response.RoleOption, error) {
	var roleList []models.SysRole
	err := global.OPS_DB.Find(&roleList).Error
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
