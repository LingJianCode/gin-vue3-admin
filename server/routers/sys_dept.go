package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitDeptRoutes(router *gin.RouterGroup) {
	deptRouter := router.Group("/dept")
	{
		deptRouter.GET("", v1.DeptApiApp.GetDeptTree)
		deptRouter.POST("", v1.DeptApiApp.CreateDept)
		deptRouter.GET("/options", v1.DeptApiApp.GetDeptOptions)
	}
}
