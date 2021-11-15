package main

import "fmt"

func main() {
	var banana, cerveja, abacate, salgadinho float32

	banana = 1.25
	cerveja = 2.98
	abacate = 4.59
	salgadinho = 7.29

	var quantidadeBanana, quantidadeCerveja, quantidadeAbacate, quantidadeSalgadinho float32

	quantidadeBanana = 2.17
	quantidadeCerveja = 6
	quantidadeAbacate = 5.65
	quantidadeSalgadinho = 3

	var precoBanana, precoCerveja, precoAbacate, precoSalgadinho, precoTotal float32

	precoBanana = banana * quantidadeBanana
	precoCerveja = cerveja * quantidadeCerveja
	precoAbacate = abacate * quantidadeAbacate
	precoSalgadinho = salgadinho * quantidadeSalgadinho
	precoTotal = precoBanana + precoCerveja + precoAbacate + precoSalgadinho

	fmt.Printf("Preço unitário banana: %v. Peso: %v. Valor total: %v\n", banana, quantidadeBanana, precoBanana)
	fmt.Printf("Preço unitário cerveja: %v. Peso: %v. Valor total: %v\n", cerveja, quantidadeCerveja, precoCerveja)
	fmt.Printf("Preço unitário abacate: %v. Peso: %v. Valor total: %v\n", abacate, quantidadeAbacate, precoAbacate)
	fmt.Printf("Preço unitário salgadinho: %v. Peso: %v. Valor total: %v\n", salgadinho, quantidadeSalgadinho, precoSalgadinho)
	fmt.Printf("Valor total: %v", (precoTotal))
}
