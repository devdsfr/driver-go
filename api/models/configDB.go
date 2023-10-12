package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:1254@tcp(127.0.0.1:3306)/driver?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&Vehicle{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&User{}, &Vehicle{}, &Rent{})
	if err != nil {
		return nil
	}

	return db
}
