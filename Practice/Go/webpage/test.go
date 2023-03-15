package main

import (
	"fmt"
	"net/http"
)

const portNumber1 = ":8080"

type Router struct {
}

func funcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Hello, World")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact Page")
}

func (router Router) serverHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		funcHandler(w, r)
	} else if r.URL.Path == "/contact" {
		contactHandler(w, r)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {

	//var router Router
	//http.Handle("/", router)
	////
	//http.ListenAndServe(portNumber1, http.HandlerFunc(homeHandler))
	//http.HandleFunc("/", homeHandler)
	//http.ListenAndServe(portNumber1, nil)

	//http.ListenAndServe(portNumber1, http.HandleFunc(router))
}
