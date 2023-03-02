package main

import (
	"fmt"
	"sort"
)

func main() {

	//Slice
	var slice1 = []int{3, 1, 2}
	slice2 := []int{1, 2, 3}

	fmt.Println(slice1, slice2)

	slice1 = append(slice1, 6)
	fmt.Println(slice1)

	sort.Ints(slice2)
	fmt.Println(slice2)

}
