package routers

import (
	v1 "my-ops-admin/api/v1"
	"my-ops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/users").Use(middleware.OperationLog())
	userRouterWithoutLog := router.Group("/users")
	{
		userRouter.POST("", v1.UserApiApp.CreateUser)
		userRouter.PUT("/:userId/password/reset", v1.UserApiApp.ResetUserPassword)
		userRouter.PUT("/:userId", v1.UserApiApp.UpdateUserInfo)
		userRouter.DELETE("/:userId", v1.UserApiApp.DeleteUser)
	}
	{
		userRouterWithoutLog.GET("/me", v1.UserApiApp.GetCurrentUserInfo)
		userRouterWithoutLog.GET("/page", v1.UserApiApp.GetUserListPagination)
		userRouterWithoutLog.GET("/:userId/form", v1.UserApiApp.GetUserInfoFormById)
		userRouterWithoutLog.GET("/profile", v1.UserApiApp.GetUserProfile)
	}
}
