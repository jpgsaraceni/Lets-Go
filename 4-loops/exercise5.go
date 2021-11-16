package main

import "fmt"

func main() {
	numbers := []string{
		"Zero",
		"One",
		"Two",
		"Three",
		"Four",
	}

	for i, value := range numbers {
		fmt.Println(i, value)
	}
}
