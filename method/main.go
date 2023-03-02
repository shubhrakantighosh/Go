package main

import (
	"fmt"
	"shubhrakantighosh.com/golang/method/person"
)

func main() {
	var person1 = person.Person{"Shubhra", "Ghosh", 26, "Kolkata"}
	fmt.Println(person1.FirstName, person1.LastName, person1.Age, person1.Address)
	fmt.Println(person1.GetFirstName(), person1.GetLastName(), person1.GetAge(), person1.GetAddress())

	var person2 person.Person

	person2.SetFirstName("Chirag")
	person2.SetLastName("Gupta")
	person2.SetAge(25)
	person2.SetAddress("Delhi")
	fmt.Println(person2.GetAll())

	var person3 person.Person

	person3.SetAll("Sourik", "Braman", 25, "Bardhaman")
	fmt.Println(person3.GetAll())

}
