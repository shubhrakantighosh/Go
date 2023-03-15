package main

import "fmt"

func PrintStudents(students map[Student]Address) {

	for i, j := range students {
		fmt.Println(i, " : ", j)
	}

}
