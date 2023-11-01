package model

import "gorm.io/gorm"

type AlarmEquipment struct {
	gorm.Model
	Name                 string
	UserId               int64
	InspectEquipmentList []InspectEquipment `gorm:"many2many:InspectAlarm;"`
}
