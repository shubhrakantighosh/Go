package main

import (
	"log"
	"time"
)

type Person struct {
	FirstName   string
	LastName    string
	Age         int
	DateOfBirth time.Time
}

func (name *Person) printName() string {
	return name.FirstName
}

func main() {

	person := Person{
		FirstName: "Arjun",
	}

	var person1 Person
	person1.FirstName = "Ram"

	// log.Println(person.FirstName)
	// log.Println(person1.FirstName)

	log.Println(person.printName())
	log.Println(person1.printName())

}
