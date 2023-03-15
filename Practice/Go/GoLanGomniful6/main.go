package main

import (
	"fmt"
	"time"
)

func main() {

	//defer panic("Hiii")
	//fmt.Println("Haan Ji")
	arr := []int{1, 2, 3}

	fmt.Println(Array(arr, 9))
	time.Sleep(800000000)
	fmt.Println("Haan Ji")
	fmt.Println(ulalal)
	fmt.Println(gingalala)

}

func Array(arr []int, index int) int {
	defer Check()
	if index > (len(arr) - 1) {
		panic("Array Out Bound")
	} else {
		return arr[index]
	}

}

func Check() {
	r := recover()
	if r != nil {
		fmt.Println("Recover Panic ", r)
	}
}

const (
	ulalal    string = "haannji"
	gingalala int    = 7
)
