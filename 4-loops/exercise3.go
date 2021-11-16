package main

import "fmt"

func main() {
	var hours int

	for hours < 24 {
		var minutes int
		for minutes < 60 {
			fmt.Println(hours, minutes)
			minutes++
		}
		hours++
	}
}
