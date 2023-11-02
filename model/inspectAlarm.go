package model

import "gorm.io/gorm"

type InspectAlarm struct {
	gorm.Model
	InspectId    uint
	AlarmId      uint
	IsBuzzing    bool
	IsFlashing   bool
	ShowLocation bool
}
