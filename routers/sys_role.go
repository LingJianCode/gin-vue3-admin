package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRoleRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/roles")
	{
		roleRouter.POST("", v1.CreateRole)
		roleRouter.GET("/options", v1.GetRoleOptions)
		roleRouter.GET("/page", v1.GetRolePagenation)
		roleRouter.GET("/:roleId/form", v1.GetRoleForm)
		roleRouter.PUT("/:roleId", v1.UpdateRole)
		roleRouter.GET("/:roleId/menuIds", v1.GetRoleMenu)
		roleRouter.PUT("/:roleId/menus", v1.AssignMenuToRole)
	}
}
