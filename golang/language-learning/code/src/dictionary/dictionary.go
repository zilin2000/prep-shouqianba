package dictionary

import "errors"

var ErrNotFound = errors.New("this word is not in the dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}
