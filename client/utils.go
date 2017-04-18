package client

import "fmt"

//A10Bool is a boolean value that converts between 0 and 1
type a10Bool bool

func (a a10Bool) MarshalJSON() ([]byte, error) {
	if a {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

func (a *a10Bool) UnmarshalJSON(data []byte) error {
	asString := string(data)
	switch asString {
	case "0":
		*a = false
		return nil
	case "1":
		*a = true
		return nil
	default:
		return fmt.Errorf("Uncastable value: %v", asString)
	}
}
