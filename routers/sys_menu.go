package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitMenuRoutes(router *gin.RouterGroup) {
	menuRouter := router.Group("/menus")
	{
		menuRouter.GET("", v1.MenuApiApp.GetMenusList)
		menuRouter.POST("", v1.MenuApiApp.CreateMenu)
		menuRouter.PUT("/:menuId", v1.MenuApiApp.UpdateMenu)
		menuRouter.DELETE("/:menuId", v1.MenuApiApp.DeleteMenu)
		menuRouter.GET("/:menuId/form", v1.MenuApiApp.GetMenuForm)
	}
	{
		menuRouter.GET("/options", v1.MenuApiApp.GetMenuOptions)
		menuRouter.GET("/routes", v1.MenuApiApp.GetMenuRoutes)
	}
}
