package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Roll    int
	Name    string
	Teacher []Teacher `gorm:"many2many:student_teacher;"`
}
