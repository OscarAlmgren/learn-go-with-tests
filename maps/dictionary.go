package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrAlreadyFound = errors.New("word already exists in map")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if !ok {
		d[word] = definition
	} else {
		return ErrAlreadyFound
	}
	return nil
}
