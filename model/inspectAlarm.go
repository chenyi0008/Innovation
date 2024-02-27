package model

import (
	"Innovation/config"
	"fmt"
	"gorm.io/gorm"
)

type InspectAlarm struct {
	gorm.Model
	InspectId uint
	AlarmId   uint
}

func InspectAlarmSave(inspectAlarm *InspectAlarm) bool {
	db := config.GetDb()

	var list []InspectAlarm

	err2 := db.Where("inspect_id = ? and alarm_id = ?", inspectAlarm.InspectId, inspectAlarm.AlarmId).
		Or("inspect_id = ? and alarm_id = ?", inspectAlarm.AlarmId, inspectAlarm.InspectId).
		Find(&list).Error

	if err2 != nil {
		fmt.Println(err2)
	}
	if len(list) != 0 {
		return false
	}

	err := db.Save(inspectAlarm).Error

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func InspectAlarmRemove(id uint) bool {
	db := config.GetDb()
	err := db.Delete(&InspectAlarm{}, id).Error
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		return false
	}
	return true
}

func InspectAlarmGetById(id uint) (bool, *InspectAlarm) {
	db := config.GetDb()
	var inspectAlarm InspectAlarm
	err := db.First(&inspectAlarm, id).Error
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}
	return true, &inspectAlarm

}
