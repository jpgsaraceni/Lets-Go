package main

import (
	"fmt"
	"time"
)

func main() {
	var yearOfBirth int
	fmt.Scan(&yearOfBirth)
	age := time.Now().Year() - yearOfBirth

	fmt.Println(age)
}
