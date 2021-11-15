package main

import "fmt"

func main() {
	someMap := map[string]string{
		"Branco": "#ffffff",
		"Preto":  "#000000",
	}
	fmt.Println("Before delete:", someMap)

	delete(someMap, "Preto")

	fmt.Println("After delete", someMap)
}
