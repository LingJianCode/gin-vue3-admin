package initalize

import (
	"my-ops-admin/middleware"
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
		public := api.Group("/v1")
		private := api.Group("/v1")
		private.Use(middleware.JwtAuthMiddleware())
		{
			routers.InitAuthRoutes(public, private)
			routers.InitUserRoutes(private)
			routers.InitDepartmentRoutes(private)
		}
	}
	return opsRouter
}
