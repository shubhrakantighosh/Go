package main

import "os"

func main() {
	Folder()
}

func Folder() {

	folder := os.Mkdir("/Users/shubhrakantighosh/Desktop/Go/udemy/Omniful", os.ModePerm)

	if folder != nil {
		folder.Error()
	}

	os.RemoveAll("/Users/shubhrakantighosh/Desktop/Go/udemy/Omniful")

}
