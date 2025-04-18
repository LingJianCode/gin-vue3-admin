package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("", v1.CreateUser)
	}
}
