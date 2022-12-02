package dictionary

import "errors"

type Dictionary map[string]string

var KeyNotFoundError = errors.New("Not an existing key")

func (d Dictionary) Search(key string) (string, error) {
	value, existing := d[key]
	if !existing {
		return "", KeyNotFoundError
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) {
	// A map value is a pointer to a runtime.hmap structure
	d[key] = value

	//TODO: avoid overwrite map value in Add()
}
