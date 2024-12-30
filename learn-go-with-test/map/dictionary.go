package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error)  {

	// before
	// if val, ok := d[word]; ok {
	// 	return val, nil
	// } else {
	// 	return "", NOT_FOUNT
	// }

	val, ok := d[word]
	if !ok {
		return "", ErrNotFound 
	}

	return val, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
			d[word] = definition
	case nil:
			return ErrWordExists
	default:
			return err
	}

	return nil
}
