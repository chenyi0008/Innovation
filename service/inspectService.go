package service

import (
	"Innovation/model"
	"Innovation/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type createRequest struct {
	Status   string
	Name     string
	Location string
}

func InspectCreate(c *gin.Context) {
	var request *createRequest
	c.BindJSON(&request)

	if value, exists := c.Get("id"); exists {

		id, _ := value.(uint)
		println("userId", id)
		inspect := &model.Inspect{
			Status:   request.Status,
			Name:     request.Name,
			Location: request.Location,
			UserId:   id,
		}

		model.InspcetSave(inspect)
		c.JSON(200, model.NewResponse(1, "创建成功", nil))
		return
	}
	c.JSON(500, model.NewResponse(0, "服务器出现问题", nil))
	return

}

func InspectGetAll(c *gin.Context) {
	if value, exists := c.Get("id"); exists {
		id := utils.ToUint(value)
		list := model.InspectGetAll(id)
		c.JSON(200, model.NewResponse(1, "查询成功", list))
		return
	}
	c.JSON(500, model.NewResponse(0, "服务器异常", nil))
	return

}

type updateRequest struct {
	Status   string
	Name     string
	Location string
	ID       uint
}

func InspectUpdate(c *gin.Context) {
	var request *updateRequest
	c.BindJSON(&request)

	id := utils.GetContextData(c, "id")
	res := model.InspectGetById(request.ID)

	if res.UserId != id {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "你没有权限访问此资源！", nil))
		return
	}

	println("userId", id)
	inspect := &model.Inspect{
		Status:   request.Status,
		Name:     request.Name,
		Location: request.Location,
		UserId:   id,
		Model: gorm.Model{
			ID: request.ID,
		},
	}
	model.InspcetSave(inspect)
	c.AbortWithStatusJSON(200, model.NewResponse(1, "修改成功", nil))

}

func InspectDeleteById(c *gin.Context) {
	id := c.Param("id")
	inspectIdInt, _ := strconv.Atoi(id)
	inspectId := uint(inspectIdInt)
	res := model.InspectGetById(inspectId)
	userId := utils.GetContextData(c, "id")
	if res.UserId != userId {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "你没有权限访问此资源！", nil))
		return
	}
	model.InspectDeleteById(inspectId)
	c.AbortWithStatusJSON(200, model.NewResponse(1, "删除成功", nil))
}
