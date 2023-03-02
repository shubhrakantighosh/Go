package main

import (
	"fmt"
	"reflect"
)

func main() {

	var float1 float64 = 9.99

	var integer1 int = (int(float1))

	fmt.Println(float1, " -> ", reflect.TypeOf(float1), " : ", integer1, " -> ", reflect.TypeOf(integer1))

	integer2 := 8
	float2 := float64(integer2)

	fmt.Println(integer2, " -> ", reflect.TypeOf(integer2), " : ", float2, " -> ", reflect.TypeOf(float2))

}
