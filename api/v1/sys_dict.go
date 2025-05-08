package v1

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/request"
	"my-ops-admin/service"
	"my-ops-admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var DictApiApp = new(SysDictApi)

type SysDictApi struct{}

func getDictIdFromParam(c *gin.Context) (uint, error) {
	dictId := c.Param("dictId")
	if dictId == "" {
		return 0, errors.New("dictId不能为空")
	}
	id, err := strconv.Atoi(dictId)
	if err != nil {
		return 0, errors.New("转换dictId失败")
	}
	return uint(id), nil
}

func (a *SysDictApi) GetDicts(c *gin.Context) {}
func (a *SysDictApi) GetDictsPagination(c *gin.Context) {
	var dpi request.DictPaginationInfo
	err := c.ShouldBindQuery(&dpi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.DictServiceApp.GetDictPagination(dpi)
	if err != nil {
		global.OPS_LOGGER.Error("获取字典分页列表出错:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
func (a *SysDictApi) CreateDict(c *gin.Context) {}
func (a *SysDictApi) GetDictForm(c *gin.Context) {
	dictId, err := getDictIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取dictId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.DictServiceApp.GetDictForm(dictId)
	if err != nil {
		global.OPS_LOGGER.Error("获取dict信息失败:", zap.Error(err))
		utils.FailWithMessage("失败:", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
func (a *SysDictApi) UpdateDict(c *gin.Context) {}
func (a *SysDictApi) DeleteDict(c *gin.Context) {}

func getDictCodeFromParam(c *gin.Context) (string, error) {
	dictCode := c.Param("dictCode")
	if dictCode == "" {
		return dictCode, errors.New("dictCode不能为空")
	}
	return dictCode, nil
}

func GetDictItemListPagination(c *gin.Context) {
	dictCode, err := getDictCodeFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数dictCode获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	var dipi request.DictItemPaginationInfo
	err = c.ShouldBindQuery(&dipi)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数DictItemPagenitionInfo获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.DictItemServiceApp.GetDictItemListPagination(dictCode, dipi)
	if err != nil {
		global.OPS_LOGGER.Error("获取字典项分页列表出错:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func GetDictItemList(c *gin.Context) {
	dictCode, err := getDictCodeFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数dictCode获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.DictItemServiceApp.GetDictItemList(dictCode)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}

func CreateDictItem(c *gin.Context) {}
func GetDictItemForm(c *gin.Context) {
	dictCode, err := getDictCodeFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数dictCode获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}

	itemId, err := getDictItemIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("请求参数itemId获取失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.DictItemServiceApp.GetDictItemForm(dictCode, itemId)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
func UpdateDictItem(c *gin.Context) {}
func DeleteDictItem(c *gin.Context) {}

func getDictItemIdFromParam(c *gin.Context) (uint, error) {
	itemId := c.Param("itemId")
	if itemId == "" {
		return 0, errors.New("itemId不能为空")
	}
	id, err := strconv.Atoi(itemId)
	if err != nil {
		return 0, errors.New("转换itemId失败")
	}
	return uint(id), nil
}
