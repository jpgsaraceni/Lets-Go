// Construa dois métodos: um que retorna a área e
// outro que retorna o perímetro de uma estrutura que representa um círculo.
// Printe a área e o perímetro de um círculo.

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func main() {
	circle := Circle{
		Radius: 5,
	}
	fmt.Printf("Área: %.2f\nPerímetro: %.2f\n", circle.GetArea(), circle.GetPerimeter())
}

func (c Circle) GetArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) GetPerimeter() float64 {
	return c.Radius * 2 * math.Pi
}
