package main

import "fmt"

type Vehicle interface {
	HornSound()
	HeadLightColour()
}

type Car struct {
	Wheel  int
	Colour string
}

func (car Car) HeadLightColour() string {
	return "Pure White"
}

func (car Car) HornSound() string {
	return "Beep Beep"
}

type Bike struct {
	CompanyName string
	TopSpeed    int
}

func (bike Bike) HeadLightColour() string {
	return "Yellow"
}

func (bike Bike) HornSound() string {
	return "Beep"
}

func main() {

	car := Car{4, "Blue"}
	fmt.Println(car.HornSound())
	fmt.Println(car.HeadLightColour())

	bike := Bike{"BMW", 200}
	fmt.Println(bike.HornSound())
	fmt.Println(bike.HeadLightColour())

}
