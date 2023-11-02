package main

import (
	"Innovation/config"
	"Innovation/router"
	"Innovation/utils"
)

func main() {
	config.InitMysql()
	utils.InitRedis()
	r := router.NewRouter()

	//db := config.GetDb()
	//db.AutoMigrate(&model.User{},
	//	&model.Alarm{},
	//	&model.Inspect{},
	//	&model.InspectAlarm{},
	//	&model.History{},
	//	&model.SmsNum{},
	//	&model.AlarmNum{})

	r.Run("0.0.0.0:8088") // 监听并在 0.0.0.0:8080 上启动服务
}
