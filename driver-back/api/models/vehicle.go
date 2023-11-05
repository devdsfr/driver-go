package models

type Vehicle struct {
	ID           uint   `gorm:"primaryKey"`
	Model        string `json:"model"`
	LicensePlate string `json:"license_plate" gorm:"unique"`
	Status       string `json:"status"`
	UserID       *uint  `json:"user_id"`
}

func (Vehicle) TableName() string {
	return "vehicles"
}
