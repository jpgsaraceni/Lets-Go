package main

import (
	"fmt"
	"os"
)

func main() {
	var number int

	for {
		fmt.Print("Insira um número par\n> ")
		fmt.Fscanf(os.Stdin, "%d", &number)
		if number != 0 && number%2 == 0 {
			fmt.Println("Obrigado amigo, você é um amigo.")
			return
		}
	}

}
