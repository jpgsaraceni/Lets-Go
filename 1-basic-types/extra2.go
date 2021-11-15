package main

import (
	"fmt"
	"time"
)

func main() {
	var yearOfBirth int
	fmt.Println("Informe o ano de nascimento")
	fmt.Scan(&yearOfBirth)
	age := time.Now().Year() - yearOfBirth

	fmt.Printf("Idade: %v\n", age)
}
