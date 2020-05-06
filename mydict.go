package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExists = errors.New("The word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(key, word string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		d[key] = word
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(key, word string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = word
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
