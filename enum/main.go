package main

import "fmt"

const (
	January   = 1
	February  = 2
	March     = 3
	April     = 4
	May       = 5
	June      = 6
	July      = 7
	August    = 8
	September = 9
	October   = 10
	November  = 11
	December  = 12
)

const (
	Small      int     = 1
	Medium             = "2"
	Large      string  = "3"
	ExtraLarge float64 = 4.98765
)

func main() {
	fmt.Println(January, December, March)
	fmt.Println(Small, Large, ExtraLarge)
}
