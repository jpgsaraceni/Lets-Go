package main

import "fmt"

func main() {
	var age int = 20

	switch {
	case age < 18:
		fmt.Println("Menor de idade")
	case age >= 18 && age < 60:
		fmt.Println("Maior de idade")
	case age >= 60:
		fmt.Println("Idoso")
	default:
		fmt.Println("Valor inválido")
	}
}
