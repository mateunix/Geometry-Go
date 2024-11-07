package main

import (
	"fmt"
	"math"
)

type gshape interface {
	area() float64
	perimeter() float64
}
type square struct {
	side float64
}
type rectangle struct {
	length float64
	height float64
}
type circle struct {
	radius float64
}
type etriangle struct { //equilateral triangle
	length float64
	height float64
}
type trapezium struct {
	base   float64
	top    float64
	height float64
	side1  float64
	side2  float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

func (r rectangle) area() float64 {
	return r.length * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.length + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (e etriangle) area() float64 {
	return e.length * e.height / 2
}

func (e etriangle) perimeter() float64 {
	return 3 * e.length
}

func (t trapezium) area() float64 {
	return 0.5 * t.height * (t.base + t.top)
}

func (t trapezium) perimeter() float64 {
	return t.base + t.top + t.side1 + t.side2
}

func main() {
	var shapeType string

	// Ask user for the type of shape
	fmt.Print("Please enter the geometric shape (square, rectangle, circle, etriangle, trapezium) that you want to calculate: ")
	fmt.Scanf("%s", &shapeType)

	// Handle input based on shape type
	switch shapeType {
	case "square":
		var side float64
		fmt.Print("Enter side length of the square: ")
		fmt.Scanf("%f", &side)
		sq := square{side: side}
		fmt.Printf("Area of square: %.2f\n", sq.area())
		fmt.Printf("Perimeter of square: %.2f\n", sq.perimeter())

	case "rectangle":
		var length, height float64
		fmt.Print("Enter length of the rectangle: ")
		fmt.Scanf("%f", &length)
		fmt.Print("Enter height of the rectangle: ")
		fmt.Scanf("%f", &height)
		r := rectangle{length: length, height: height}
		fmt.Printf("Area of rectangle: %.2f\n", r.area())
		fmt.Printf("Perimeter of rectangle: %.2f\n", r.perimeter())

	case "circle":
		var radius float64
		fmt.Print("Enter radius of the circle: ")
		fmt.Scanf("%f", &radius)
		c := circle{radius: radius}
		fmt.Printf("Area of circle: %.2f\n", c.area())
		fmt.Printf("Perimeter (circumference) of circle: %.2f\n", c.perimeter())

	case "etriangle":
		var length, height float64
		fmt.Print("Enter length of the equilateral triangle: ")
		fmt.Scanf("%f", &length)
		fmt.Print("Enter height of the equilateral triangle: ")
		fmt.Scanf("%f", &height)
		e := etriangle{length: length, height: height}
		fmt.Printf("Area of equilateral triangle: %.2f\n", e.area())
		fmt.Printf("Perimeter of equilateral triangle: %.2f\n", e.perimeter())

	case "trapezium":
		var base, top, height, side1, side2 float64
		fmt.Print("Enter base of the trapezium: ")
		fmt.Scanf("%f", &base)
		fmt.Print("Enter top of the trapezium: ")
		fmt.Scanf("%f", &top)
		fmt.Print("Enter height of the trapezium: ")
		fmt.Scanf("%f", &height)
		fmt.Print("Enter side1 of the trapezium: ")
		fmt.Scanf("%f", &side1)
		fmt.Print("Enter side2 of the trapezium: ")
		fmt.Scanf("%f", &side2)
		t := trapezium{base: base, top: top, height: height, side1: side1, side2: side2}
		fmt.Printf("Area of trapezium: %.2f\n", t.area())
		fmt.Printf("Perimeter of trapezium: %.2f\n", t.perimeter())

	default:
		fmt.Println("Unknown shape type.")
	}
}
