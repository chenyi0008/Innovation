package service

import (
	"Innovation/model"
	"Innovation/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type createRequest struct {
	Status    string
	Name      string
	Location  string
	SerialNum string
}

func InspectCreate(c *gin.Context) {
	var request *createRequest
	c.BindJSON(&request)

	if value, exists := c.Get("id"); exists {

		id, _ := value.(uint)
		println("userId", id)
		inspect := &model.Inspect{
			Status:    request.Status,
			Name:      request.Name,
			Location:  request.Location,
			UserId:    id,
			SerialNum: request.SerialNum,
		}
		fmt.Println("SerialNum:", request.SerialNum)

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

type inspectUpdateRequest struct {
	Status   string
	Name     string
	Location string
	ID       uint
}

func InspectUpdate(c *gin.Context) {
	var request *inspectUpdateRequest
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

// InspectBindAlarm 绑定/**
func InspectBindAlarm(c *gin.Context) {
	userId := utils.GetContextData(c, "id")
	var inspectAlarm *model.InspectAlarm
	err := c.BindJSON(&inspectAlarm)
	if err != nil {
		c.AbortWithStatusJSON(500, model.NewResponse(0, err.Error(), nil))
		return
	}

	inspect := model.InspectGetById(inspectAlarm.InspectId)
	if inspect.UserId != userId {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "你没有权限访问此资源！", nil))
		return
	}
	alarm := model.AlarmGetById(inspectAlarm.AlarmId)
	if alarm.UserId != userId {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "你没有权限访问此资源！", nil))
		return
	}

	save := model.InspectAlarmSave(inspectAlarm)
	if save {
		c.AbortWithStatusJSON(200, model.Succeed("保存成功", nil))
		return
	} else {
		c.AbortWithStatusJSON(500, model.Failed("保存失败", nil))
		return
	}
}

// InspectUnBindAlarm 解绑/*
func InspectUnBindAlarm(c *gin.Context) {
	userId := utils.GetContextData(c, "id")

	inspectIdStr := c.Query("inspectId")
	alarmIdStr := c.Query("alarmId")
	fmt.Println("inspectId:", inspectIdStr)
	fmt.Println("alarmId:", alarmIdStr)

	inspectId, err2 := strconv.Atoi(inspectIdStr)
	alarmId, err3 := strconv.Atoi(alarmIdStr)

	if err2 != nil || err3 != nil {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "参数请传整数", nil))
		return
	}

	inspectIdUint := uint(inspectId)
	alarmIdUint := uint(alarmId)

	inspect := model.InspectGetById(inspectIdUint)
	alarm := model.AlarmGetById(alarmIdUint)

	if inspect.UserId != userId {
		c.AbortWithStatusJSON(403, model.NewResponse(0, "你没有权限访问！", nil))
		return
	}

	remove := model.InspectUnbindAlarm(inspect, alarm)
	if remove {
		c.AbortWithStatusJSON(200, model.Succeed("解绑成功", nil))
		return
	} else {
		c.AbortWithStatusJSON(500, model.Failed("解绑失败", nil))
		return
	}
}

func InspectGetInfoById(c *gin.Context) {
	userId := utils.GetContextData(c, "id")
	inspectIdStr := c.Param("id")
	inspectId, err := strconv.Atoi(inspectIdStr)
	if err != nil {
		c.AbortWithStatusJSON(500, model.Failed("传参出错", nil))
		return
	}
	flag, inspect := model.InspectGetInfoById(uint(inspectId))
	if !flag {
		c.AbortWithStatusJSON(500, model.Failed("查询失败", nil))
		return
	}
	if inspect.UserId != userId {
		c.AbortWithStatusJSON(403, model.Failed("你没有权限访问此资源！", nil))
		return
	}

	c.AbortWithStatusJSON(200, model.Succeed("查询成功", inspect))
	return
}
