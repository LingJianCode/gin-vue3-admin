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

func getMenuIdFromParam(c *gin.Context) (uint, error) {
	menuId := c.Param("menuId")
	if menuId == "" {
		return 0, errors.New("menuId不能为空")
	}
	id, err := strconv.Atoi(menuId)
	if err != nil {
		return 0, errors.New("转换menuId失败")
	}
	return uint(id), nil
}
func GetMenusList(c *gin.Context) {
	menus, err := service.GetMenuTree()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(menus, "获取成功", c)
}

func CreateMenu(c *gin.Context) {
	var menuReq request.MenuInfo
	err := c.ShouldBindJSON(&menuReq)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	menu := models.SysMenu{
		AlwaysShow: menuReq.AlwaysShow,
		Component:  menuReq.Component,
		Icon:       menuReq.Icon,
		KeepAlive:  menuReq.KeepAlive,
		Name:       menuReq.Name,
		Params:     menuReq.Params,
		Perm:       menuReq.Perm,
		Redirect:   menuReq.Redirect,
		RouteName:  menuReq.RouteName,
		RoutePath:  menuReq.RoutePath,
		Sort:       menuReq.Sort,
		Visible:    menuReq.Visible,
		ParentID:   menuReq.ParentID,
	}
	err = service.CreateMenu(menu)
	if err != nil {
		global.OPS_LOGGER.Error("添加失败!", zap.Error(err))
		utils.FailWithMessage("添加失败", c)
		return
	}
	utils.SuccessWithMessage("添加成功", c)
}

func GetMenuRoutes(c *gin.Context) {
	userId, err := utils.GetUserID(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	menus, err := service.GetMenuRouteTree(userId)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(menus, "获取成功", c)
}

func GetMenuOptions(c *gin.Context) {
	onlyParent := c.DefaultQuery("onlyParent", "false")
	global.OPS_LOGGER.Debug(onlyParent, zap.String("type", utils.GetType(onlyParent)))
	var onlyParentBool bool = false
	if onlyParent == "true" {
		onlyParentBool = true
	}
	menuOptions, err := service.GetMenuOptions(onlyParentBool)
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(menuOptions, "获取成功", c)
}

func UpdateMenu(c *gin.Context) {
	menuId, err := getMenuIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取menuId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	var menuReq request.MenuInfo
	err = c.ShouldBindJSON(&menuReq)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	err = service.UpdateMenu(menuId, menuReq)
	if err != nil {
		global.OPS_LOGGER.Error("更新失败!", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}

func DeleteMenu(c *gin.Context) {
	menuId, err := getMenuIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取menuId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	err = service.DeleteMenu(menuId)
	if err != nil {
		global.OPS_LOGGER.Error("删除失败!", zap.Error(err))
		utils.FailWithMessage("失败", c)
		return
	}
	utils.SuccessWithMessage("成功", c)
}

func GetMenuForm(c *gin.Context) {
	menuId, err := getMenuIdFromParam(c)
	if err != nil {
		global.OPS_LOGGER.Error("获取menuId失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	res, err := service.GetMenuForm(menuId)
	if err != nil {
		global.OPS_LOGGER.Error("获取菜单信息失败:", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.SuccessWithDetailed(res, "获取成功", c)
}
