package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitApiRoutes(private *gin.RouterGroup) {
	privateAuthRouter := private.Group("/apis")
	{
		privateAuthRouter.GET("/options", v1.GetApiOptionsList)
	}
}
