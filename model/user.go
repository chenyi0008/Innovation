package model

import (
	"Innovation/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account     string
	Password    string
	InspectList []Inspect `gorm:"foreignKey:UserId"`
}

func UserLogin(account, password string) (*User, bool) {
	db := config.GetDb()
	var user *User
	err := db.Model(&User{}).Where("account = ? and password = ?", account, password).First(&user).Error
	if err != nil {
		return nil, false
	}

	return user, true
}

func UserIfExist(account string) bool {
	db := config.GetDb()
	var num int64
	err2 := db.Model(&User{}).Where("account = ?", account).Count(&num).Error
	if err2 != nil {
		panic(err2.Error())
	}
	return num > 0
}

func UserCreate(user *User) {
	db := config.GetDb()
	err := db.Create(user).Error
	if err != nil {
		println(err)
	}
}

func UserSave(user *User) {
	db := config.GetDb()
	err := db.Create(user).Error
	if err != nil {
		panic(err)
	}
}
