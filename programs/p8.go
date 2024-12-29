package programs

import "fmt"

func Program8() {
	code := `
package main

import (
	"fmt"
	"math"
)

// Define the Shape interface
type Shape interface {
	area() float64
	perimeter() float64
}

// Square struct implementing Shape interface
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}

// Circle struct implementing Shape interface
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	// Create a square and circle
	s := Square{side: 5}
	c := Circle{radius: 3}

	// Calculate and display area and perimeter for both shapes
	fmt.Printf("Square: Area = %.2f, Perimeter = %.2f\n", s.area(), s.perimeter())
	fmt.Printf("Circle: Area = %.2f, Perimeter = %.2f\n", c.area(), c.perimeter())
}

`
	fmt.Println("Program 8 Code:")
	fmt.Println(code)
}
