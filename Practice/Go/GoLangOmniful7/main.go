package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	Print("Shubhra")
}

func Print(str string) {
	myError := errors.New("Bhai Tu kya kar raha hain")

	if str == "Shubhra" {
		fmt.Println("Shubhra")
	} else {
		log.Println(myError)
	}

}
