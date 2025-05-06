package middleware

import (
	mycasbin "my-ops-admin/pkg/my_casbin"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleIds, _ := utils.GetUserRoleIds(c)
		//获取请求匹配的路由uri
		obj := c.FullPath()
		// 获取请求方法
		act := c.Request.Method

		var ok bool = false
		for _, sub := range roleIds {
			ok, _ = mycasbin.Casbin.EnforceRole(sub, obj, act)
			if ok {
				break
			}
		}
		if !ok {
			utils.FailWithDetailed(gin.H{}, "API权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
