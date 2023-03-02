package main

import "fmt"

func main() {

	//Map
	var map1 = make(map[int]string)
	map1[1] = "One"
	map1[2] = "Two"

	map2 := make(map[int]string)
	map2[1] = "One"
	map2[2] = "Two"

	map3 := map[int]string{
		1: "One",
		2: "Two",
	}

	fmt.Println(map1, map2, map3)

}
