package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db  *gorm.DB
	err error
)

func InitMysql() {

	dsn := "entrepreneurship:qweewqQWEEWQ@tcp(106.52.223.188:3306)/entrepreneurship?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为 Info
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动创建表结构

}

func GetDb() *gorm.DB {
	return Db
}
