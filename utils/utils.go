package utils

import (
	"Innovation/model"
	"github.com/gin-gonic/gin"
)

func ToUint(value any) uint {
	id, _ := value.(uint)
	return id
}

func GetContextData(c *gin.Context, s string) uint {
	if value, exists := c.Get(s); exists {
		id, _ := value.(uint)
		return id
	}
	c.AbortWithStatusJSON(500, model.NewResponse(0, "服务器出错", nil))
	return 0
}
