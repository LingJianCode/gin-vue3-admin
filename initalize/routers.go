package initalize

import (
	"my-ops-admin/routers"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	opsRouter := gin.New()
	opsRouter.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		opsRouter.Use(gin.Logger())
	}
	api := opsRouter.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			routers.InitAuthRoutes(v1)
			routers.InitUserRoutes(v1)
			routers.InitDepartmentRoutes(v1)
		}
	}
	return opsRouter
}
