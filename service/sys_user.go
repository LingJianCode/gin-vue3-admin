package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/utils"

	"gorm.io/gorm"
)

func CreateUser(u models.SysUser) error {
	var user models.SysUser
	if !errors.Is(global.OPS_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	return global.OPS_DB.Create(&u).Error
}
