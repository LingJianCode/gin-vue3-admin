package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	publicAuthRouter := public.Group("/auth")
	{
		publicAuthRouter.GET("/captcha", v1.AuthApiApp.Captcha)
		publicAuthRouter.POST("/login", v1.AuthApiApp.Login)
	}
	privateAuthRouter := private.Group("/auth")
	{
		privateAuthRouter.DELETE("/logout", v1.AuthApiApp.Logout)
	}
}
