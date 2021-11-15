package main

import "fmt"

func main() {
	x := 2
	y := 3
	z := 4

	if x > y && x > z {
		fmt.Println("The largest number is x")
		return
	}
	if y > x && y > z {
		fmt.Println("The largest number is y")
		return
	}
	if z > x && z > y {
		fmt.Println("The largest number is z")
		return
	}

	fmt.Println("Invalid or repeated values.")
}
