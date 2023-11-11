package main

import (
	"Innovation/config"
	"Innovation/mqtt"
	"Innovation/router"
	"Innovation/utils"
)

func main() {
	config.InitMysql()
	utils.InitRedis()
	utils.Init(1)
	utils.InitConfig()
	utils.QiniuConfigInit()
	r := router.NewRouter()
	mqtt.MqttMain()
	//db := config.GetDb()
	//db.AutoMigrate(&model.User{},
	//	&model.Alarm{},
	//	&model.Inspect{},
	//	&model.InspectAlarm{},
	//	&model.History{},
	//	&model.SmsNum{},
	//	&model.AlarmNum{})

	r.Run("0.0.0.0:9000") // 监听并在 0.0.0.0:8080 上启动服务
}
