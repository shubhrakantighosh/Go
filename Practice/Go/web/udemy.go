package main

import (
	"fmt"
	"net/http"
)

func HaanJi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hiii")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello, World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf(string(n)))
	})

	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/9", HaanJi)
	http.ListenAndServe(":8080", nil)

}
