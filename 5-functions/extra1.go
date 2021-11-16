package main

import "fmt"

func main() {
	list, sum, err := ModifyNumbers(1, 2, 5, 10, 20, 22)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v, %v\n", list, sum)
}

func ModifyNumbers(numbersList ...int) ([]int, int, error) {
	var newNumbersList = make([]int, 0)
	var sum int
	for _, number := range numbersList {
		if number < 0 {
			return newNumbersList, 0, fmt.Errorf("Essa função aceita somente números naturais.")
		}
		if number%2 == 0 {
			newNumbersList = append(newNumbersList, number/2)
		}
		if number%2 != 0 {
			newNumbersList = append(newNumbersList, number*2)
		}
		sum += number
	}
	return newNumbersList, sum, nil
}
