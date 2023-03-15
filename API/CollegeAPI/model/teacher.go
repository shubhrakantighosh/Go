package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	AddressId int
	Address   Address   `gorm:"foreignKey:AddressId"`
	Student   []Student `gorm:"foreignKey:teacher_id"`
}
