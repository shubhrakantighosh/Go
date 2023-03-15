package main

import "log"

type Teacher struct {
	FirstName string
}

func main() {

	myMap := make(map[string]string)

	myMap["dog"] = "shampoo"
	myMap["fish"] = "ninja"

	log.Println(myMap["dog"])
	log.Println(myMap["fish"])
	log.Println(myMap["bird"])

	myMap["dog"] = "hulu"

	log.Println(myMap["dog"])

	myMap1 := make(map[Teacher]string)

	teacher := Teacher{
		FirstName: "Shubhra",
	}

	myMap1[teacher] = "Go Lang"

	log.Println(myMap1[teacher])

}
