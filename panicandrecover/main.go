package main

import "fmt"

func main() {
	fmt.Println("Started")
	defer Recover()
	ArrayElements(7)
	fmt.Println("Ended")
}

func ArrayElements(index int) {
	arr := []int{1, 7, 5, 3}
	if index > len(arr) {
		panic("Of out bound access")
	}
	fmt.Println(arr[index])
}

func Recover() {
	r := recover()

	if r != nil {
		fmt.Println("Recovering. ", r)
	}

}
