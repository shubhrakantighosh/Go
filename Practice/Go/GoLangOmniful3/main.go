package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hii")
	//defer println("Kaise Ji")
	//defer fmt.Println("Hello")
	//defer log.Println("Haan Ji")
	//log.Println("Lo Ji")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			defer Kaise_Ji(i)
		} else {
			NaaJi()
		}
	}
}

func Kaise_Ji(i int) {
	fmt.Println("Kaise_Ji ", i)
}

func NaaJi() {
	fmt.Println("NaaJi ")
}
