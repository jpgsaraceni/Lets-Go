package main

import "fmt"

func main() {
	falseIsGreater := 3 > 4
	trueIsLess := 2 < 10
	trueIsEqual := 3 == 3
	falseIsDifferent := 5 != 5
	trueAnd := true && true

	fmt.Printf("Valor: %v; Tipo: %T\n", falseIsGreater, falseIsGreater)
	fmt.Printf("Valor: %v; Tipo: %T\n", trueIsLess, trueIsLess)
	fmt.Printf("Valor: %v; Tipo: %T\n", trueIsEqual, trueIsEqual)
	fmt.Printf("Valor: %v; Tipo: %T\n", falseIsDifferent, falseIsDifferent)
	fmt.Printf("Valor: %v; Tipo: %T\n", trueAnd, trueAnd)
}
