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

func HistoryGetAll(userId uint) (bool, []History) {
	db := config.GetDb()
	var historyList []History
	err := db.Where("user_id = ?", userId).Find(historyList).Error
	if err != nil {
		return false, nil
	}
	return true, historyList
}
