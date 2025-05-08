package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRoleRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/roles")
	{
		roleRouter.POST("", v1.RoleApiApp.CreateRole)
		roleRouter.GET("/options", v1.RoleApiApp.GetRoleOptions)
		roleRouter.GET("/page", v1.RoleApiApp.GetRolePagination)
		roleRouter.GET("/:roleId/form", v1.RoleApiApp.GetRoleForm)
		roleRouter.PUT("/:roleId", v1.RoleApiApp.UpdateRole)
		roleRouter.GET("/:roleId/menuIds", v1.RoleApiApp.GetRoleMenus)
		roleRouter.PUT("/:roleId/menus", v1.RoleApiApp.AssignMenuToRole)
		roleRouter.GET("/:roleId/apiIds", v1.RoleApiApp.GetRoleApis)
		roleRouter.PUT("/:roleId/apiIds", v1.RoleApiApp.AssignApiToRole)
	}
}
