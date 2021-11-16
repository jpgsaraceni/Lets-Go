package main

import "fmt"

type Apartment struct {
	Number     int
	Owner      string
	hasParking bool
}

func main() {
	apartment1 := Apartment{
		Number:     107,
		Owner:      "João",
		hasParking: true,
	}

	apartment2 := Apartment{
		Number:     108,
		Owner:      "Irineu",
		hasParking: false,
	}

	apartments := []Apartment{apartment1, apartment2}

	for _, apartment := range apartments {
		var parkingMessage string
		if apartment.hasParking == true {
			parkingMessage = "Possui vaga na garagem."
		} else {
			parkingMessage = "Não possui vaga na garagem."
		}
		fmt.Printf(
			"##Informações do apatamento %v##\nProprietário: %v\n%v\n\n",
			apartment.Number,
			apartment.Owner,
			parkingMessage,
		)
	}
}
