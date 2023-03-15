package main

import (
	"log"
	"time"
)

var s = "Hiii"

func main() {
	// var s = "Red"
	// sayHello(s)

	user := User{
		FirstName: "Shubhra",
		LastName:  "Ghosh",
	}

	log.Println(user.FirstName)

	var user1 User = User{
		FirstName: "Chirag",
		Age:       25,
	}
	log.Println(user1.FirstName, user1.Age)

}

// var str1 string inside the package (In java public )
// var Str1 string outSide the package (In java private)

type User struct {
	FirstName   string
	LastName    string
	Age         int
	DateOfBirth time.Time
}

// func sayHello(s string) string {
// 	log.Println(s)
// 	return s
// }
