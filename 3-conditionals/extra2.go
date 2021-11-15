package main

import "fmt"

func main() {
	x := 6
	var divisors = make([]int, 0)

	if x == 0 {
		fmt.Println("Zero.")
		return
	}
	if x%2 == 0 {
		divisors = append(divisors, 2)
	}
	if x%3 == 0 {
		divisors = append(divisors, 3)
	}

	switch len(divisors) {
	case 0:
		fmt.Println("O número não é múltipo de 2 ou 3.")
		return
	case 2:
		fmt.Println("O número é múltiplo de 2 e 3.")
		return
	case 1:
		fmt.Printf("O número é múltiplo de %v", divisors[0])
		return
	}

}
