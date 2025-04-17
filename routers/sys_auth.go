package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.GET("/captcha", v1.Captcha)
	}
}
