package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title     string
	Content   string
	Done      bool
	CreatedBy uint
}
