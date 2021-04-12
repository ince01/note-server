package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title     string
	Icon      string
	Content   string
	Parent    *uint
	CreatedBy uint
}
