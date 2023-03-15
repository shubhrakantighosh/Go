package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addTwoNumber(1, 2)
	fmt.Fprint(w, fmt.Sprintf("Sum of two number is %d ", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float64 = 100.0, 0.0
	value, err := divideValues(x, y)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, value)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	http.ListenAndServe(portNumber, nil)
}

func addTwoNumber(x, y int) int {
	return x + y
}

func divideValues(x, y float64) (float64, error) {

	if y <= 0 {
		return 0, errors.New("y can't be 0 or less")
	}
	return x / y, nil

}
