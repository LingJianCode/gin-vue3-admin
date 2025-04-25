package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRoleRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/roles")
	{
		roleRouter.POST("", v1.CreateRole)
	}
}
