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

// couldn't run this code because The Go Playground doesn't support input
