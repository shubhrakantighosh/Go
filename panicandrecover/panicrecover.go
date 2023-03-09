package main

import "fmt"

func main() {
	fmt.Println("Started")

	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
			fmt.Println("Recovering from panic")
		}
	}()

	panic("Panic started")

	fmt.Println("Ended")
}
