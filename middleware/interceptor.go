package middleware

import (
	"Innovation/model"
	"Innovation/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证请求
		token := c.GetHeader("Authorization")
		token = strings.Split(token, " ")[1]
		flag, username, id := utils.ParseToken(token)
		println("token:", token)
		ctx := context.Background()
		result, err := utils.GetRedisHelper().Get(ctx, token).Result()
		if err != nil {
			//panic(err)
			c.JSON(403, model.NewResponse(0, "token无效", nil))
			c.Abort()
			return
		}

		user := utils.UnMarshalToUser(result)
		_ = user

		if !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"msg":  "unauthorized",
			})
			return
		}

		println("interceptor:", id)
		c.Set("username", username)
		c.Set("id", id)

		// 继续处理请求
		c.Next()
	}
}
