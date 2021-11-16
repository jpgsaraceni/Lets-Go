package main

import "fmt"

type Stack []int

func main() {
	var someStack Stack = Stack{1, 2, 3}
	fmt.Println(someStack)
	fmt.Println(someStack.pop())
	fmt.Println(someStack.push(5))
	someStack = someStack.pop()
	fmt.Println(someStack)
	someStack = someStack.push(1)
	fmt.Println(someStack)
}

func (s Stack) push(entry int) []int {
	return append(s, entry)
}

func (s Stack) pop() []int {
	return s[1:]
}
