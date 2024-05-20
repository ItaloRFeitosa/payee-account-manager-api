package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Document is an interface that represents value objects for government or tax identification documents (Ex: CPF and CNPJ)
type Document interface {
	// Value returns the underlying value of document, without formating
	Value() string
	// String returns the underlying value of document formatted
	// If CNPJ, return as 00.000.000/0001-00
	// If CPF, return as 000.000.000-00
	String() string
}

var ErrInvalidDocument = errors.New("invalid document")

// NewDocument returns an new instance of Document value object, it can be CPF or CNPJ
// If value not match CPF or CNPJ rules, ErrInvalidDocument is returned
func NewDocument(v string) (Document, error) {
	var nilDocument Document

	if cpf, err := NewCPF(v); err == nil {
		return cpf, nil
	}

	if cnpj, err := NewCNPJ(v); err == nil {
		return cnpj, nil
	}

	return nilDocument, fmt.Errorf("%w: %s", ErrInvalidDocument, v)
}

type CPF struct {
	value string
}

// Value returns the underlying value of document, without formating
func (c CPF) Value() string {
	return c.value
}

var cpfFormatPattern = regexp.MustCompile(`([\d]{3})([\d]{3})([\d]{3})([\d]{2})`)

// String returns the underlying value of document formatted (Ex: 000.000.000-00)
func (c CPF) String() string {
	return cpfFormatPattern.ReplaceAllString(c.value, "$1.$2.$3-$4")
}

var (
	EmptyCPF CPF

	ErrInvalidCPF = errors.New("invalid cpf number")

	CPFRegex = regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`)
)

// NewCPF returns an new instance of CPF value object
func NewCPF(v string) (CPF, error) {
	cpf, err := cleanAndValidateCPF(v)
	if err != nil {
		return EmptyCPF, err
	}

	return CPF{cpf}, nil
}

func cleanAndValidateCPF(documentNumber string) (string, error) {
	if !CPFRegex.MatchString(documentNumber) {
		return "", ErrInvalidCPF
	}

	documentNumber = keepOnlyNumbers(documentNumber)

	if len(documentNumber) != 11 {
		return "", ErrInvalidCPF
	}

	var cpfDigits [11]int
	for i, char := range documentNumber {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return "", ErrInvalidCPF
		}
		cpfDigits[i] = digit
	}

	// Check if all CPF digits are the same (a common invalid CPF condition)
	allSame := true
	for i := 1; i < 11; i++ {
		if cpfDigits[i] != cpfDigits[i-1] {
			allSame = false
			break
		}
	}
	if allSame {
		return "", ErrInvalidCPF
	}

	// Validate the CPF using the algorithm
	sum := 0
	for i := 0; i < 9; i++ {
		sum += cpfDigits[i] * (10 - i)
	}
	remainder := sum % 11

	// Calculate the first verification digit
	expectedDigit1 := 11 - remainder
	if expectedDigit1 == 10 || expectedDigit1 == 11 {
		expectedDigit1 = 0
	}
	if cpfDigits[9] != expectedDigit1 {
		return "", ErrInvalidCPF
	}

	// Calculate the second verification digit
	sum = 0
	for i := 0; i < 10; i++ {
		sum += cpfDigits[i] * (11 - i)
	}
	remainder = sum % 11

	expectedDigit2 := 11 - remainder
	if expectedDigit2 == 10 || expectedDigit2 == 11 {
		expectedDigit2 = 0
	}
	if cpfDigits[10] != expectedDigit2 {
		return "", ErrInvalidCPF
	}

	return documentNumber, nil
}

type CNPJ struct {
	value string
}

// Value returns the underlying value of document, without formating
func (c CNPJ) Value() string {
	return c.value
}

var cnpjFormatPattern = regexp.MustCompile(`([\d]{2})([\d]{3})([\d]{3})([\d]{4})([\d]{2})`)

// String returns the underlying value of document formatted (Ex: 00.000.000/0001-00)
func (c CNPJ) String() string {
	return cnpjFormatPattern.ReplaceAllString(c.value, "$1.$2.$3/$4-$5")
}

var (
	EmptyCNPJ CNPJ

	ErrInvalidCNPJ = errors.New("invalid cnpj number")

	CNPJRegex = regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
)

// NewCNPJ returns an new instance of CNPJ value object
func NewCNPJ(v string) (CNPJ, error) {
	cnpj, err := cleanAndValidateCNPJ(v)
	if err != nil {
		return EmptyCNPJ, err
	}

	return CNPJ{cnpj}, nil
}

func cleanAndValidateCNPJ(documentNumber string) (string, error) {
	if !CNPJRegex.MatchString(documentNumber) {
		return "", ErrInvalidCNPJ
	}

	documentNumber = keepOnlyNumbers(documentNumber)

	if len(documentNumber) != 14 {
		return "", ErrInvalidCNPJ
	}

	var cnpjDigits [14]int
	for i, char := range documentNumber {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return "", ErrInvalidCNPJ
		}
		cnpjDigits[i] = digit
	}

	sum := 0
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 12; i++ {
		sum += cnpjDigits[i] * weights[i]
	}
	remainder := sum % 11

	expectedDigit1 := 11 - remainder
	if expectedDigit1 >= 10 {
		expectedDigit1 = 0
	}
	if cnpjDigits[12] != expectedDigit1 {
		return "", ErrInvalidCNPJ
	}

	sum = 0
	weights = append([]int{6}, weights...) // Include the first position for the second digit calculation
	for i := 0; i < 13; i++ {
		sum += cnpjDigits[i] * weights[i]
	}
	remainder = sum % 11

	expectedDigit2 := 11 - remainder
	if expectedDigit2 >= 10 {
		expectedDigit2 = 0
	}
	if cnpjDigits[13] != expectedDigit2 {
		return "", ErrInvalidCNPJ
	}

	return documentNumber, nil
}

// RestoredPixKey is an struct to wrap tempered document value, that means
// if document is changed from database to invalid state, it wont break api
type restoredDocument struct {
	value string
}

func (c restoredDocument) Value() string {
	return c.value
}

func (c restoredDocument) String() string {
	return c.value
}
