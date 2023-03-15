package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	PinCode string
	City    string
	State   string
}
