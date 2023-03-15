package main

import (
	"OneToMany/config"
	"OneToMany/model"
	"fmt"
)

func main() {
	Db := config.DataBaseConnection()
	Db.AutoMigrate(&model.Teacher{}, &model.Student{})

	//teacher1 := model.Teacher{Name: "Ratan", Course: "Java"}
	//
	//Db.Save(&teacher1)

	//student1 := model.Student{Roll: 11, Name: "Shubhra", TeacherId: 1}
	//student2 := model.Student{Roll: 12, Name: "Chirag", TeacherId: 1}
	//
	//Db.Save(&student1)
	//Db.Save(&student2)

	//student3 := model.Student{}
	//
	//Db.Find(&student3, "id", 5)
	//
	//student3.Roll = 16
	//student3.Name = "Yatin"
	//
	//Db.Updates(&student3)

	//student4 := model.Student{}
	//
	//Db.Find(&student4, "roll", 16)
	//
	//Db.Delete(&student4)
	//
	//student5 := model.Student{}
	//
	//Db.Find(&student5, "roll", 16)
	//
	//fmt.Println(student5)

	teacher2 := model.Teacher{}

	//Db.Preload("Student").Find(&teacher2, "id", 1)

	Db.Model(&model.Teacher{}).Preload("Student").Find(&teacher2)
	fmt.Println(teacher2)

}
