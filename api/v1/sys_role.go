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

func CreateRole(c *gin.Context) {
	var crr request.CreateRoleReq
	err := c.ShouldBindJSON(&crr)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	role := models.SysRole{
		Code:      crr.Code,
		DataScope: crr.DataScope,
		Name:      crr.Name,
		Sort:      crr.Sort,
		Status:    crr.Status,
	}
	err = service.CreateRole(role)
	if err != nil {
		global.OPS_LOGGER.Error("添加失败!", zap.Error(err))
		utils.FailWithMessage("添加失败", c)
		return
	}
	utils.SuccessWithMessage("添加成功", c)
}

func GetRoleOptions(c *gin.Context) {
	roleOptions, err := service.GetRoleOptions()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(roleOptions, "获取成功", c)
}

func GetRolePagination(c *gin.Context) {
	var rpi request.RolePaginationInfo
	err := c.ShouldBindQuery(&rpi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetRolePagination(rpi)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
func getRoleIdFromParam(c *gin.Context) (uint, error) {
	roleId := c.Param("roleId")
	if roleId == "" {
		return 0, errors.New("roleId不能为空")
	}
	id, err := strconv.Atoi(roleId)
	if err != nil {
		return 0, errors.New("转换roleId失败")
	}
	return uint(id), nil
}
func GetRoleForm(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetRoleForm(roleId)
	if err != nil {
		global.OPS_LOGGER.Error("获取角色信息失败:", zap.Error(err))
		utils.FailWithMessage("获取角色信息失败:", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
func UpdateRole(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	var role models.SysRole
	err = c.ShouldBindJSON(&role)
	if err != nil {
		global.OPS_LOGGER.Error("更新角色参数异常", zap.Error(err))
		utils.FailWithMessage("更新角色参数异常", c)
		return
	}
	err = service.UpdateRole(roleId, role)
	if err != nil {
		global.OPS_LOGGER.Error("更新角色信息失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}
func GetRoleMenus(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetRoleMenus(roleId)
	if err != nil {
		global.OPS_LOGGER.Error("获取角色菜单失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func AssignMenuToRole(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	var menuIds []uint
	err = c.ShouldBindJSON(&menuIds)
	if err != nil {
		global.OPS_LOGGER.Error("更新角色参数异常", zap.Error(err))
		utils.FailWithMessage("更新角色参数异常", c)
		return
	}

	err = service.AssignMenuToRole(roleId, menuIds)
	if err != nil {
		global.OPS_LOGGER.Error("分配角色菜单失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}

func GetRoleApis(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetRoleApis(roleId)
	if err != nil {
		global.OPS_LOGGER.Error("获取角色api失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func AssignApiToRole(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	var apiIds []uint
	err = c.ShouldBindJSON(&apiIds)
	if err != nil {
		global.OPS_LOGGER.Error("更新角色参数异常", zap.Error(err))
		utils.FailWithMessage("更新角色参数异常", c)
		return
	}
	err = service.AssignApiToRole(roleId, apiIds)
	if err != nil {
		global.OPS_LOGGER.Error("分配角色api失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}
