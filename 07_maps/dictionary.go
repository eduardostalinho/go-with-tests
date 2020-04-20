package main

const (
	ErrUnexistingWord = DictionaryErr("word not found in dictionary")
	ErrExistingWord   = DictionaryErr("word already exists in dictionary")
)

type DictionaryErr string

type Dictionary map[string]string

func (d Dictionary) Add(w, definition string) error {
	_, err := d.Search(w)
	switch err {
	case ErrUnexistingWord:
		d[w] = definition
		return nil
	case nil:
		return ErrExistingWord
	default:
		return err
	}
	return nil

}

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrUnexistingWord
	}
	return definition, nil
}

func (err DictionaryErr) Error() string {
	return string(err)
}
