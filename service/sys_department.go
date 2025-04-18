package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"

	"gorm.io/gorm"
)

func CreateDepartment(d models.SysDepartment) error {
	var sd models.SysDepartment
	if !errors.Is(global.OPS_DB.Where("code = ?", d.Code).First(&sd).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册")
	}
	return global.OPS_DB.Create(&d).Error
}
