// Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
// Crie três variáveis do tipo pessoa.
// Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas.

package main

import "fmt"

type Pessoa struct {
	Nome         string
	Idade        int
	CorPreferida string
}

func main() {
	joao := Pessoa{
		Nome:         "João",
		Idade:        25,
		CorPreferida: "Azul",
	}
	lina := Pessoa{
		Nome:         "Lina",
		Idade:        6,
		CorPreferida: "Rosa",
	}
	luana := Pessoa{
		Nome:         "Luana",
		Idade:        24,
		CorPreferida: "Vermelho",
	}

	fmt.Printf("Um pouco sobre %v, %v e %v:\n", joao.Nome, lina.Nome, luana.Nome)
	fmt.Printf("%v tem %v anos e sua cor preferida é %v\n", joao.Nome, joao.Idade, joao.CorPreferida)
	fmt.Printf("%v tem %v anos e sua cor preferida é %v\n", lina.Nome, lina.Idade, lina.CorPreferida)
	fmt.Printf("%v tem %v anos e sua cor preferida é %v\n", luana.Nome, luana.Idade, luana.CorPreferida)
}
