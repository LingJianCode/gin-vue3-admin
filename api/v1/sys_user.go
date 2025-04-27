package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/service"
	"my-ops-admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var cu request.CreateUser
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
	utils.OkWithMessage("注册成功", c)
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
	utils.OkWithDetailed(ReqUser, "获取成功", c)
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
	utils.OkWithDetailed(userList, "获取成功", c)
}

func GetUserInfoFormById(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		utils.FailWithMessage("获取失败，userId不能为空。", c)
		return
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		global.OPS_LOGGER.Error("获取用户信息失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetUserInfoFormById(uint(id))
	if err != nil {
		global.OPS_LOGGER.Error("获取用户信息失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.OkWithDetailed(res, "获取成功", c)
}
