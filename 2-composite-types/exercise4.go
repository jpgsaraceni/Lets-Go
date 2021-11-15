package main

import "fmt"

func main() {
	shoppingList := []string{"Arroz", "Feijão"}
	fmt.Println(shoppingList)

	shoppingList = append(shoppingList, "Batata")
	fmt.Println(shoppingList)
}
