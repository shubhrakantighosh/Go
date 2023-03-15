package main

import (
	"fmt"
	"time"
)

func main() {
	//TimeSleep()
	Check()
}

func TimeSleep() {
	fmt.Println("Started...")
	time.Sleep(3000000000)
	fmt.Println("Ended...")
}

func Check() {
	date := time.DateTime
	fmt.Println(date)
	fmt.Println(time.Month(12))

	t, _ := time.LoadLocation("India/kolkata")
	fmt.Println(t)

	fmt.Println(time.Now().Location())

	fmt.Println(time.Now().Format("01//YYYY"))

}
