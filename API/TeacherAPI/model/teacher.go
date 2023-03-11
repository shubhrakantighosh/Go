package model

type Teacher struct {
	Id         int `gorm:"primary_key"`
	Name       string
	Age        int
	Department string
	Address    Address `gorm:"embedded"`
}
