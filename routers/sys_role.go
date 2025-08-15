package routers

import (
	v1 "my-ops-admin/api/v1"
	"my-ops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoleRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/roles").Use(middleware.OperationLog())
	roleRouterWithoutLog := router.Group("/roles")
	{
		roleRouter.POST("", v1.RoleApiApp.CreateRole)
		roleRouter.PUT("/:roleId", v1.RoleApiApp.UpdateRole)
		roleRouter.PUT("/:roleId/menus", v1.RoleApiApp.AssignMenuToRole)
		roleRouter.PUT("/:roleId/apiIds", v1.RoleApiApp.AssignApiToRole)
	}
	{
		roleRouterWithoutLog.GET("/options", v1.RoleApiApp.GetRoleOptions)
		roleRouterWithoutLog.GET("/page", v1.RoleApiApp.GetRolePagination)
		roleRouterWithoutLog.GET("/:roleId/form", v1.RoleApiApp.GetRoleForm)
		roleRouterWithoutLog.GET("/:roleId/menuIds", v1.RoleApiApp.GetRoleMenus)
		roleRouterWithoutLog.GET("/:roleId/apiIds", v1.RoleApiApp.GetRoleApis)
	}
}
