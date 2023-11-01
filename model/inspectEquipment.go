package model

import "gorm.io/gorm"

type InspectEquipment struct {
	gorm.Model
	Status             string
	Name               string
	Location           string
	UserId             int64
	AlarmEquipmentList []AlarmEquipment `gorm:"many2many:InspectAlarm;"`
	HistoryList        []History        `gorm:"foreignKey:InspectId"`
	SmsNumList         []SmsNum         `gorm:"foreignKey:InspectId"`
	AlarmNumList       []AlarmNum       `gorm:"foreignKey:InspectId"`
}
