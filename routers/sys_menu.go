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
	}
}
