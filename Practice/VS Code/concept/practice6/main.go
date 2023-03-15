package main

import "log"

type Animal interface {
	sayhello() string
	noOfLegs() int
}

type Dog struct {
	Name string
	Eat  string
}

// type Fish struct {
// 	FishType string
// 	Eye      int
// }

func main() {

	dog := Dog{
		Name: "Shampoo",
		Eat:  "Meat",
	}

	// fish := Fish{
	// 	FishType: "Gold Fish",
	// 	Eye:      2,
	// }

	printAnimal(&dog)

}

func printAnimal(animal Animal) {
	log.Println(animal.sayhello(), animal.noOfLegs())
}

func (a *Dog) sayhello() string {
	return "Woof"
}
func (a *Dog) noOfLegs() int {
	return 4
}
