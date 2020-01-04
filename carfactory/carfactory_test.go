// carfactory_test.go tests the functions in carfactory

package carfactory

import (
	"fmt"
	"testing"
)

type CarMock struct {
	Name string
}

func (cm CarMock) Run() {
	fmt.Println("Mocking car " + cm.Name + " is running")
}

type CarFactoryMock struct{}

func (cf CarFactoryMock) MakeCar(name string) ICar {
	return CarMock{name}
}

func TestTransport(t *testing.T) {
	Transport(CarFactoryMock{})
}
