package v1

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/request"
	"my-ops-admin/service"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDicts(c *gin.Context) {}
func GetDictsPagenation(c *gin.Context) {
	var dpi request.DictPagenationInfo
	err := c.ShouldBindQuery(&dpi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetDictPagenation(dpi)
	if err != nil {
		global.OPS_LOGGER.Error("获取字典分页列表出错:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
func CreateDict(c *gin.Context)  {}
func GetDictForm(c *gin.Context) {}
func UpdateDict(c *gin.Context)  {}
func DeleteDict(c *gin.Context)  {}

func getDictCodeFromParam(c *gin.Context) (string, error) {
	dictCode := c.Param("dictCode")
	if dictCode == "" {
		return dictCode, errors.New("dictCode不能为空")
	}
	return dictCode, nil
}

func GetDictItemPagenation(c *gin.Context) {
	dictCode, err := getDictCodeFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数dictCode获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	var dipi request.DictItemPagenationInfo
	err = c.ShouldBindQuery(&dipi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数DictItemPagenationInfo获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetDictItemPagenation(dictCode, dipi)
	if err != nil {
		global.OPS_LOGGER.Error("获取字典项分页列表出错:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func GetDictItems(c *gin.Context) {
	dictCode := c.Param("dictCode")
	if dictCode == "" {
		utils.FailWithMessage("获取失败，dictCode不能为空。", c)
		return
	}
	res, err := service.GetDictItem(dictCode)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func CreateDictItem(c *gin.Context)  {}
func GetDictItemForm(c *gin.Context) {}
func UpdateDictItem(c *gin.Context)  {}
func DeleteDictItem(c *gin.Context)  {}
