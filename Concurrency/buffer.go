package main

import "fmt"

func main() {
	ch := make(chan int, 2) // sends block when channel is full and blocks receives when empty
	ch <- 1
	ch <- 2
	// ch <- 3 // deadlock because the buffered channel is full
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // deadlock because the buffered channel is empty
}
