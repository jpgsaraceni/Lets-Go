package main

import "fmt"

type NumberList struct {
	Numbers []int
}

func main() {
	list1 := NumberList{
		Numbers: []int{1, 2, 3, 4},
	}
	fmt.Printf("Soma: %v\nMédia: %.2f\n", list1.GetSum(), list1.GetAvarage())
}

func (list NumberList) GetAvarage() float64 {
	var sum = list.GetSum()
	return float64(sum) / float64(len(list.Numbers))
}

func (list NumberList) GetSum() int {
	var sum int
	for _, number := range list.Numbers {
		sum += number
	}
	return sum
}
