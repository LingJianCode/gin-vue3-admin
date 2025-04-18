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

func CreateDepartment(c *gin.Context) {
	var cd request.CreateDepartment
	err := c.ShouldBindJSON(&cd)
	if err != nil {
		global.OPS_LOGGER.Error("创建部门参数异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	d := &models.SysDepartment{
		Name:     cd.Name,
		Code:     cd.Code,
		ParentId: cd.ParentID,
		Sort:     cd.Sort,
	}
	err = service.CreateDepartment(*d)
	if err != nil {
		global.OPS_LOGGER.Error("创建部门异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.OkWithMessage("创建部门成功", c)
}
