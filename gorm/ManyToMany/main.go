package main

import (
	"UserAPI2/config"
	"UserAPI2/model"
	"encoding/json"
	"fmt"
)

func main() {

	Db := config.DataBaseConnection()
	Db.AutoMigrate(&model.Student{}, &model.Teacher{})

	//student1 := model.Student{Roll: 11, Name: "Shubhra"}
	//student2 := model.Student{Roll: 12, Name: "Yatin"}

	//Db.Save(&student1)
	//Db.Save(&student2)

	//teacher1 := model.Teacher{Name: "Ratan", Course: "Java"}
	//teacher2 := model.Teacher{Name: "Ankush", Course: "DSA"}

	//Db.Save(&teacher1)
	//Db.Save(&teacher2)

	//student3 := model.Student{}

	//teacher3 := model.Teacher{}
	//teacher4 := model.Teacher{}
	//
	//Db.Find(&student3, "roll", 11)
	//
	//Db.Find(&teacher3, "name", "Ankush")
	//
	//Db.Find(&teacher4, "name", "Ratan")
	//
	//teachers := []model.Teacher{
	//	teacher4, teacher3,
	//}
	//
	//student3.Teacher = teachers

	//Db.Updates(&student3)
	//
	//student4 := model.Student{}
	//Db.Preload("Teacher").Find(&student4, "roll", 11)
	//
	//json_person, err := json.MarshalIndent(student4, " ", "    ")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(string(json_person))

	//student5 := model.Student{}
	//
	//teacher5 := model.Teacher{}
	//
	//Db.Find(&student5, "roll", 12)
	//
	//Db.Find(&teacher5, "name", "Ratan")
	//
	//teacher6 := []model.Teacher{
	//	teacher5,
	//}
	//
	//student5.Teacher = teacher6
	//
	//Db.Updates(&student5)

	teacher7 := model.Teacher{}

	Db.Preload("Student").Find(&teacher7, "name", "Ankush")

	json_person, err := json.MarshalIndent(teacher7, " ", "    ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json_person))

}
