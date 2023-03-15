package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	student1 := Student{11, "Shubhra", "Ghosh", 26, "Kolkata"}

	student_json, err1 := json.Marshal(student1)

	if err1 != nil {
		log.Println(err1)
	}
	fmt.Println(string(student_json))

	students1 := []Student{
		{11, "Shubhra", "Ghosh", 26, "Kolkata"},
		{12, "Chirag", "Gupta", 26, "Durgapur"},
		{13, "Soumya", "Mishra", 26, "Bardhaman"},
		{14, "Kosa", "Mama", 26, "Bardhaman"},
		{15, "Subhadeep", "Chatterjee", 26, "Durgapur"},
		{16, "Sourik", "Barman", 26, "Bardhaman"},
	}

	students1_json, err2 := json.Marshal(students1)

	if err2 != nil {
		log.Println(err2)
	}
	fmt.Println(string(students1_json))

	var student5 Student

	json_student3 := []byte(`{"Roll":11,"FirstName":"Shubhra","LastName":"Ghosh","Age":26,"Address":"Kolkata"}`)

	student3 := json.Unmarshal(json_student3, &student5)

	println("Haan Ji")

	if student3 != nil {
		log.Println(student3)
	}

	fmt.Println(student5.FirstName, student5.LastName, student5.Roll, student5.Address)

	var students12 []Student

	students3_json := []byte(`
	[{"Roll":11,"FirstName":"Shubhra","LastName":"Ghosh","Age":26,"Address":"Kolkata"},
	{"Roll":12,"FirstName":"Chirag","LastName":"Gupta","Age":26,"Address":"Durgapur"},
	{"Roll":13,"FirstName":"Soumya","LastName":"Mishra","Age":26,"Address":"Bardhaman"},
	{"Roll":14,"FirstName":"Kosa","LastName":"Mama","Age":26,"Address":"Bardhaman"},
	{"Roll":15,"FirstName":"Subhadeep","LastName":"Chatterjee","Age":26,"Address":"Durgapur"},
	{"Roll":16,"FirstName":"Sourik","LastName":"Barman","Age":26,"Address":"Bardhaman"}]`)

	students11 := json.Unmarshal(students3_json, &students12)

	if students11 != nil {
		log.Println(students11)
	}

	for i := range students12 {
		fmt.Println(students12[i])
	}

	fmt.Println(students12)

}

type Student struct {
	Roll      int
	FirstName string
	LastName  string
	Age       int
	Address   string
}
