package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	//
	arr := []int{5, 3, 9, 0, 1, 2}
	//log.Println(BubbleSort(arr))
	//
	//MyMap()

	//dog := Dog{Name: "Rappy", Sleeping: false}
	//
	//fmt.Println(dog.sayHello(), "  :  ", dog.eat())

	//teacher := Teacher{Roll: 12, Address: "Kolkata"}
	//println(teacher.name(), " : ", teacher.eat())

	println(TypeConversion(arr))

}

func MyMap() {

	myMap := make(map[int]string)

	myMap[0] = "Rahul"
	myMap[1] = "Shubhra"
	myMap[2] = "Saurav"
	myMap[3] = "Chirag"

	for i, j := range myMap {
		log.Println(i, " ", j)
	}

	var names []string

	for _, i := range myMap {
		log.Println(i)
		names = append(names, i)
	}

	log.Println(names)
	for _, i := range names {
		log.Println(i)
	}

	students := make(map[Roll]Student)

	roll1 := Roll{RollNumber: 1}
	roll2 := Roll{RollNumber: 2}
	roll3 := Roll{RollNumber: 3}
	roll4 := Roll{RollNumber: 4}
	roll5 := Roll{RollNumber: 5}

	student1 := Student{FirstName: "Rahul,", LastName: "Banerjee", Age: 24}
	student2 := Student{FirstName: "Chirag,", LastName: "Gupta", Age: 26}
	student3 := Student{FirstName: "Soumya,", LastName: "Mishra", Age: 25}
	student4 := Student{FirstName: "Kosa", LastName: "Mama", Age: 25}
	student5 := Student{FirstName: "Sourik,", LastName: "Barman", Age: 26}

	students[roll1] = student1
	students[roll2] = student2
	students[roll3] = student3
	students[roll4] = student4
	students[roll5] = student5

	var studentNames []string

	for i, j := range students {
		fmt.Println(i, " : ", j)
		studentNames = append(studentNames, j.FirstName)
	}

	fmt.Println(studentNames)
	sort.Strings(studentNames)
	fmt.Println(studentNames)

}

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

type Roll struct {
	RollNumber int
}

type Animal interface {
	sayHello()
	eat()
}

type Dog struct {
	Name     string
	Sleeping bool
}

func (dog Dog) sayHello() string {
	return "Shampoo"
}

func (dog Dog) eat() string {
	return "meat"
}

type Person interface {
	name()
	eat()
}

type Teacher struct {
	Roll    int
	Address string
}

func (teacher Teacher) name() string {
	return "Alex"
}

func (teacher Teacher) eat() string {
	return "Rice"
}

func TypeConversion(numbers []int) float64 {
	var total, count int

	for _, i := range numbers {
		total += i
		count++
	}
	var avg float64
	avg = float64(total / count)

	var double float64 = 3.543
	var integer int = 3

	println(int(double))
	println(float64(integer))

	return avg
}
