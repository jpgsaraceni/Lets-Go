package main

import "fmt"

func main() {
	var someNumber int
	someNumber = 5
	fmt.Println(someNumber)

	var someString string
	someString = "String content"
	fmt.Println(someString)

	var someBool bool
	someBool = someNumber > 0
	fmt.Println(someBool)

	var someOtherBool bool
	someOtherBool = someNumber < 0
	fmt.Println(someOtherBool)
}
