package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	user := User{"Shubhra", "Ghosh", 26}

	fmt.Println(user.FirstName, user.LastName, user.Age)

}
