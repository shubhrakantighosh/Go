package main

import "fmt"

func main() {
	str := "Hi"
	fmt.Println(&str)
	Pointer(&str)
	fmt.Println(str)
	fmt.Println(&str)
}

func Pointer(str *string) string {
	*str = "Kela"
	
	return *str
}
