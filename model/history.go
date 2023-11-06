package model

import (
	"Innovation/config"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	Position  string
	Notes     string
	VideoId   string
	PicId     string
	InspectId uint
}

func HistorySave(history *History) bool {
	db := config.GetDb()
	err := db.Save(history).Error
	if err != nil {
		println(err)
		return false
	}
	return true
}
