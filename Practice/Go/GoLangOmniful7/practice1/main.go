package main

import "fmt"

func main() {
	fmt.Println(Divide(0, 1))
}

type Exception struct {
	message string
}

func (e Exception) Error() string {
	return "Wrong haan ji"
}

func Divide(a int, b int) (int, error) {

	if a == 0 || b == 0 {
		return 0, Exception{}
	} else {
		return a / b, nil
	}

}
