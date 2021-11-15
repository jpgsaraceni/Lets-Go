package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5, 6}
	var slice2 = []int{7, 8, 9, 10, 11, 12}

	var appendedSlice = append(slice1, slice2...)

	fmt.Println(appendedSlice)
}
