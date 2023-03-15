package main

import (
	"fmt"
	"time"
)

func main() {
	go Ram()
	go Shyam()
	time.Sleep(time.Second)
}

func Ram() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("Ram ", i)
	}
}

func Shyam() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("Shyam ", i)
	}
}
