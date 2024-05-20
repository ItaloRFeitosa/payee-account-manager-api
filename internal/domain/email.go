package domain

import (
	"errors"
	"regexp"
	"strings"
)

// Email is a value object to handle email address values
type Email struct {
	value string
}

// Value returns the undelying value of Email as string
func (e Email) Value() string {
	return e.value
}

var (
	EmptyEmail Email

	EmailRegex = regexp.MustCompile("^[a-z0-9+_.-]+@[a-z0-9.-]+$")

	ErrInvalidEmail = errors.New("invalid email address format")
)

// NewEmail create a new instance of Email value object, and validate against EmailRegex
func NewEmail(v string) (Email, error) {
	trimmedEmail := strings.TrimSpace(v)

	if !EmailRegex.MatchString(trimmedEmail) {
		return EmptyEmail, ErrInvalidEmail
	}

	return Email{trimmedEmail}, nil
}
