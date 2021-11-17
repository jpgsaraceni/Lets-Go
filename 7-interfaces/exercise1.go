package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius int
}

type Square struct {
	Side int
}

type GeometricShape interface {
	getArea() float64
}

func (c Circle) getArea() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}

func (s Square) getArea() float64 {
	return math.Pow(float64(s.Side), 2)
}

func printArea(shape GeometricShape) {
	fmt.Printf("Área: %.2f\n", shape.getArea())
}

func main() {
	circle := Circle{Radius: 2}
	square := Square{Side: 2}
	printArea(circle)
	printArea(square)
}
