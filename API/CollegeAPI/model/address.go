package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	PinCode string `json:"pin_code"`
	City    string `json:"city"`
	State   string `json:"state"`
}
