package main

import "fmt"

func main() {
	number := 1
	if number > 0 {
		fmt.Println("Positivo")
		return
	}
	fmt.Println("Negativo")
}
