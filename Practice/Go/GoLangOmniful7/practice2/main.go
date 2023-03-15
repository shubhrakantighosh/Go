package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	StringCheck("Omnifu")
	fmt.Println(Divide(0, 1))

}

func StringCheck(str string) {
	myError := errors.New("Galat Hain")
	if str == "Omniful" {
		fmt.Println(str)
	} else {
		log.Println(myError)
	}
}

func Divide(num1 int, num2 int) (int, error) {
	if num1 == 0 || num2 == 0 {
		return 0, Exception{}
	} else {
		return num1 / num2, nil
	}

}

type Exception struct {
	Message string
}

func (e Exception) Error() string {
	return "Galat hain "
}
