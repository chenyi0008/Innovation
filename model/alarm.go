package model

import (
	"Innovation/config"
	"fmt"
	"gorm.io/gorm"
)

type Alarm struct {
	gorm.Model
	SerialNum            string
	Name                 string
	UserId               uint
	IsBuzzing            bool
	IsFlashing           bool
	ShowLocation         bool
	InspectEquipmentList []Inspect `gorm:"many2many:InspectAlarm;"`
}

func AlarmSave(alarm *Alarm) bool {
	db := config.GetDb()
	err := db.Save(alarm).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func AlarmGetAll(userId uint) (*[]Alarm, bool) {
	db := config.GetDb()
	var alarmList []Alarm
	err := db.Where("user_id = ?", userId).Find(&alarmList).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	return &alarmList, true
}

func AlarmGetById(id uint) *Alarm {
	db := config.GetDb()
	var alarm Alarm
	err := db.First(&alarm, id).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return &alarm
}

func AlarmDeleteById(id uint) bool {
	db := config.GetDb()
	err := db.Delete(&Alarm{}, id).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
