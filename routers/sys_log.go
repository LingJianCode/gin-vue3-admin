package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitLogRoutes(private *gin.RouterGroup) {
	logRouter := private.Group("/logs")
	{
		logRouter.GET("/page", v1.LogApiApp.GetApiOptionsListPagination)
	}
}
