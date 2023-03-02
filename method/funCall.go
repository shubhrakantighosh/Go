package main

import "fmt"

func main() {
	fmt.Println("Sum of a and b is : ", SumOfTwoNumber(2, 9))
	SayHello("Hi, User")
	Multiply(2, 6)
}

func SumOfTwoNumber(a int, b int) int {
	return a + b
}

func SayHello(str string) {
	fmt.Println(str)
}

func Multiply(a int, b int) {
	OddOrEven(a * b)
}

func OddOrEven(num int) {
	if num%2 == 0 {
		fmt.Println("Even Number : ", num)
	} else {
		fmt.Println("Odd Number : ", num)
	}
}
