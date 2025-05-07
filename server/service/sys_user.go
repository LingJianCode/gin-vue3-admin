package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	mycasbin "my-ops-admin/pkg/my_casbin"
	"my-ops-admin/request"
	"my-ops-admin/response"
	"my-ops-admin/utils"
	"time"

	"gorm.io/gorm"
)

// CreateUser 这里应该使用事务？
//
//	@param models.SysUser
//	@return error
//	@author lingjian
//	@date 2025-04-30 15:05:01
func CreateUser(u models.SysUser) error {
	return global.OPS_DB.Transaction(func(tx *gorm.DB) error {
		var user models.SysUser
		// 这里感觉写的有问题，应该所有错误都返回
		if !errors.Is(tx.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
			return errors.New("用户名已注册")
		}
		u.Password = utils.BcryptHash(u.Password)
		err := tx.Create(&u).Error
		if err != nil {
			return err
		}
		// casbin处理

		var userList []string
		var roleIds []uint
		for _, v := range u.Roles {
			userList = append(userList, u.Username)
			roleIds = append(roleIds, v.ID)
		}
		ok, err := mycasbin.Casbin.AddUserRoles(userList, roleIds)
		if !ok {
			return errors.New("casbin fasle")
		}
		return err
	})
}

func GetCurrentUserInfoByID(id uint) (cu response.CurrentUser, err error) {
	var user models.SysUser
	err = global.OPS_DB.Preload("Roles").First(&user, id).Error
	if err != nil {
		return
	}
	cu = response.CurrentUser{
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
		UserID:   user.ID,
		Username: user.Username,
	}

	var roleIds []uint
	for _, v := range user.Roles {
		cu.Roles = append(cu.Roles, v.Code)
		roleIds = append(roleIds, v.ID)
	}
	// 获取角色的按钮权限
	/* 这里需要修改为使用join
		 SELECT
	            DISTINCT t2.perm
	        FROM
	            sys_role_menu t1
	                INNER JOIN sys_menu t2 ON t2.id = t1.menu_id
	                INNER JOIN sys_role t3 ON t3.id = t1.role_id
	        WHERE
	            t2.type = '${@com.youlai.boot.system.enums.MenuTypeEnum@BUTTON.getValue()}'
	            AND t2.perm IS NOT NULL
	            AND t3.CODE IN
	            <foreach collection="roles" item="role" separator="," open="(" close=")">
	                #{role}
	            </foreach>
	*/
	// var roleList []models.SysRole
	// flag := map[uint]bool{}
	// err = global.OPS_DB.Where("id in ?", roleIds).Preload("Menus").Find(&roleList).Error
	// for _, role := range roleList {
	// 	for _, menu := range role.Menus {
	// 		if menu.Type == models.SYS_MENU_TYPE_BUTTON && !flag[menu.ID] {
	// 			flag[menu.ID] = true
	// 			cu.Perms = append(cu.Perms, menu.Perm)
	// 		}
	// 	}
	// }
	var menuList []*models.SysMenu
	err = global.OPS_DB.
		Distinct(`"sys_menu"."id","sys_menu"."perm","sys_menu"."sort"`).
		Joins(`JOIN "sys_role_menu" ON "sys_role_menu"."sys_menu_id" = "sys_menu"."id" AND "sys_role_menu"."sys_role_id" in ?`, roleIds).
		Where(`"sys_menu"."type" = ?`, models.SYS_MENU_TYPE_BUTTON).
		Order(`"sys_menu"."sort"`).
		Preload("Params").
		Find(&menuList).Error
	cu.Perms = []string{}
	for _, m := range menuList {
		cu.Perms = append(cu.Perms, m.Perm)
	}
	return
}

func GetUserListPagination(upi request.UserPaginationInfo) (ulp response.UserListPagination, err error) {
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

// UpdateUserInfo 更新用户角色时，还需要对casbin进行处理
//
//	@param uint
//	@param request.UserInfo
//	@return error
//	@author lingjian
//	@date 2025-04-30 11:19:27
func UpdateUserInfo(id uint, ui request.UserInfo) error {
	return global.OPS_DB.Transaction(func(tx *gorm.DB) error {
		var user models.SysUser
		err := tx.First(&user, id).Error
		if err != nil {
			return err
		}

		// 删除用户关联角色
		tx.Where("sys_user_id = ?", id).Delete(&models.SysUserRole{})

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
		// 更新
		err = tx.Save(&user).Error
		if err != nil {
			return err
		}
		// casbin处理

		// 删除
		_, err = mycasbin.Casbin.DeleteUserRole(user.Username)
		if err != nil {
			return err
		}
		// 重新添加
		var userList []string
		var roleIds []uint
		for _, v := range user.Roles {
			userList = append(userList, user.Username)
			roleIds = append(roleIds, v.ID)
		}
		_, err = mycasbin.Casbin.AddUserRoles(userList, roleIds)
		return err
	})
}

func DeleteUserById(id uint) error {
	return global.OPS_DB.Transaction(func(tx *gorm.DB) error {
		var u models.SysUser
		err := tx.Delete(&u, id).Error
		if err != nil {
			return err
		}
		err = tx.Where("sys_user_id = ?", id).Delete(&models.SysUserRole{}).Error
		if err != nil {
			return err
		}
		// casbin处理

		// 删除
		ok, err := mycasbin.Casbin.DeleteUserRole(u.Username)
		if !ok {
			return errors.New("casbin fasle")
		}
		return err
	})
}
