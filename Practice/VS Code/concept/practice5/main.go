package main

import "log"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	var persons []Person

	persons = append(persons, Person{"Shubhra", "Ghosh", 26})
	persons = append(persons, Person{"Chirag", "Gupta", 26})
	persons = append(persons, Person{"Soumya", "Mishra", 26})
	persons = append(persons, Person{"Kosa", "Mama", 26})
	persons = append(persons, Person{"Sourik", "Barman", 26})

	for _, i := range persons {
		log.Println("FirstName : ", i.FirstName, "\n", "LastName : ", i.LastName, "\n", "Age : ", i.Age)
	}

}
