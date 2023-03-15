package main

import (
	"fmt"
	_ "net/http"
)

func main() {
	//helloWorld()
	//dataType()
	//array()
	//forLoop()
	//whileLoop()
	//switchCase()
	//input()
	//primeOrNot()
	//parameter("hii")
	//parameter("Shubhra")
	//andOrStatement("Shubhra", 0)
	//output := retunType(101)
	//fmt.Println(output)

}

func helloWorld() {
	println("Hello World")
}

func dataType() {

	var integer int = 6
	var float float32 = 9.089
	var str string = "Shubhra"
	var boolean bool = true

	println(integer)
	println(float)
	println(str)
	println(boolean)

	var x = "kawasaki"
	var y string
	y = "ninja"
	println(x, y)

	var a, b int
	a = 9
	b = 8

	println(a, b)

	k := 8
	l := "300"

	println(k, l)

	java, c, python, ruby := true, 9, 7.0, "Ruby"
	println(java, c, python, ruby)

}

func array() {

	var arr [2]string
	arr[0] = "Shubhra"
	arr[1] = "Chirag"

	println(len(arr))

	var array [3]int

	array[0] = 1
	array[1] = 3

	println(array[2])
	println(len(array))

	arr1 := [5]string{"Hi", "Shubhra", "Kawasaki", "Ninja"}
	arr2 := [2]int{1}

	println(arr1[4], arr2[1])

	fmt.Print("Array ")

}

func forLoop() {

	arr := [4]int{1, 2, 3, 4}

	fmt.Print("Array")
	fmt.Print(arr[0])

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}

func whileLoop() {

	i := 0

	for i <= 100 {

		if i%2 == 0 {
			fmt.Println("Even Number : ", i)
		} else {
			fmt.Println("Odd Number : ", i)
		}

		i++

	}

}

func switchCase() {

	str := "shubhra"

	switch str {

	case "Shubhra":
		fmt.Print("Shubhra")

	case "Chirag":
		fmt.Println("Chirag")

	default:
		fmt.Print("Wrong Input")

	}

}

func input() {

	name := ""

	fmt.Print("Enter Your Name : ")
	fmt.Scanf("%s", &name)
	fmt.Println("Good Morning : ", name)

	var i, j int

	fmt.Println("Enter Two Numbers ")
	fmt.Scanf("%d", &i)
	fmt.Scanf("%d", &j)
	fmt.Print("Sum of Two Number is : ", (i + j))

}

func primeOrNot() {

	var integer int
	count := 0

	fmt.Print("Enter a Number : ")
	fmt.Scanf("%d", &integer)

	for i := 1; i <= integer; i++ {

		if integer%i == 0 {
			count++
		}

	}

	if count == 2 {
		fmt.Println("Prime Number")
	} else {
		fmt.Println("Non Prime ")
	}

}

func parameter(name string) {

	if name == "Shubhra" {
		fmt.Println("Shubhra")
	} else {
		fmt.Println("Wrong Input")
	}

}

func andOrStatement(name string, age int) {

	if name == "Shubhra" && age == 26 {
		fmt.Println("Shubhra")
	} else if name == "Shubhra" || age == 26 {
		fmt.Println("Shubhra Or 26")
	} else {
		fmt.Println("Wrong Inputs.")
	}

}

func retunType(integer int) string {

	if integer > 0 && integer <= 100 {
		return ("1 to 100 in this range")
	} else {
		return ("Not in rage")
	}

}
