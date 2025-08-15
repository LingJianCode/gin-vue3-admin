package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitApiRoutes(private *gin.RouterGroup) {
	apiRouter := private.Group("/apis")
	{
		apiRouter.GET("/options", v1.ApiApiApp.GetApiOptionsList)
	}
}
