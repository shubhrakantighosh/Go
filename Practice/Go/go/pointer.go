package main

import "fmt"

func main() {
	//a := new(int)
	//*a = 1
	//fmt.Println(a)

	b := 10

	c := &b

	fmt.Println(&b, &c)
	fmt.Println(b, *c)

}
