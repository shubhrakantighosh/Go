package main

import (
	"fmt"
	"math/rand"
	"regexp"
)

func main() {

	str := "ILoveMyCountry"

	var txt []byte = []byte("ILove")

	match1, err1 := regexp.MatchString("Love", str)
	fmt.Println(match1, " : ", err1)

	match2, err2 := regexp.Match("ILove", txt)

	fmt.Println(match2, " : ", err2)

	fmt.Println(rand.Int())
}
