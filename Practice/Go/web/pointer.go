package main

import "fmt"

func main() {

	value := 7

	valueAss := &value

	fmt.Println(valueAss, &value)

	//val := ValueChange(valueAss)

}

func ValueChange(a *int) int {
	return *a + 1
}
