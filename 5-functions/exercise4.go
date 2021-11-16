package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print(WriteGreeting("João", time.Now().Hour()))
}

func WriteGreeting(name string, time int) string {
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
	return fmt.Sprintf("Olá, %s %s!\n", name, greeting)
}
