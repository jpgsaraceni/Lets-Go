package main

import "fmt"

func main() {
	var hours int

	for hours < 3 {
		var minutes int
		for minutes < 60 {
			var seconds int
			for seconds < 60 {
				fmt.Printf("%.2d:%.2d:%.2d\n", hours, minutes, seconds)
				seconds++
			}
			minutes++
		}
		hours++
	}
}
