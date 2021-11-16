// Crie duas funções, uma que receba uma string e retorne uma nova string
// alterando cada caractere para o terceiro caractere seguinte, e outra que reverta esse processo.
// Exemplo, o caractere “a” viraria “d”.

package main

import "fmt"

func main() {
	fmt.Println(TransformLetters("abcdefghijklmnopqrstuvwxyz123.,?:"))
	fmt.Println(RollbackTransformLetters("defghijklmnopqrstuvwxyzabc123.,?:"))
}

func TransformLetters(receivedText string) string {
	var newText string
	for _, letter := range receivedText {
		if letter < 65 || letter > 122 || (letter > 90 && letter < 97) { // not a letter
			newText += string(letter)
		}
		if (letter >= 65 && letter <= 87) || (letter >= 97 && letter <= 119) { // A-W || a-w
			newText += string(letter + 3)
		}
		if (letter >= 88 && letter <= 90) || (letter >= 120 && letter <= 122) { // X, Y, Z || x, y, z
			newText += string(letter - 23) // A, B, C
		}
	}
	return newText
}

func RollbackTransformLetters(receivedText string) string {
	var newText string
	for _, letter := range receivedText {
		if letter < 65 || letter > 122 || (letter > 90 && letter < 97) { // not a letter
			newText += string(letter)
		}
		if (letter >= 68 && letter <= 90) || (letter >= 100 && letter <= 122) { // D-Z || d-z
			newText += string(letter - 3)
		}
		if (letter >= 65 && letter <= 67) || (letter >= 97 && letter <= 99) { // A, B, C || a, b, c
			newText += string(letter + 23) // X, Y, Z
		}
	}
	return newText
}
