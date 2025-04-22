package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/utils"
)

func Login(u *models.SysUser) (*models.SysUser, error) {
	var user models.SysUser
	err := global.OPS_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	user.Password = ""
	return &user, err
}
