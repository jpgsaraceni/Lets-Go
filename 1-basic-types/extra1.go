package main

import (
	"fmt"
)

func main() {
	// var quilometros int8 // will throw error because int 8 only supports number ranging
	// from -128 to 127.
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}
