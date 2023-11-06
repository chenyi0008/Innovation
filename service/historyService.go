package service

import (
	"Innovation/model"
	"Innovation/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	path := viper.GetString("file.path") + strconv.Itoa(int(id)) + "_" + file.Filename
	fmt.Println("path:", path)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(500, model.NewResponse(0, "保存视频失败！", nil))
		return
	} else {
		c.AbortWithStatusJSON(200, model.NewResponse(1, "视频保存成功", nil))
		return
	}
}
