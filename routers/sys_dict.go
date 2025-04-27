package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitDictRoutes(router *gin.RouterGroup) {
	dictRouter := router.Group("/dicts")
	// 字典
	{
		dictRouter.GET("", v1.GetDicts)
		dictRouter.GET("/page", v1.GetDictsPagenation)
		dictRouter.POST("", v1.CreateDict)
		dictRouter.GET("/:id/form", v1.GetDictForm)
		dictRouter.PUT("/:id", v1.UpdateDict)
		dictRouter.DELETE("/:ids", v1.DeleteDict)
	}
	// 字典项
	dictItemRouter := router.Group("/dict")
	{
		dictItemRouter.GET("/:dictCode/items/page", v1.GetDictItemPagenation)
		dictItemRouter.GET("/:dictCode/items", v1.GetDictItems)
		dictItemRouter.POST("/:dictCode/items", v1.CreateDictItem)
		dictItemRouter.GET("/:dictCode/items/:itemId/form", v1.GetDictItemForm)
		dictItemRouter.PUT("/:dictCode/items/:itemId", v1.UpdateDictItem)
		dictItemRouter.DELETE("/:dictCode/items/:itemIds", v1.DeleteDictItem)
	}
}
