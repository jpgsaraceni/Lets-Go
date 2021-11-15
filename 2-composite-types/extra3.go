// Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.

// Passos:

// Criar um mapa com chave do tipo string e valor do tipo string (map[string]string)
// onde a chave seja o nome da cidade e o valor o nome do país.

// Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país.
// Esse mapa segundo mapa terá tipo map[string]int,
// sendo a chave o nome do país e o valor a quantidade de menções.

// Printe na tela os valores do último mapa.

package main

import "fmt"

func main() {

	citiesLived := map[string]string{
		"Petrópolis":      "Brasil",
		"Rio de Janeiro":  "Brasil",
		"Cabo Frio":       "Brasil",
		"Toronto":         "Canadá",
		"Innisfil":        "Canadá",
		"Paty do Alferes": "Brasil",
		"Arraial do Cabo": "Brasil",
		"Macaé":           "Brasil",
		"Miguel Pereira":  "Brasil",
		"Volta Redonda":   "Brasil",
	}

	var citiesPerCountry = make(map[string]int)

	for _, country := range citiesLived {
		citiesPerCountry[country] += 1
	}

	fmt.Printf("%v", citiesPerCountry)
}
