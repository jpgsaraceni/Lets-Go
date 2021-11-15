package main

import "fmt"

func main() {
	var array1 [3]int
	var array2 [3]int

	array1 = [3]int{1, 2, 3}
	array2 = array1

	fmt.Println(array2)
}
