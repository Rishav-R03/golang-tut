package main

import "fmt"

type Shape interface {
	Area() float64
	Scale(factor float64)
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Breadth * r.Length
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (c *Circle) Scale(factor float64) {
	c.Radius = c.Radius * factor
}

func (r *Rectangle) Scale(factor float64) {
	r.Length = r.Length * factor
	r.Breadth = r.Breadth * factor
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Rectangle{Length: 10, Breadth: 2},
	}
	fmt.Println("--- Before Scaling ---")
	for _, s := range shapes {
		fmt.Printf("Area: %.2f\n", s.Area())
	}

	fmt.Println("\n--- Scaling all shapes by 2x ---")
	for _, s := range shapes {
		s.Scale(2) // This modifies the actual Circle/Rectangle in the slice
		fmt.Printf("New Area: %.2f\n", s.Area())
	}
}
