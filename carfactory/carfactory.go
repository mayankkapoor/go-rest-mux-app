// carfactory.go creates a car factory to manufacture cars then "run" them.
// This package demonstrates how to test third party packages which use
// structs, and how to mock structs.

package carfactory

import "fmt"

type Car struct {
	Name string
}

func (c Car) Run() {
	fmt.Println("Real car " + c.Name + " is running")
}

type CarFactory struct{}

func (cf CarFactory) MakeCar(name string) Car {
	return Car{name}
}

func Transport(cf CarFactory) {
	car := cf.MakeCar("lamborghini")
	car.Run()
}
