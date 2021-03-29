package models

import (
	"github.com/ince01/note-server/pkg/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Gender    string `gorm:"not null"`
	Phone     string
	AvatarUrl string
	Notes     []Note `gorm:"foreignKey:CreatedBy"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if hashedPassword, err := helpers.HashPassword(user.Password); err == nil {
		tx.Statement.SetColumn("password", hashedPassword)
	}
	return err
}

func (user *User) ComparePassword(rawPassword string) bool {
	isMatched := helpers.CheckPasswordHash(rawPassword, user.Password)

	return isMatched
}
