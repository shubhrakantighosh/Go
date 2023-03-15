package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	students := InputUser()

	stu := make(map[Student]Address)

	for i, j := range students {
		stu[i] = j
	}

	students_json, err := json.Marshal(stu)

	fmt.Println(err, string(students_json))

	//PrintStudents(students)
	//
	//
	//f, err := os.Create("data.txt")
	//
	//if err != nil {
	//	fmt.Println(f)
	//}
	//
	//_, err2 := f.WriteString("hii")
	//if err2 != nil {
	//
	//}

}
