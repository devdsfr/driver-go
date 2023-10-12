package models

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"unique"`
	Cpf        string    `json:"cpf"`
	Nascimento time.Time `json:"nascimento"`
	CreatedAt  time.Time `json:"-"`
}
