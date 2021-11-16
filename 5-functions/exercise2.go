package main

import "fmt"

func main() {
	PrintHello("João")
}

func PrintHello(name string) {
	fmt.Printf("Hello %s!\n", name)
}
