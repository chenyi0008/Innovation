package model

import "gorm.io/gorm"

type InspectAlarm struct {
	gorm.Model
	InspectId    int64
	AlarmId      int64
	IsBuzzing    bool
	IsFlashing   bool
	ShowLocation bool
}
