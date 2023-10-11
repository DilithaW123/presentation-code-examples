package main

import "fmt"

type Vehicle struct {
	NumberOfWheels int
	Brand          string
}

// Vehicle object is copied onto the stack, not the same object
func (v Vehicle) editWheels(n int) {
	v.NumberOfWheels = n
}

func (v *Vehicle) editWheelsPointer(n int) {
	v.NumberOfWheels = n
}

func editingWheels(v *Vehicle, n int) {
	v.NumberOfWheels = n
}

func main() {
	car := Vehicle{4, "Ford"}
	println(fmt.Sprintf("Car has %d wheels", car.NumberOfWheels))

	car.editWheels(3)
	println(fmt.Sprintf("Car has %d wheels", car.NumberOfWheels))

	car.editWheelsPointer(2)
	println(fmt.Sprintf("Car has %d wheels", car.NumberOfWheels))

	editingWheels(&car, 1)
	println(fmt.Sprintf("Car has %d wheels", car.NumberOfWheels))
}
