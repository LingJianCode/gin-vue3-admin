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
