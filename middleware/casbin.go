package middleware

import (
	"my-ops-admin/global"
	mycasbin "my-ops-admin/pkg/my_casbin"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleIds, _ := utils.GetUserRoleIds(c)
		//获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		casbin, err := mycasbin.NewCasbinHandler(global.OPS_DB, global.OPS_LOGGER)
		if err != nil {
			utils.FailWithDetailed(gin.H{}, err.Error(), c)
			c.Abort()
			return
		}
		var ok bool = false
		for _, sub := range roleIds {
			ok, _ = casbin.EnforceRole(sub, obj, act)
			if ok {
				break
			}
		}
		if !ok {
			utils.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
