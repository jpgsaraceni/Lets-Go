package main

import "fmt"

type Sale struct {
	Value int
	Name  string
	Date  string
}

func main() {
	sale1 := Sale{
		Value: 10,
		Name:  "João",
		Date:  "2021-11-01",
	}

	sale2 := Sale{
		Value: 5,
		Name:  "João",
		Date:  "2021-11-02",
	}
	sale3 := Sale{
		Value: 30,
		Name:  "Luana",
		Date:  "2021-11-01",
	}

	salesBySalesperson, totalSales, err := FilterBySalesperson(sale1, sale2, sale3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(salesBySalesperson, totalSales)

}

func FilterBySalesperson(salesList ...Sale) (map[string]int, int, error) {
	var totalSold int
	var salesFiltered = make(map[string]int)
	for _, sale := range salesList {
		if sale.Value == 0 || sale.Name == "" {
			return make(map[string]int), 0, fmt.Errorf("Missing sales value or name.")
		}
		salesFiltered[sale.Name] += sale.Value
		totalSold += sale.Value
	}

	return salesFiltered, totalSold, nil
}
