package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/response"
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

func GetUserInfoByID(id uint) (cu response.CurrentUser, err error) {
	var user models.SysUser
	err = global.OPS_DB.Preload("Roles").First(&user, id).Error
	cu = response.CurrentUser{
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
		UserID:   user.ID,
		Username: user.Username,
	}
	for _, v := range user.Roles {
		cu.Roles = append(cu.Roles, v.Code)
	}
	return
}
