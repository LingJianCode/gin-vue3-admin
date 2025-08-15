package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/request"
	"my-ops-admin/service"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var LogApiApp = new(SysLogApi)

type SysLogApi struct{}

func (a *SysLogApi) GetApiOptionsListPagination(c *gin.Context) {
	var lpi request.LogPaginationInfo
	err := c.ShouldBindQuery(&lpi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.SysLogServiceApp.GetDictPagination(lpi)
	if err != nil {
		global.OPS_LOGGER.Error("获取字典分页列表出错:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
