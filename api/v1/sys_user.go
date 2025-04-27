package v1

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/service"
	"my-ops-admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func getUserIdFromParam(c *gin.Context) (uint, error) {
	userId := c.Param("userId")
	if userId == "" {
		return 0, errors.New("userId不能为空")
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		return 0, errors.New("转换userId失败")
	}
	return uint(id), nil
}

func CreateUser(c *gin.Context) {
	var cu request.UserInfo
	err := c.ShouldBindJSON(&cu)
	if err != nil {
		global.OPS_LOGGER.Error("创建用户参数异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}

	var roles []models.SysRole
	for _, v := range cu.RoleIDS {
		role := models.SysRole{OPS_MODEL: global.OPS_MODEL{ID: v}}
		roles = append(roles, role)
	}
	// user := models.SysUser{}：会在栈上分配内存，当该变量超出作用域时，系统会自动回收内存。
	// user := &models.SysUser{}：会在堆上分配内存，需依靠 Go 的垃圾回收机制来回收内存。
	user := &models.SysUser{
		Username: cu.Username,
		Password: cu.Password,
		Nickname: cu.Nickname,
		Gender:   cu.Gender,
		DeptID:   cu.DeptID,
		Avatar:   cu.Avatar,
		Mobile:   cu.Mobile,
		Status:   cu.Status,
		Email:    cu.Email,
		Openid:   cu.OpenID,
		Roles:    roles,
	}

	err = service.CreateUser(*user)
	if err != nil {
		global.OPS_LOGGER.Error("创建用户异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("注册成功", c)
}

func GetCurrentUserInfo(c *gin.Context) {
	id, err := utils.GetUserID(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	ReqUser, err := service.GetCurrentUserInfoByID(id)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(ReqUser, "获取成功", c)
}

func GetUserListPagenation(c *gin.Context) {
	var upi request.UserPagenationInfo
	err := c.ShouldBindQuery(&upi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	userList, err := service.GetUserListPagenation(upi)
	if err != nil {
		global.OPS_LOGGER.Error("获取用户分页列表出错:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(userList, "获取成功", c)
}

func GetUserInfoFormById(c *gin.Context) {
	userId, err := getUserIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取userId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetUserInfoFormById(userId)
	if err != nil {
		global.OPS_LOGGER.Error("获取用户信息失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func ResetUserPassword(c *gin.Context) {
	userId, err := getUserIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取userId失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	password := c.Query("password")
	if password == "" {
		global.OPS_LOGGER.Error("重置密码参数异常，password不能为空", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	err = service.ResetUserPassword(userId, password)
	if err != nil {
		global.OPS_LOGGER.Error("重置密码异常", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}

func UpdateUserInfo(c *gin.Context) {
	var ui request.UserInfo
	err := c.ShouldBindJSON(&ui)
	if err != nil {
		global.OPS_LOGGER.Error("用户参数异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	userId, err := getUserIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取userId失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	err = service.UpdateUserInfo(userId, ui)
	if err != nil {
		global.OPS_LOGGER.Error("更新用户信息失败", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}
