package main

import (
	"fmt"
	"reflect"
)

const variable1 = "Hi"
const variable2 string = "Hii"

const (
	C = "Circle"
	S = "Square"
)

func main() {

	const name = 8

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(variable1))
	fmt.Println(C)

	const flag = true

}
