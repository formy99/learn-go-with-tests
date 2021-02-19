package maps

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound          = errors.New("could not find the word you were looking for")
	ErrWordExists        = errors.New("cannot add word because it already exists")
	ErrWordDoseNotExists = errors.New("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists

	default:
		return err

	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoseNotExists
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
