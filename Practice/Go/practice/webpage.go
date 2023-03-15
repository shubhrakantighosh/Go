package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Staring server")
}
