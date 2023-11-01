package model

import "gorm.io/gorm"

type AlarmNum struct {
	gorm.Model
	IsAlarm   bool
	Num       int64
	InspectId int64
}
