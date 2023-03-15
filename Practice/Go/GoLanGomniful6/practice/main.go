package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{1, 23, 3}

	defer fmt.Println("Hi...1")
	fmt.Println("Hi...2")

	fmt.Println(ArrayNumber(arr, 6))

}

func ArrayNumber(arr []int, index int) int {
	defer Check()
	if index > len(arr) {
		panic("array Of bound ")
	}
	return arr[index]

}

func Check() {
	r := recover()
	if r != nil {
		log.Println("Recovering", r)
	}
}
