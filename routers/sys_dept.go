package routers

import (
	v1 "my-ops-admin/api/v1"
	"my-ops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitDeptRoutes(router *gin.RouterGroup) {
	deptRouter := router.Group("/dept").Use(middleware.OperationLog())
	deptRouterWithoutLog := router.Group("/dept")
	{
		deptRouter.POST("", v1.DeptApiApp.CreateDept)

		deptRouterWithoutLog.GET("", v1.DeptApiApp.GetDeptTree)
		deptRouterWithoutLog.GET("/options", v1.DeptApiApp.GetDeptOptions)
	}
}
