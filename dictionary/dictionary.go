package dictionary

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrorKeyNotFound     = DictionaryError("Not an existing key")
	ErrorKeyExisting     = DictionaryError("Key existing")
	ErrorWordNotExisting = DictionaryError("Word not existing")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, existing := d[key]
	if !existing {
		return "", ErrorKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	// A map value is a pointer to a runtime.hmap structure
	_, err := d.Search(key)
	switch err {
	case ErrorKeyNotFound:
		d[key] = value
	case nil:
		return ErrorKeyExisting
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorKeyNotFound:
		return ErrorWordNotExisting
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
