package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitDictRoutes(router *gin.RouterGroup) {
	dictRouter := router.Group("/dicts")
	// 字典
	{
		dictRouter.GET("", v1.DictApiApp.GetDicts)
		dictRouter.GET("/page", v1.DictApiApp.GetDictsPagination)
		dictRouter.POST("", v1.DictApiApp.CreateDict)
		dictRouter.GET("/:dictId/form", v1.DictApiApp.GetDictForm)
		dictRouter.PUT("/:dictId", v1.DictApiApp.UpdateDict)
		dictRouter.DELETE("/:dictId", v1.DictApiApp.DeleteDict)
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
