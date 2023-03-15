package main

import (
	"github/shubhrakantighosh/customGo/helpers"
	"log"
)

func main() {

	var customMy helpers.Custom

	customMy.FirstName = "Shubhra"
	customMy.LastName = "Ghosh"

	log.Printf(customMy.FirstName, customMy.LastName)

}
