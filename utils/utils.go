package utils

import (
	"Innovation/model"
	"github.com/gin-gonic/gin"
	"strconv"
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
	c.AbortWithStatusJSON(500, model.NewResponse(0, "请传参:"+s, nil))
	return 0
}

func QueryAndParseUint(c *gin.Context, s string) uint {
	value := c.Query(s)
	atoi, err := strconv.Atoi(value)
	if err != nil {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "传参不合法:"+s+"请传数字", nil))
		return 0
	}
	return uint(atoi)

}
