package main

import "fmt"

func main() {
	age := 25

	if age < 18 {
		fmt.Println("Menor de idade")
		return
	}
	if age >= 18 && age < 60 {
		fmt.Println("Maior de idade")
		return
	}
	fmt.Println("Idoso")
}
