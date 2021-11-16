package main

import "fmt"

func main() {
	shoppingList := []string{"Pão", "Vinho"}

	for i := 0; i < len(shoppingList); i++ {
		fmt.Println(i, shoppingList[i])
	}
}
