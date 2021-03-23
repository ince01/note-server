package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Phone     string
	AvatarUrl string
	Notes     []Note `gorm:"foreignKey:CreatedBy"`
}
