package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Roll      int
	Name      string
	TeacherId int
}
