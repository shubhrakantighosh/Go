package main

import "fmt"

func main() {

	//a :=0

	//a := new(int)
	//*a = 10
	//
	//a := 2
	//b := &a
	//
	//fmt.Println(a)
	//fmt.Println(*b)

	integer := 8
	fmt.Println(&integer)
	Pointer(&integer)
}

func Pointer(integer *int) {
	*integer = *integer * *integer
	fmt.Println(integer, &integer)
}
