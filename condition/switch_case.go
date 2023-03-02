package main

import "fmt"

func main() {

	example1 := 1

	switch example1 {

	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	default:
		fmt.Println("Wrong Input")

	}

	example2 := "One"

	switch example2 {

	case "One":
		fmt.Println(1)

	case "Two":
		fmt.Println(2)

	default:
		fmt.Println("Wrong In put")

	}

}
