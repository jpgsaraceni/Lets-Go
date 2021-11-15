// Considerando os times do item anterior, crie uma slice para representar cada um.

// Adicione o jogador Luis Inácio no time vermelho usando a função append()
// Printe na tela os nomes dos jogadores do time vermelho

package main

import "fmt"

func main() {
	var (
		teamRed []string
		// teamYellow []string
	)

	teamRed = []string{
		"Fernando",
		"João",
		"Lúcia",
		"Marina",
		"Ana",
	}

	// teamYellow = []string{
	// 	"Helena",
	// 	"Jonas",
	// 	"José",
	// 	"Juliana",
	// }

	teamRed = append(teamRed, "Luis Inácio")

	fmt.Printf("Time vermelho: %v", teamRed)
}
