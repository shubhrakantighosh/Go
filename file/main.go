package main

import (
	"fmt"
	"os"
)

func main() {
	defer DeleteFile()
	CreateFile()
	WriteFiles()
	ReadFile()

}

func CreateFile() {
	fileName, _ := os.Create("file/Omniful.txt")
	fileName.Close()
	fmt.Println(fileName.Name())
	fmt.Println("File created..")
}

func DeleteFile() {

	_, err := os.Stat("file/Omniful.txt")

	if err == nil {
		fmt.Println("Exists")

	} else {
		panic(err)
	}

	os.Remove("file/Omniful.txt")
	fmt.Println("File deleted.")

}

func WriteFiles() {

	_, err := os.Stat("file/Omniful.txt")

	if err == nil {
		fmt.Println("Exists")

	} else {
		fmt.Println("Not Exists")
	}

	str := []byte("Hi, User" + "\n" + "    How are you?")

	err = os.WriteFile("file/Omniful.txt", str, 0600)

	fmt.Println("Done.")

}

func ReadFile() {

	_, err := os.Stat("file/Omniful.txt")

	if err == nil {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not Exists")
	}

	fileName, _ := os.ReadFile("file/Omniful.txt")

	fmt.Println(string(fileName))

}
