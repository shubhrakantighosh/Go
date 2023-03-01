package main

import "fmt"

func main() {

	//Integer
	var integer1 int = 1
	var integer2 int
	integer2 = 8
	integer3 := 9

	fmt.Println(integer1, integer2, integer3)

	//Float

	var float1 float64 = 2.78
	var float2 float64
	float2 = 3.4
	float3 := 5.987

	fmt.Println(float1, float2, float3)

	//String
	var str1 string = "Hi"
	var str2 string
	str2 = "Hi"
	str3 := "Hi"

	fmt.Println(str1, str2, str3)

	//Array
	var arr1 = [2]int{1, 3}
	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr1, arr2)

	//Slice
	var slice1 = []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	fmt.Println(slice1, slice2)

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

	//Boolean
	var bool1 bool = true
	var bool2 bool
	bool2 = false
	bool3 := true

	fmt.Println(bool1, bool2, bool3)

}
