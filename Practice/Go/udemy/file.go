package main

import (
	"fmt"
	"os"
)

func main() {
	FileCreate()
}

func FileCreate() {
	file, err := os.Create("Omnifil.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	file.WriteString("hii")

	d, _ := os.ReadFile("Omnifil.txt")

	fmt.Println(d)

	os.RemoveAll("/Users/shubhrakantighosh/Desktop/Go/Omnifil.txt")

}
