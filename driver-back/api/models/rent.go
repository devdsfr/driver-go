package models

import (
	"gorm.io/gorm"
	"time"
)

type Rent struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	UserID     uint           `json:"user_id"`
	VehicleID  uint           `json:"vehicle_id"`
	RentDate   time.Time      `json:"rent_date" gorm:"type:date"`
	ReturnDate *time.Time     `json:"return_date" gorm:"default:NULL"`
	User       User           `json:"user" gorm:"foreignKey:UserID"`
	Vehicle    Vehicle        `json:"vehicle" gorm:"foreignKey:VehicleID"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
