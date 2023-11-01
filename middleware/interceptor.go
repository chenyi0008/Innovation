package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	utils "goodsManagement/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证请求
		token := c.GetHeader("Authorization")
		flag, username, id := utils.ParseToken(token)

		ctx := context.Background()
		get := utils.GetRedisHelper().Get(ctx, token)
		var permission string
		get.Scan(&permission)

		fmt.Println(permission)
		if !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "unauthorized",
			})
			return
		}

		c.Set("username", username)
		c.Set("id", id)
		// 继续处理请求
		c.Next()
	}
}
