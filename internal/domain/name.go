package domain

import (
	"errors"
	"strings"
)

// Name is value object that represents a Person Name or Legal Name
type Name struct {
	value string
}

// Value returns the undelying value of Name as string
func (n Name) Value() string {
	return n.value
}

var (
	// EmptyName is zero value of Name, can help in assertions
	EmptyName Name

	ErrNameEmptyString      = errors.New("name cannot be empty")
	ErrNameLessThenTwoWords = errors.New("name must contain at least two words")
	ErrShortFirstName       = errors.New("first name must have at least two characters")
)

// NewName creates a new instance of Name value object, return error if name value is not valid
func NewName(v string) (Name, error) {
	trimmedName := strings.Join(strings.Fields(v), " ")

	if trimmedName == "" {
		return EmptyName, ErrNameEmptyString
	}

	words := strings.Split(trimmedName, " ")

	if len(words) < 2 {
		return EmptyName, ErrNameLessThenTwoWords
	}

	if len(words[0]) < 2 {
		return EmptyName, ErrShortFirstName
	}

	return Name{trimmedName}, nil
}
