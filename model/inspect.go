package model

import (
	"Innovation/config"
	"fmt"
	"gorm.io/gorm"
)

type Inspect struct {
	gorm.Model
	Status             string
	Name               string
	Location           string
	UserId             uint
	AlarmEquipmentList []Alarm    `gorm:"many2many:InspectAlarm;"`
	HistoryList        []History  `gorm:"foreignKey:InspectId"`
	SmsNumList         []SmsNum   `gorm:"foreignKey:InspectId"`
	AlarmNumList       []AlarmNum `gorm:"foreignKey:InspectId"`
}

func InspcetSave(inspect *Inspect) {
	db := config.GetDb()
	err := db.Save(inspect).Error
	if err != nil {
		panic(err)
	}
}

func InspectGetAll(userId uint) []Inspect {
	db := config.GetDb()
	var inspectList []Inspect
	err := db.Where("user_id = ?", userId).Find(&inspectList).Error
	if err != nil {
		panic(err)
	}
	return inspectList
}

func InspectGetById(id uint) *Inspect {
	db := config.GetDb()
	var inspect Inspect
	err := db.First(&inspect, id).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return &inspect
}

func InspectDeleteById(id uint) bool {
	db := config.GetDb()

	err := db.Delete(&Inspect{}, id).Error
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}
