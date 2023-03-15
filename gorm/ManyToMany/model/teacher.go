package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string
	Course  string
	Student []Student `gorm:"many2many:student_teacher;"`
	//If you remove previous line you can't fetch data from teacher side
}
