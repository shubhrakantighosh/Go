package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Roll      int
	FirstName string
	LastName  string
	Age       int
}

func main() {

	student := Student{11, "Shubhra", "Ghosh", 26}

	student_json, err := json.Marshal(student)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(student_json))

	students := []Student{
		{11, "Shubhra", "Ghosh", 26},
		{12, "Chirag", "Gupta", 26},
		{13, "Soumya", "Mishra", 27},
		{14, "Sourik", "Barman", 25},
		{15, "Yatin", "Kumar", 25},
	}

	students_json, err := json.Marshal(students)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(students_json))

}
