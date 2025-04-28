package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitMenuRoutes(router *gin.RouterGroup) {
	menuRouter := router.Group("/menus")
	{
		menuRouter.GET("", v1.GetMenusList)
		menuRouter.POST("", v1.CreateMenu)
		menuRouter.PUT("/:menuId", v1.UpdateMenu)
		menuRouter.DELETE("/:menuId", v1.DeleteMenu)
		menuRouter.GET("/:menuId/form", v1.GetMenuForm)
	}
	{
		menuRouter.GET("/options", v1.GetMenuOptions)
		menuRouter.GET("/routes", v1.GetMenuRoutes)
	}
}
