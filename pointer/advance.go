package main

import "fmt"

func main() {
	var i *int
	var j int = 8
	i = &j
	IncrementPointer(i)
	fmt.Println("j values is -> ", j)
	fmt.Println("i value is -> ", *i)

}

func IncrementPointer(i *int) {
	*i++
}
