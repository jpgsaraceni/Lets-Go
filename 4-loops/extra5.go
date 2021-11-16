package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var letterCounter = map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	var text string

	fmt.Print("Insira um texto:\n> ")
	fmt.Fscanf(os.Stdin, "%s", &text)

	for _, letter := range text {
		letterCounter[string(letter)] += 1
	}

	lettersKeys := make([]string, 0, len(letterCounter))

	for k := range letterCounter {
		lettersKeys = append(lettersKeys, k)
	}

	sort.Strings(lettersKeys)

	for _, key := range lettersKeys {
		switch letterCounter[key] {
		case 0:
			fmt.Printf("%s => não aparece no texto.\n", key)
		case 1:
			fmt.Printf("%s => aparece 1 vez no texto.\n", key)
		default:
			fmt.Printf("%s => aparece %d vezes no texto.\n", key, letterCounter[key])
		}
	}
}
