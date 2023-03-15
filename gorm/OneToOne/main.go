package main

import (
	"OneToOne/config"
	"OneToOne/model"
)

func main() {

	Db := config.DataBaseConnection()
	Db.AutoMigrate(&model.Student{}, &model.Address{})

	//address1 := model.Address{PinCode: "700102", City: "Kolkata", State: "WB"}
	//student1 := model.Student{Roll: 11, Name: "Shubhra", AddressId: 1}

	//Db.Save(&address1)
	//Db.Save(&student1)

	student2 := model.Student{}

	Db.Find(&student2, "roll", 11)

	address2 := model.Address{PinCode: "713130", City: "Katwa", State: "WB"}

	student2.Address = address2

	Db.Updates(&student2)

}
