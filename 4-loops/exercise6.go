package main

import "fmt"

func main() {
	shoppingList := []string{"Pão", "Vinho"}

	for i, product := range shoppingList {
		fmt.Println(i, product)
	}
}
