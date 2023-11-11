package service

import (
	"Innovation/model"
	"Innovation/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type alarmCreateRequest struct {
	Name      string
	SerialNum string
}

func AlarmCreate(c *gin.Context) {
	var request alarmCreateRequest
	c.BindJSON(&request)
	if len(request.SerialNum) == 0 || len(request.Name) == 0 {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "设备命名和序列号不能为空", nil))
		return
	}

	value, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "服务器出现问题", nil))
		return
	}
	userId := value.(uint)

	alarm := model.Alarm{
		Name:      request.Name,
		UserId:    userId,
		SerialNum: request.SerialNum,
	}
	save := model.AlarmSave(&alarm)
	if save {
		c.AbortWithStatusJSON(200, model.NewResponse(1, "保存成功", nil))
	} else {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "保存失败", nil))
	}

}

func AlarmGetAll(c *gin.Context) {
	value, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "token无效", nil))
		return
	}
	userId := value.(uint)
	list, flag := model.AlarmGetAll(userId)
	if !flag {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "服务器出错", nil))
		return
	}
	c.AbortWithStatusJSON(200, model.NewResponse(1, "", list))

}

type alarmUpdateRequest struct {
	Name string
	Id   uint
}

func AlarmUpdate(c *gin.Context) {
	var request *alarmUpdateRequest
	c.BindJSON(&request)

	userId := utils.GetContextData(c, "id")
	res := model.AlarmGetById(request.Id)
	if res.UserId != userId {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "你没有权限访问此资源！", nil))
		return
	}

	alarm := &model.Alarm{
		Name:   request.Name,
		UserId: res.UserId,
		Model: gorm.Model{
			ID: res.ID,
		},
	}
	model.AlarmSave(alarm)
	c.AbortWithStatusJSON(200, model.NewResponse(1, "修改成功", nil))

}

func AlarmDeleteById(c *gin.Context) {
	id := c.Param("id")
	alarmIdInt, _ := strconv.Atoi(id)
	alarmId := uint(alarmIdInt)
	res := model.AlarmGetById(alarmId)
	userId := utils.GetContextData(c, "id")
	if res.UserId != userId {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "你没有权限访问此资源！", nil))
		return
	}
	model.AlarmDeleteById(alarmId)
	c.AbortWithStatusJSON(200, model.NewResponse(1, "删除成功", nil))
}
