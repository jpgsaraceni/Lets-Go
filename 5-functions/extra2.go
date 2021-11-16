package main

import "fmt"

func main() {
	fmt.Println(CountLetterAppearences("afdsfa", "a"))
}

func CountLetterAppearences(text string, letter string) int {
	var letterCount int

	for _, letterInText := range text {
		if string(letterInText) == letter {
			letterCount++
		}
	}

	return letterCount
}
