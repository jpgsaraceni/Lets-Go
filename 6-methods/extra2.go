// No nosso conteúdo, temos um exemplo de validação de CPF.
// Porém, esse se encontra incompleto. Complete o exemplo adicionando
//  o algoritmo de verificação dos dígitos verificadores,
// e faça com que o cpf formatado (e.g 111.111.110-30) seja aceito
//  também por meio de um construtor.

package main

import (
	"fmt"
	"os"
	"strconv"
)

type CPF struct {
	Input       string
	ParsedInput string
	IsValid     bool
}

func main() {
	var cpf CPF

	fmt.Print("Insira um cpf para validar\n> ")
	fmt.Fscanf(os.Stdin, "%s", &cpf.Input)

	// cpf.Input = "111.444.777-35"

	err1 := cpf.parse()

	if err1 != nil {
		fmt.Println("Entrada inválida! Digite o CPF no formato 123.456.789-01 ou 12345678901")
		return
	}

	err2 := cpf.checkDigitsMatch()

	if err2 != nil {
		fmt.Println("Entrada inválida.")
		return
	}

	if cpf.IsValid {
		fmt.Println("CPF válido!")
		return
	}
	fmt.Println("CPF inválido!")
}

func (cpf *CPF) checkDigitsMatch() error {

	verifiedCpf, err := cpf.getVerifyingDigits()

	if err != nil {
		fmt.Println("Entrada inválida.")
		return fmt.Errorf("Error calculating digits.")
	}

	if verifiedCpf == cpf.ParsedInput {
		cpf.IsValid = true
	}
	return nil
}

func (cpf *CPF) parse() error {
	var parsed string

	if len(cpf.Input) != 11 || len(cpf.Input) != 14 {
		return fmt.Errorf("Invalid input.")

	}

	for i, character := range cpf.Input {
		if (character >= 48) && (character <= 57) { // 0-9
			parsed += string(character)
			continue
		}
		if len(cpf.Input) == 14 { // XXX.XXX.XXX-XX format
			if (character == 46) && (i == 3 || i == 7) { // .
				continue
			}
			if (character == 45) && (i == 11) { // -
				continue
			}
		}
		if len(cpf.Input) != 11 {
			return fmt.Errorf("Invalid input.")
		}
		return fmt.Errorf("Invalid input.")
	}
	cpf.ParsedInput = parsed
	return nil
}

func (cpf CPF) getVerifyingDigits() (string, error) {

	var cpfString = cpf.ParsedInput[0:9]

	cpfStringWithFirstDigit, err1 := cpf.iterateDigits(cpfString)

	if err1 != nil {
		return "", fmt.Errorf("Error calculating first digit")
	}

	cpfStringWithSecondDigit, err2 := cpf.iterateDigits(cpfStringWithFirstDigit)

	if err2 != nil {
		return "", fmt.Errorf("Error calculating second digit")
	}
	return cpfStringWithSecondDigit, nil
}

func (cpf CPF) iterateDigits(cpfString string) (string, error) {
	var sum int
	var factor int = len(cpfString) + 1

	for i := 0; i < len(cpfString); i++ {
		var char string = string(cpfString[i])
		digit, err := strconv.Atoi(char)
		sum += digit * factor
		factor--

		if err != nil {
			return "", fmt.Errorf("Error calculating second digit.")
		}
	}
	if sum%11 < 2 {
		return (cpfString + "0"), nil
	}
	return (cpfString + strconv.Itoa(11-sum%11)), nil
}
