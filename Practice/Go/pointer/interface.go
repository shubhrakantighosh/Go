package main

import "fmt"

type Person interface {
	Speak()
	Walk()
}

type Student struct {
	Name       string
	SchoolName string
}

type Employee struct {
	CompanyName    string
	CodingLanguage string
}

func (s Student) Speak() {
	fmt.Println("ABCD..")
}

func (s Student) Walk() {
	fmt.Println("1 KM")
}

func (e Employee) Speak() {
	fmt.Println("Haan Ji")
}

func (e Employee) Walk() {
	fmt.Println("10 KM")
}

func main() {
	s := Student{"student", "Masai"}
	s.Speak()
	s.Walk()

	e := Employee{"Omniful", "Go"}
	e.Speak()
	e.Walk()
}
