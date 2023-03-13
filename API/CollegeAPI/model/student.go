package model

type Student struct {
	Roll      int     `json:"roll"  gorm:"primary_key ; autoIncrement:false"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	Marks     int     `json:"marks"`
	Grade     Grade   `json:"grade" gorm:"embedded"`
	Address   Address `gorm:"embedded"`
}
