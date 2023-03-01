package main

import "fmt"

func main() {

	str := "Hi,Shubhra"

	var i int

	for i <= len(str) {

		fmt.Println(string(str[i]))

		i++

		if i == len(str) {
			break
		}

	}

}
