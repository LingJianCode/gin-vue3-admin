package routers

import (
	v1 "my-ops-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("", v1.UserApiApp.CreateUser)
		userRouter.GET("/me", v1.UserApiApp.GetCurrentUserInfo)
		userRouter.GET("/page", v1.UserApiApp.GetUserListPagination)
		userRouter.GET("/:userId/form", v1.UserApiApp.GetUserInfoFormById)
		userRouter.PUT("/:userId/password/reset", v1.UserApiApp.ResetUserPassword)
		userRouter.PUT("/:userId", v1.UserApiApp.UpdateUserInfo)
		userRouter.DELETE("/:userId", v1.UserApiApp.DeleteUser)
		userRouter.GET("/profile", v1.UserApiApp.GetUserProfile)
	}
}
