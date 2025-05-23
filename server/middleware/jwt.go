package middleware

import (
	"errors"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.ExtractToken(c)
		if token == "" {
			utils.NoAuth("请传入token", c)
			c.Abort()
			return
		}
		claims, err := utils.ParseTokenHs256(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) {
				utils.NoAuth("授权已过期", c)
				c.Abort()
				return
			}
			utils.NoAuth(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
