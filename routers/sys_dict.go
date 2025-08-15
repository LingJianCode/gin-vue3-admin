package routers

import (
	v1 "my-ops-admin/api/v1"
	"my-ops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitDictRoutes(router *gin.RouterGroup) {
	dictRouter := router.Group("/dicts").Use(middleware.OperationLog())
	dictRouterWithoutLog := router.Group("/dicts")
	// 字典
	{
		dictRouter.POST("", v1.DictApiApp.CreateDict)
		dictRouter.PUT("/:dictId", v1.DictApiApp.UpdateDict)
		dictRouter.DELETE("/:dictId", v1.DictApiApp.DeleteDict)

		dictRouterWithoutLog.GET("", v1.DictApiApp.GetDicts)
		dictRouterWithoutLog.GET("/page", v1.DictApiApp.GetDictsPagination)
		dictRouterWithoutLog.GET("/:dictId/form", v1.DictApiApp.GetDictForm)

	}
	// 字典项
	dictItemRouter := router.Group("/dict").Use(middleware.OperationLog())
	dictItemRouterWithoutLog := router.Group("/dict")
	{

		dictItemRouter.POST("/:dictCode/items", v1.CreateDictItem)
		dictItemRouter.PUT("/:dictCode/items/:itemId", v1.UpdateDictItem)
		dictItemRouter.DELETE("/:dictCode/items/:itemIds", v1.DeleteDictItem)

		dictItemRouterWithoutLog.GET("/:dictCode/items/page", v1.GetDictItemListPagination)
		dictItemRouterWithoutLog.GET("/:dictCode/items", v1.GetDictItemList)
		dictItemRouterWithoutLog.GET("/:dictCode/items/:itemId/form", v1.GetDictItemForm)
	}
}
