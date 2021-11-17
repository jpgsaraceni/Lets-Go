// Dada a seguinte função, o que devemos escrever para que ela printe
// “inteiro” caso o parâmetro recebido seja int32 e
// “ponto flutuante” caso o parâmetro seja float32?
// Caso o tipo não seja reconhecido, a função deverá retornar um erro informando o ocorrido.

package main

import "fmt"

// func printType(i interface{}) error {
// 	// lógica aqui
// }

// func main() {
// 	printType(int32(2))
// 	printType(float32(2.23))
// }

type Number interface {
	Type() string
}

type Integer struct {
	number int32
}

func (i Integer) Type() string {
	return "Inteiro"
}

type Float struct {
	number float32
}

func (f Float) Type() string {
	return "Ponto flutuante"
}

func printType(n Number) error {
	if n.Type() == "Inteiro" || n.Type() == "Ponto flutuante" {
		fmt.Println(n.Type())
		return nil
	}
	return fmt.Errorf("Not int32 or float32")
}

func main() {
	integer := Integer{1}
	float := Float{1.1}
	printType(integer)
	printType(float)
}
