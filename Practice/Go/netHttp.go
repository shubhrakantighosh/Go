package main

import (
	"fmt"
	"net/http"
)

func main() {

	NetHttp()

}

func NetHttp() {

	getRes, _ := http.Get("http://google.com")

	fmt.Println(getRes.StatusCode)
	fmt.Println(getRes.Status)

	getRes.Body.Close()

}
