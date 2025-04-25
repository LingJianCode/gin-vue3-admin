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

func GetMenusList(c *gin.Context) {
	menus, err := service.GetMenuTree()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.OkWithDetailed(menus, "获取成功", c)
}

func CreateMenu(c *gin.Context) {
	var menuReq request.CreateMenuReq
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
	utils.OkWithMessage("添加成功", c)
}

func GetMenuRoutes(c *gin.Context) {
	menus, err := service.GetMenuRouteTree()
	if err != nil {
		global.OPS_LOGGER.Error("获取失败!", zap.Error(err))
		utils.FailWithMessage("获取失败", c)
		return
	}
	utils.OkWithDetailed(menus, "获取成功", c)
}

func GetMenuOptions(c *gin.Context) {
	onlyParent := c.Query("onlyParent")
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
	utils.OkWithDetailed(menuOptions, "获取成功", c)
}
