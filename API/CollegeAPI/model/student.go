package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Roll      int     `json:"roll" gorm:"unique"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	Marks     int     `json:"marks"`
	Grade     Grade   `json:"grade" validate:"regexp=^[A-B-C-D]*$"`
	AddressId int     `json:"address_id"`
	Address   Address `gorm:"foreignKey:address_id"`
	TeacherId int     `json:"teacher_id"`
}
