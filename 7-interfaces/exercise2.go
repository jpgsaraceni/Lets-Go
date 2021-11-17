// Crie tipos que representam diferentes animais, com atributos
// que façam sentido para cada um deles

// Crie uma interface que descreve o comportamento de
// apresentar um animal com a seguinte assinatura: Apresenta()

// Cada animal saberá como se apresentar.
// Sendo assim, faça com que cada um dos tipos que você criou
// implemente o método Apresenta(), que deve printar uma frase
// apresentando o animal e seus atributosDemonstre que todos os
// tipos implementam a interface que você criou declarando uma
// slice de animais e percorrendo-a com um for range que, em todas as voltas,
//  chama o método Apresenta().

package main

import "fmt"

type Pikachu struct {
	VoltageInVolts int
}

func (p Pikachu) Says() {
	fmt.Printf("Pika pika! My shock is %dV.\n", p.VoltageInVolts)
}

type Charmander struct {
	TemperatureInCelsius int
}

func (c Charmander) Says() {
	fmt.Printf("Char! My tail is %dºC!\n", c.TemperatureInCelsius)
}

type Pokemon interface {
	Says()
}

func main() {
	pika := Pikachu{1000}
	charme := Charmander{250}

	pokedex := []Pokemon{pika, charme}

	for _, poke := range pokedex {
		poke.Says()
	}
}
