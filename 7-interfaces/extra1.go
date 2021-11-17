package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Printer interface {
	PrintContent()
}

type File struct {
	file os.File
}

func (f *File) PrintContent() {
	content, _ := ioutil.ReadAll(&f.file)
	fmt.Printf("%s\n", string(content))
}

type Text struct {
	text string
}

func (t Text) PrintContent() {
	fmt.Printf("%s\n", t.text)
}

func main() {
	myFile, _ := os.OpenFile("README.md", os.O_RDONLY, os.ModeDevice)
	file := File{*myFile}
	file.PrintContent()

	myString := Text{"this is my content"}
	myString.PrintContent()
}
