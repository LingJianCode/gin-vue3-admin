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

var DeptApiApp = new(SysDeptApi)

type SysDeptApi struct{}

func (a *SysDeptApi) CreateDept(c *gin.Context) {
	var cd request.CreateDept
	err := c.ShouldBindJSON(&cd)
	if err != nil {
		global.OPS_LOGGER.Error("创建部门参数异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	d := &models.SysDept{
		Name:     cd.Name,
		Code:     cd.Code,
		ParentID: cd.ParentID,
		Sort:     cd.Sort,
	}
	err = service.DeptServiceApp.CreateDept(*d)
	if err != nil {
		global.OPS_LOGGER.Error("创建部门异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("创建部门成功", c)
}

func (a *SysDeptApi) GetDeptTree(c *gin.Context) {
	dept, err := service.DeptServiceApp.GetDeptTree()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}

	utils.SuccessWithDetailed(dept, "获取成功", c)
}

func (a *SysDeptApi) GetDeptOptions(c *gin.Context) {
	deptOptions, err := service.DeptServiceApp.GetDeptOptionsTree()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(deptOptions, "获取成功", c)
}
