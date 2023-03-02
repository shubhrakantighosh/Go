package main

import (
	"fmt"
	"os"
)

func main() {
	defer FolderDelete()
	FolderCreate()
	FolderRename()
}

func FolderCreate() {

	os.Mkdir("folder/shubhra", os.ModePerm)
	fmt.Println("Folder created.")
}

func FolderRename() {
	_, err := os.Stat("folder/shubhra")

	if err == nil {
		fmt.Println("Exists")
	}

	os.Rename("folder/shubhra", "folder/Omniful")
	fmt.Println("Folder renamed")
}

func FolderDelete() {
	_, err := os.Stat("folder/Omniful")

	if err == nil {
		fmt.Println("Exists")
	}

	os.Remove("folder/Omniful")
	fmt.Println("Folder deleted.")
}
