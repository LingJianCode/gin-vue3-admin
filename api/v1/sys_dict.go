package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/service"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDicts(c *gin.Context)           {}
func GetDictsPagenation(c *gin.Context) {}
func CreateDict(c *gin.Context)         {}
func GetDictForm(c *gin.Context)        {}
func UpdateDict(c *gin.Context)         {}
func DeleteDict(c *gin.Context)         {}

func GetDictItemPagenation(c *gin.Context) {}

func GetDictItems(c *gin.Context) {
	dictCode := c.Param("dictCode")
	res, err := service.GetDictItem(dictCode)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.OkWithDetailed(res, "获取成功", c)
}

func CreateDictItem(c *gin.Context)  {}
func GetDictItemForm(c *gin.Context) {}
func UpdateDictItem(c *gin.Context)  {}
func DeleteDictItem(c *gin.Context)  {}
