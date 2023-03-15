package main

import (
	"fmt"
	"log"
)

var myName = "Shubhra"

func main() {

	log.Println(myName)

	fmt.Println("Hello, World.")

	var str string
	str = "Hi, shubhra."

	var integer int
	integer = 9
	fmt.Println(str)
	fmt.Println("i value is ", integer)

	fmt.Println(saySomething())

	anothherWayStore1, anothherWayStore2 := saySomething()
	fmt.Println(anothherWayStore1, anothherWayStore2)

	var store1, store2 string = saySomething()
	fmt.Println(store1, store2)

	var str1 string
	str1 = "Red"
	log.Println("value of str1 is ", str1)
	pointer(&str1)
	log.Println("current value", str1)

}

func saySomething() (string, string) {
	return "something", "Chana"
}

func pointer(s *string) {
	fmt.Println("Chnage 1 ", s)
	log.Println("fmt1 ", s)
	newValue := "Green"
	*s = newValue
	fmt.Println("Chnage 2 ", s)
	log.Println("fmt2 ", s)
}
