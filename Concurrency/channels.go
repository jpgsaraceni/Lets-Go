package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) { // cahn<- means this channel is send-only.
	for i := 0; ; i++ {
		c <- "ping" // sends "ping" to channel, which is then blocked.
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong" // waits until c is unblocked to send "pong"
	}
}

func printer(c <-chan string) { // <-chan means this channel is receive-only.
	for {
		msg := <-c // receives the message from c, which is then unblocked.
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c) // these three goroutines will work synchronously because they share the same channel

	var input string
	fmt.Scanln(&input)
}
