package main

import "fmt"

func main() {
	var yearOfBirth int = 1996
	var currentYear int = 2021

	if age := currentYear - yearOfBirth; age >= 16 {
		fmt.Println("Já pode votar")
		return
	}
	fmt.Println("Não pode votar ainda.")
}
