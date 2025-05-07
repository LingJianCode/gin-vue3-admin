package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/service"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetApiOptionsList(c *gin.Context) {
	apiOptions, err := service.GetApiOptions()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(apiOptions, "获取成功", c)
}
