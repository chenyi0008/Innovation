package model

import "gorm.io/gorm"

type SmsNum struct {
	gorm.Model
	IsAlarm   bool
	Num       int64
	InspectId uint
}
