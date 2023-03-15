package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const portNumber = ":8080"

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/about":
		aboutHandler(w, r)
	default:
		http.Error(w, "Wrong URL", http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprint("Home Page"), http.StatusOK)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprint("About Page"), http.StatusOK)
}

func main() {
	
	http.HandleFunc("/", pathHandler)
	http.ListenAndServe(portNumber, nil)
	fmt.Println(c)

}
