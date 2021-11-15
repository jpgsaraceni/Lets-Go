package main

import "fmt"

func main() {
	amExcitedToLearn := true
	amGoExpert := false
	willBecomeGoExpert := true

	haveWhatINeed := amExcitedToLearn && !amGoExpert && willBecomeGoExpert

	fmt.Print(haveWhatINeed)
}
