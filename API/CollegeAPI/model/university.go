package model

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name      string  `json:"name"`
	AddressId int     `json:"address_id"`
	Address   Address `gorm:"foreignKey:address_id"`
}
