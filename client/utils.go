package client

import "fmt"

//numBool is a boolean value that converts between 0 and 1
type numBool bool

func (a numBool) MarshalJSON() ([]byte, error) {
	if a {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

func (a *numBool) UnmarshalJSON(data []byte) error {
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
