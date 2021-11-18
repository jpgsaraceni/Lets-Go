package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const englishDefaultName = "World"
const spanishDefaultName = "Mundo"
const frenchDefaultName = "Monde"

func greetingPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func defaultName(language string) string {
	switch language {
	case "Spanish":
		return spanishDefaultName
	case "French":
		return frenchDefaultName
	default:
		return englishDefaultName
	}
}

func Hello(name, language string) string {
	prefix := greetingPrefix(language)

	if name == "" {
		name = defaultName(language)
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("João", "English"))
}
