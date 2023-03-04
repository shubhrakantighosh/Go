package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	HairColour string `json:"hair_colour"`
	HasDog     bool   `json:"has_dog"`
}

func main() {

	jsonUsers := `[
       {
        "first_name"  : "Shubhra",
        "last_name"   : "Ghosh",
        "hair_colour" : "black",
         "has_dog"    : false
        },
        {
        "first_name"  : "Chirag",
        "last_name"   : "Gupta",
        "hair_colour" : "red",
         "has_dog"    :   true
        }
       ]`

	var users []Person

	err := json.Unmarshal([]byte(jsonUsers), &users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)

	var persons []Person

	var person1 Person
	person1.FirstName = "Shubhra"
	person1.LastName = "Ghosh"
	person1.HairColour = "Black"
	person1.HasDog = false

	var person2 Person
	person2.FirstName = "Sourik"
	person2.LastName = "Barman"
	person2.HairColour = "Red"
	person2.HasDog = true

	persons = append(persons, person1)
	persons = append(persons, person2)

	json_person, err := json.MarshalIndent(persons, " ", "    ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json_person))

}
