package main

import (
	"fmt"
	"time"
)

func main() {
	SayHello("João", time.Now().Hour())
}

func SayHello(name string, time int) {
	var greeting string
	if time < 5 && time > 5 {
		greeting = "boa madrugada"
	}
	if time >= 5 && time < 12 {
		greeting = "bom dia"
	}
	if time >= 12 && time < 18 {
		greeting = "boa tarde"
	}
	if time >= 18 && time < 24 {
		greeting = "boa noite"
	}
	fmt.Printf("Olá, %s %s!\n", greeting, name)
}
