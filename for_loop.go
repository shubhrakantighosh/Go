package main

import "fmt"

func main() {

	str := "Hi,Shubhra"

	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, j := range arr {
		fmt.Println(j)
	}

	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	mmyMap := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
	}

	for i, j := range mmyMap {
		fmt.Println(i, j)
	}

	for i := 0; i < len(mmyMap); i++ {
		fmt.Println(mmyMap[i])
	}

	booleans := [4]bool{true, false, true, true}

	for i := 0; i < len(booleans); i++ {
		fmt.Println(booleans[i])
	}

}
