package domain

import "fmt"

type typedError interface {
	error
	Code() string
}

type baseTypedError struct {
	errCode string
}

func (te baseTypedError) Code() string {
	return te.errCode
}

func (te baseTypedError) Error() string {
	return fmt.Sprintf("Error of type \"%s\" happened", te.errCode)
}

type outOfRangeError struct {
	typedError
	rangeMin int
	rangeMax int
	value    int
}

func NewOutOfRangeErr(rangeMin, rangeMax, value int) *outOfRangeError {
	return &outOfRangeError{
		baseTypedError{"Value out of range"},
		rangeMin,
		rangeMax,
		value,
	}
}

func (re *outOfRangeError) Error() string {
	return fmt.Sprintf("Value %d is out of range %d - %d", re.value, re.rangeMin, re.rangeMax)
}
