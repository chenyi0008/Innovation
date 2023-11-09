package service

import (
	"Innovation/model"
	"Innovation/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AlarmAndUpload(c *gin.Context) {
	id, err := utils.GetID()
	if err != nil {
		panic(err)
	}
	file, err2 := c.FormFile("video")
	if err2 != nil {
		panic(err2)
	}

	inspectId := c.Query("id")
	inspectIdUint, err3 := strconv.Atoi(inspectId)
	if err3 != nil {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "请正确传递检测设备Id", nil))
		return
	}

	inspect := model.InspectGetById(uint(inspectIdUint))
	fileName := strconv.Itoa(int(id)) + "_" + file.Filename
	history := model.History{
		InspectId: inspect.ID,
		Position:  inspect.Location,
		VideoId:   fileName,
		PicId:     "0",
	}

	save := model.HistorySave(&history)
	if !save {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "保存记录失败！", nil))
		return
	}

	bytes, err5 := utils.ConvertFileHeaderToBytes(file)

	utils.QiniuUpload(fileName, bytes)
	videoFile, err4 := utils.ExtractImageFromMP4(bytes)
	utils.QiniuUpload("output.jpg", videoFile)

	if err != nil || err4 != nil || err5 != nil {
		fmt.Println(err4.Error())
		c.AbortWithStatusJSON(500, model.NewResponse(0, "保存视频失败！", nil))
		return
	} else {
		c.AbortWithStatusJSON(200, model.NewResponse(1, "视频保存成功", nil))
		return
	}
}

func HistoryDownload(c *gin.Context) {
	fileId := c.Query("fileId")
	download := utils.QiniuDownload(fileId)
	c.AbortWithStatusJSON(200, model.Succeed("视频查询成功", download))
	return
}

//func HistoryGetAll(c *gin.Context){
//	userId := utils.GetContextData(c, "id")
//
//}
