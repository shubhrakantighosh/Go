package main

import (
	"encoding/json"
	"fmt"
)

type Students struct {
	Roll      int
	FirstName string
	LastName  string
	Age       int
}

func main() {
	student_json := []byte(`{"Roll":11,"FirstName":"Shubhra","LastName":"Ghosh","Age":26}`)

	students_json := []byte(`[{"Roll":11,"FirstName":"Shubhra","LastName":"Ghosh","Age":26},
                       {"Roll":12,"FirstName":"Chirag","LastName":"Gupta","Age":26},
                       {"Roll":13,"FirstName":"Soumya","LastName":"Mishra","Age":27},
                       {"Roll":14,"FirstName":"Sourik","LastName":"Barman","Age":25},
                       {"Roll":15,"FirstName":"Yatin","LastName":"Kumar","Age":25}]`)

	var student Students

	student_struct := json.Unmarshal(student_json, &student)

	if student_struct != nil {
		fmt.Println(student_struct.Error())
	}

	fmt.Println(student)

	var students []Students

	students_struct := json.Unmarshal(students_json, &students)

	if students_struct != nil {
		fmt.Println(students_struct.Error())
	}

	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

}
