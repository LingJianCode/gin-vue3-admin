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

var RoleApiApp = new(SysRoleApi)

type SysRoleApi struct{}

func (a *SysRoleApi) CreateRole(c *gin.Context) {
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
	err = service.RoleServiceApp.CreateRole(role)
	if err != nil {
		global.OPS_LOGGER.Error("添加失败!", zap.Error(err))
		utils.FailWithMessage("添加失败", c)
		return
	}
	utils.SuccessWithMessage("添加成功", c)
}

func (a *SysRoleApi) GetRoleOptions(c *gin.Context) {
	roleOptions, err := service.RoleServiceApp.GetRoleOptions()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(roleOptions, "获取成功", c)
}

func (a *SysRoleApi) GetRolePagination(c *gin.Context) {
	var rpi request.RolePaginationInfo
	err := c.ShouldBindQuery(&rpi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.RoleServiceApp.GetRolePagination(rpi)
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
func (a *SysRoleApi) GetRoleForm(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.RoleServiceApp.GetRoleForm(roleId)
	if err != nil {
		global.OPS_LOGGER.Error("获取角色信息失败:", zap.Error(err))
		utils.FailWithMessage("获取角色信息失败:", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
func (a *SysRoleApi) UpdateRole(c *gin.Context) {
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
	err = service.RoleServiceApp.UpdateRole(roleId, role)
	if err != nil {
		global.OPS_LOGGER.Error("更新角色信息失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}
func (a *SysRoleApi) GetRoleMenus(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.RoleServiceApp.GetRoleMenus(roleId)
	if err != nil {
		global.OPS_LOGGER.Error("获取角色菜单失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func (a *SysRoleApi) AssignMenuToRole(c *gin.Context) {
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

	err = service.RoleServiceApp.AssignMenuToRole(roleId, menuIds)
	if err != nil {
		global.OPS_LOGGER.Error("分配角色菜单失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}

func (a *SysRoleApi) GetRoleApis(c *gin.Context) {
	roleId, err := getRoleIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取roleId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.RoleServiceApp.GetRoleApis(roleId)
	if err != nil {
		global.OPS_LOGGER.Error("获取角色api失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func (a *SysRoleApi) AssignApiToRole(c *gin.Context) {
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
	err = service.RoleServiceApp.AssignApiToRole(roleId, apiIds)
	if err != nil {
		global.OPS_LOGGER.Error("分配角色api失败:", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}
