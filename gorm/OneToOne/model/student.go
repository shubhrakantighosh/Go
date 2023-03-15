package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Roll      int
	Name      string
	AddressId int
	Address   Address `gorm:"foreignKey:AddressId"`
}
