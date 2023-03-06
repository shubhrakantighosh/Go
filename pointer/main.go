package main

import "fmt"

func main() {

	// '&' -> Reference to that pointer // Also use to check memory address
	// '*' (asterisk) To get the value from reference that pointer we have to use

	num := 9

	var pointer = &num

	fmt.Println(pointer, &num)
	fmt.Println(*pointer, num)

	*pointer = *pointer + 18

	fmt.Println(pointer, &num)
	fmt.Println(*pointer, num)

}
