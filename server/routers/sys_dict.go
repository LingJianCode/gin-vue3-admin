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
		dictRouter.GET("/page", v1.GetDictsPagination)
		dictRouter.POST("", v1.CreateDict)
		dictRouter.GET("/:dictId/form", v1.GetDictForm)
		dictRouter.PUT("/:dictId", v1.UpdateDict)
		dictRouter.DELETE("/:dictId", v1.DeleteDict)
	}
	// 字典项
	dictItemRouter := router.Group("/dict")
	{
		dictItemRouter.GET("/:dictCode/items/page", v1.GetDictItemListPagination)
		dictItemRouter.GET("/:dictCode/items", v1.GetDictItemList)
		dictItemRouter.POST("/:dictCode/items", v1.CreateDictItem)
		dictItemRouter.GET("/:dictCode/items/:itemId/form", v1.GetDictItemForm)
		dictItemRouter.PUT("/:dictCode/items/:itemId", v1.UpdateDictItem)
		dictItemRouter.DELETE("/:dictCode/items/:itemIds", v1.DeleteDictItem)
	}
}
