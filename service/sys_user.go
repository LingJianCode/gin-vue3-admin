package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/response"
	"my-ops-admin/utils"
	"time"

	"gorm.io/gorm"
)

func CreateUser(u models.SysUser) error {
	// 这里后面需要改成事务操作
	var user models.SysUser
	// 这里感觉写的有问题，应该所有错误都返回
	if !errors.Is(global.OPS_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	return global.OPS_DB.Create(&u).Error
}

func GetCurrentUserInfoByID(id uint) (cu response.CurrentUser, err error) {
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

func GetUserListPagenation(upi request.UserPagenationInfo) (ulp response.UserListPage, err error) {
	var userList []models.SysUser
	db := global.OPS_DB.Model(&models.SysUser{})
	limit := upi.PageSize
	offset := upi.PageSize * (upi.PageNum - 1)

	if upi.Keywords != "" {
		db = db.Or("mobile LIKE ?", "%"+upi.Keywords+"%")
		db = db.Or("nickname LIKE ?", "%"+upi.Keywords+"%")
		db = db.Or("username LIKE ?", "%"+upi.Keywords+"%")
	}
	if upi.DeptID != "" {
		db = db.Where("dept_id = ?", upi.DeptID)
	}
	if upi.Status != "" {
		db = db.Where("status = ?", upi.Status)
	}
	layout := "2006-01-02"
	if upi.CreateTime0 != "" {
		startTime, err := time.Parse(layout, upi.CreateTime0)
		if err != nil {
			return ulp, err
		}
		db = db.Where("created_at >= ?", startTime)
	}
	if upi.CreateTime1 != "" {
		endTime, err := time.Parse(layout, upi.CreateTime1)
		if err != nil {
			return ulp, err
		}
		// 设置结束时间为当天的 23:59:59
		endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 23, 59, 59, 0, endTime.Location())
		db = db.Where("created_at <= ?", endTime)
	}
	err = db.Count(&ulp.Total).Error
	if err != nil {
		return
	}
	// 外键也需要Preload
	err = db.Limit(limit).Offset(offset).Preload("Dept").Preload("Roles").Find(&userList).Error
	for _, v := range userList {
		user := response.User{
			Avatar:     v.Avatar,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
			DeptName:   v.Dept.Name,
			Email:      v.Email,
			Gender:     v.Gender,
			ID:         v.ID,
			Mobile:     v.Mobile,
			Nickname:   v.Nickname,
			Status:     v.Status,
			Username:   v.Username,
		}
		roleNames := ""
		for _, r := range v.Roles {
			roleNames += r.Name
			roleNames += ","
		}
		user.RoleNames = utils.RemoveLastChar(roleNames)
		ulp.List = append(ulp.List, user)
	}
	return
}

func GetUserInfoFormById(id uint) (uf response.UserForm, err error) {
	var user models.SysUser
	err = global.OPS_DB.Preload("Roles").First(&user, id).Error
	if err != nil {
		return
	}
	uf = response.UserForm{
		Avatar:   user.Avatar,
		DeptID:   user.DeptID,
		Email:    user.Email,
		Gender:   user.Gender,
		ID:       user.ID,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		OpenID:   user.Openid,
		Status:   user.Status,
		Username: user.Username,
	}
	for _, v := range user.Roles {
		uf.RoleIDS = append(uf.RoleIDS, v.ID)
	}
	return
}

func ResetUserPassword(id uint, password string) error {
	var user models.SysUser
	err := global.OPS_DB.First(&user, id).Error
	if err != nil {
		return err
	}
	user.Password = utils.BcryptHash(password)
	return global.OPS_DB.Save(&user).Error
}

func UpdateUserInfo(id uint, ui request.UserInfo) error {
	return global.OPS_DB.Transaction(func(tx *gorm.DB) error {
		var user models.SysUser
		err := tx.First(&user, id).Error
		if err != nil {
			return err
		}

		// 删除用户关联角色
		tx.Where("sys_user_id = ?", id).Delete(&models.SysUserRole{})
		// 更新
		user.Avatar = ui.Avatar
		user.DeptID = ui.DeptID
		user.Email = ui.Email
		user.Gender = ui.Gender
		user.Mobile = ui.Mobile
		user.Nickname = ui.Nickname
		user.Openid = ui.OpenID
		user.Roles = []models.SysRole{}
		// 重新关联
		for _, v := range ui.RoleIDS {
			user.Roles = append(user.Roles, models.SysRole{OPS_MODEL: global.OPS_MODEL{ID: v}})
		}
		return tx.Save(&user).Error
	})
}
