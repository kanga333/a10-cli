package client

import "fmt"

//NumBool is a boolean value that converts between 0 and 1
type NumBool bool

func (a NumBool) MarshalJSON() ([]byte, error) {
	if a {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

func (a *NumBool) UnmarshalJSON(data []byte) error {
	asString := string(data)
	switch asString {
	case "0":
		*a = false
		return nil
	case "1":
		*a = true
		return nil
	case "false":
		*a = false
		return nil
	case "true":
		*a = true
		return nil
	default:
		return fmt.Errorf("uncastable value: %s", asString)
	}
}
