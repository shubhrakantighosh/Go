package main

import (
	"fmt"
)

func main() {

	//Array
	var arr1 = [2]int{1, 3}
	arr2 := [3]int{3, 1, 3}

	var array [3]int

	array[0] = 1
	array[1] = 1
	array[2] = 3

	fmt.Println(arr1, arr2, array)

}
