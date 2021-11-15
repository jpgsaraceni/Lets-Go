// Existem dois times de futebol, o time amarelo e o time vermelho.
// O time amarelo tem 5 jogadores (Fernando, João, Lúcia, Mariana e Ana)
// e o time vermelho tem 4 jogadores (Helena, Jonas, José e Juliana).

// Crie um array de string para cada time e nomeie com o nome do time.
// Printe na tela os nomes dos jogadores de cada time

package main

import "fmt"

func main() {
	var (
		teamYellow [5]string
		teamRed    [4]string
	)

	teamYellow = [5]string{
		"Fernando",
		"João",
		"Lúcia",
		"Marina",
		"Ana",
	}

	teamRed = [4]string{
		"Helena",
		"Jonas",
		"José",
		"Juliana",
	}

	fmt.Printf("Time amarelo:\n%v\nTime vermelho:\n%v\n", teamYellow, teamRed)
}
