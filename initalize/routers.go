package initalize

import (
	"my-ops-admin/middleware"
	"my-ops-admin/routers"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	opsRouter := gin.New()
	opsRouter.Use(gin.Recovery())
	// gin默认是Debug模式
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
			routers.InitDeptRoutes(private)
			routers.InitMenuRoutes(private)
			routers.InitRoleRoutes(private)
			routers.InitDictRoutes(private)
			routers.InitApiRoutes(private)
		}
	}
	return opsRouter
}
