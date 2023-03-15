package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path == "/" {
	//	home
	//}
}

func main() {

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":3000", nil)
}
