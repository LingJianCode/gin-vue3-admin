package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/service"
	"my-ops-admin/utils"

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
	utils.OkWithMessage("添加成功", c)
}

func GetRoleOptions(c *gin.Context) {
	roleOptions, err := service.GetRoleOptions()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.OkWithDetailed(roleOptions, "获取成功", c)
}
