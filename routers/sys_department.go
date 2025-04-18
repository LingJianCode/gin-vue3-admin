package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitDepartmentRoutes(router *gin.RouterGroup) {
	departmentRouter := router.Group("/dept")
	{
		departmentRouter.POST("", v1.CreateDepartment)
	}
}
