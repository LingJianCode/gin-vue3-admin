package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"

	"gorm.io/gorm"
)

func CreateRole(r models.SysRole) error {
	if !errors.Is(global.OPS_DB.Where("code = ?", r.Code).First(&models.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.OPS_DB.Create(&r).Error
}
