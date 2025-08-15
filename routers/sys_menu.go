package routers

import (
	v1 "my-ops-admin/api/v1"
	"my-ops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitMenuRoutes(router *gin.RouterGroup) {
	menuRouter := router.Group("/menus").Use(middleware.OperationLog())
	menuRouterWithoutLog := router.Group("/menus")
	{
		menuRouter.POST("", v1.MenuApiApp.CreateMenu)
		menuRouter.PUT("/:menuId", v1.MenuApiApp.UpdateMenu)
		menuRouter.DELETE("/:menuId", v1.MenuApiApp.DeleteMenu)
	}
	{
		menuRouterWithoutLog.GET("", v1.MenuApiApp.GetMenusList)
		menuRouterWithoutLog.GET("/:menuId/form", v1.MenuApiApp.GetMenuForm)
		menuRouterWithoutLog.GET("/options", v1.MenuApiApp.GetMenuOptions)
		menuRouterWithoutLog.GET("/routes", v1.MenuApiApp.GetMenuRoutes)
	}
}
