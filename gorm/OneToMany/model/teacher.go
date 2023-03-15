package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string
	Course  string
	Student []Student `gorm:"foreignKey:TeacherId"`
}
