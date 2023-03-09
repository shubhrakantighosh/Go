package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Roll    int `gorm:"primary_key;auto_increment;"`
	Name    string
	Address Address `gorm:"embedded"`
}

type Address struct {
	PinCode string
	City    string
	State   string
}

var students = []Student{

	{Name: "Shubhra", Address: Address{PinCode: "700102", City: "Kolkata", State: "WB"}},
	{Name: "Chirag", Address: Address{PinCode: "713241", City: "Delhi", State: "Delhi"}},
	{Name: "Soumya", Address: Address{PinCode: "876545", City: "Bardhaman", State: "WB"}},
	{Name: "Sourik", Address: Address{PinCode: "876545", City: "Bangalore", State: "KA"}},
	{Name: "Chirag", Address: Address{PinCode: "45678", City: "Indore", State: "MP"}},
}

func mySqlConnection() *gorm.DB {

	dsn := "root:shubhrakanti@tcp(localhost:3306)/db2"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected..")
	return db
}

func GetAllStudents(db *gorm.DB) []Student {

	students := []Student{}

	db.Find(&students)

	if students == nil {
		fmt.Println("Not Found")
	} else {

		for _, i := range students {
			fmt.Println(i)
		}

	}
	return students
}

func UpdateName(db *gorm.DB, roll int, name string) {
	var oldStudent Student
	db.Find(&oldStudent, roll)

	if oldStudent.Roll == 0 {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Old ", oldStudent)
		db.Model(&oldStudent).Update("name", name)
		oldStudent.Name = name
		fmt.Println("New ", oldStudent)
	}

}

func FindStudentWithRoll(db *gorm.DB, roll int) Student {

	student := Student{}

	db.Find(&student, roll)

	if student.Roll == 0 {
		fmt.Println("Not Found")
	} else {
		fmt.Println(student)
	}

	return student

}

func DeleteWithRoll(db *gorm.DB, roll int) {

	err := db.Delete(&Student{}, roll)

	if err.RowsAffected == 0 {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Deleted")
	}

}

func FindStudentWithName(db *gorm.DB, name string) {

	flag := true

	students := []Student{}

	db.Find(&students, "name=?", name)

	for _, i := range students {
		fmt.Println(i)
		flag = false
	}

	if flag {
		fmt.Println("Not Found")
	}

}

func main() {

	db := mySqlConnection()

	db.AutoMigrate(&Student{})
	db.Create(&students)
	fmt.Println("Inserted..")

	GetAllStudents(db)
	UpdateName(db, 2, "Phuchu")
	GetAllStudents(db)
	FindStudentWithRoll(db, 3)
	FindStudentWithName(db, "Chirag")
	DeleteWithRoll(db, 1)

}
