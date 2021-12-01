package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

var ErrAlreadyExists = errors.New("word already exists")

func (d Dictionary) Add(word, definition string) error {
	_, exists := d[word]

	if exists {
		return ErrAlreadyExists
	}

	d[word] = definition
	return nil
}
