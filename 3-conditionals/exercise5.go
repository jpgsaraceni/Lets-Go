package main

import "fmt"

func main() {
	var hora int = 17
	var mensagem string

	switch {
	case hora < 5:
		mensagem = "madrugada"
	case hora >= 5 && hora < 12:
		mensagem = "manhã"
	case hora >= 12 && hora < 18:
		mensagem = "tarde"
	case hora >= 18:
		mensagem = "noite"
	}

	fmt.Println(mensagem)
}
