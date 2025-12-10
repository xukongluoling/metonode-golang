package main

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 长方形
type Rectangle struct {
	length, width, height float64
}

func (rect Rectangle) Area() float64 {
	return rect.length * rect.width * rect.height
}
func (rect Rectangle) Perimeter() float64 {
	return rect.length + rect.width + rect.height
}

// Circle 圆
type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (circle Circle) Perimeter() float64 {
	return math.Pi * circle.radius * 2
}
