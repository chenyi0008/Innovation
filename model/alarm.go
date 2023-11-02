package model

import "gorm.io/gorm"

type Alarm struct {
	gorm.Model
	Name                 string
	UserId               uint
	InspectEquipmentList []Inspect `gorm:"many2many:InspectAlarm;"`
}
