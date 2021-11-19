package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Millisecond * 200)   // sends to channel every 200ms
	boom := time.After(time.Millisecond * 1000) // send to channel once after 1000ms
	for {
		select {
		case <-tick:
			fmt.Println("Tick") // will run when channel is ready
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("...") // will run when no channel is ready
			time.Sleep(time.Millisecond * 50)
		}
	}
}
