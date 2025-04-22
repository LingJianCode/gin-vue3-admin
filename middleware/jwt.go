package middleware

import (
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.ExtractToken(c)
		if token == "" {
			utils.NoAuth("请传入token", c)
			c.Abort()
			return
		}
		claims, err := utils.ParseTokenHs256(token)
		if err != nil {
			utils.NoAuth(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
