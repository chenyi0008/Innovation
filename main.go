package main

import (
	"Innovation/config"
	"Innovation/model"
	"Innovation/router"
)

func main() {
	config.InitMysql()

	r := router.NewRouter()

	db := config.GetDb()
	db.AutoMigrate(&model.User{},
		&model.AlarmEquipment{},
		&model.InspectEquipment{},
		&model.InspectAlarm{},
		&model.History{},
		&model.SmsNum{},
		&model.AlarmNum{})

	r.Run("0.0.0.0:8088") // 监听并在 0.0.0.0:8080 上启动服务
}
