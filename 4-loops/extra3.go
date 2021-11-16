package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberString string
	for i := 1; i < 11; i++ {
		numberString = numberString + strconv.Itoa(i)
		fmt.Println(numberString)
	}
}
