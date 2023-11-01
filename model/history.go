package model

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Position  string
	Notes     string
	VideoId   string
	PicId     string
	InspectId int64
}
