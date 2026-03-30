package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Square implements shape
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Engine struct {
	HorsePower int
	EngineType string
}

type Car struct {
	Engine
	Model string
}

type ElectricCar struct {
	Engine
	BatteryPower float64
}

func (ec ElectricCar) start() {
	fmt.Println("Electric engine started...")
}

func (e Engine) start() {
	fmt.Println("Engine started...")
}

func main() {
	r1 := Rectangle{Height: 1.0, Width: 5.0}
	c1 := Circle{Radius: 5.0}

	car1 := Car{
		Engine: Engine{},
		Model:  "hyundai",
	}
	ec1 := ElectricCar{Engine: Engine{HorsePower: 1000, EngineType: "Battery"}, BatteryPower: 85.5}
	ec1.start()
	car1.start()
	r1.Area()
	c1.Area()
}
