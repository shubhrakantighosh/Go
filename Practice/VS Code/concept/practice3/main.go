package main

import (
	"log"
	"sort"
)

func main() {

	var str []string

	str = append(str, "Shyam")
	str = append(str, "Arjun")
	log.Println(str)
	sort.Strings(str)
	log.Println(str)

	numbers := [5]int{4, 7, 2, 1}

	log.Println(numbers)
	log.Println(numbers[2:4])

}
