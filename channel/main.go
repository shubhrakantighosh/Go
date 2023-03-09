package main

import "fmt"

func main() {

	channel := make(chan int)
	go SayHello(channel)
	fmt.Println("Main")
	receive := <-channel
	fmt.Println(receive)

}

func SayHello(i chan int) {
	fmt.Println("Say Hello")
	i <- 9
}
