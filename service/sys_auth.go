package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/utils"
)

var AuthServiceApp = new(SysAuthService)

type SysAuthService struct{}

func (a *SysAuthService) Login(u *models.SysUser) (*models.SysUser, error) {
	var user models.SysUser
	err := global.OPS_DB.Preload("Roles").Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}
