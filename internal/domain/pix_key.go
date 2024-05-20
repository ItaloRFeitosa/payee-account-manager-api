package domain

import (
	"errors"
	"log/slog"
	"regexp"
	"strings"
)

const (
	CPFPixKeyType            = "CPF"
	CNPJPixKeyType           = "CNPJ"
	TelefonePixKeyType       = "TELEFONE"
	EmailPixKeyType          = "EMAIL"
	ChaveAleatoriaPixKeyType = "CHAVE_ALEATORIA"
)

// PixKey is a interface that wraps Brazilian Instant Payment Identification of account
type PixKey interface {
	// Type can be CPF CNPJ TELEFONE EMAIL CHAVE_ALEATORIA
	Type() string
	// Value underlying pix key
	Value() string
	// Formatted value of pix key
	String() string
}

var ErrInvalidPixKeyType = errors.New("invalid pix key type")

// NewPixKey create a instance of PixKey based on type
func NewPixKey(typ, value string) (PixKey, error) {
	switch typ {
	case CPFPixKeyType:
		return NewCPFPixKey(value)
	case CNPJPixKeyType:
		return NewCNPJPixKey(value)
	case TelefonePixKeyType:
		return NewTelefonePixKey(value)
	case EmailPixKeyType:
		return NewEmailPixKey(value)
	case ChaveAleatoriaPixKeyType:
		return NewChaveAleatoriaPixKey(value)
	default:
		return nil, ErrInvalidPixKeyType
	}
}

type CPFPixKey struct {
	CPF
}

func (CPFPixKey) Type() string {
	return CPFPixKeyType
}

var EmptyCPFPixKey CPFPixKey

func NewCPFPixKey(value string) (CPFPixKey, error) {
	cpf, err := NewCPF(value)
	if err != nil {
		return EmptyCPFPixKey, err
	}

	return CPFPixKey{CPF: cpf}, nil
}

type CNPJPixKey struct {
	CNPJ
}

func (CNPJPixKey) Type() string {
	return CNPJPixKeyType
}

var EmptyCNPJPixKey CNPJPixKey

func NewCNPJPixKey(value string) (CNPJPixKey, error) {
	cnpj, err := NewCNPJ(value)
	if err != nil {
		return EmptyCNPJPixKey, err
	}

	return CNPJPixKey{cnpj}, nil
}

type TelefonePixKey struct {
	value string
}

func (TelefonePixKey) Type() string {
	return TelefonePixKeyType
}

func (tp TelefonePixKey) Value() string {
	return tp.value
}

func (tp TelefonePixKey) String() string {
	return "+" + tp.Value()
}

var (
	EmptyTelefonePixKey TelefonePixKey

	ErrInvalidTelefone = errors.New("invalid telefone")

	TelefoneRegex = regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`)
)

func NewTelefonePixKey(value string) (TelefonePixKey, error) {
	if !TelefoneRegex.MatchString(value) {
		return EmptyTelefonePixKey, ErrInvalidTelefone
	}

	value = keepOnlyNumbers(value)

	withoutCountryCode := len(value) == 11
	if withoutCountryCode {
		value = "55" + value
	}

	return TelefonePixKey{value}, nil
}

type EmailPixKey struct {
	Email
}

func (EmailPixKey) Type() string {
	return EmailPixKeyType
}

func (e EmailPixKey) String() string {
	return e.Value()
}

var EmptyEmailPixKey EmailPixKey

func NewEmailPixKey(value string) (EmailPixKey, error) {
	email, err := NewEmail(value)
	if err != nil {
		return EmptyEmailPixKey, err
	}

	return EmailPixKey{email}, nil
}

var (
	EmptyChaveAleatoriaPixKey ChaveAleatoriaPixKey

	ErrInvalidChaveAleatoria = errors.New("invalid chave aleatoria")

	ChaveAleatoriaRegex = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

type ChaveAleatoriaPixKey struct {
	value string
}

func NewChaveAleatoriaPixKey(value string) (ChaveAleatoriaPixKey, error) {
	if !ChaveAleatoriaRegex.MatchString(value) {
		return EmptyChaveAleatoriaPixKey, ErrInvalidChaveAleatoria
	}

	value = strings.ToLower(value)

	return ChaveAleatoriaPixKey{value}, nil
}

func (tp ChaveAleatoriaPixKey) Type() string {
	return ChaveAleatoriaPixKeyType
}

func (tp ChaveAleatoriaPixKey) Value() string {
	return tp.value
}

func (tp ChaveAleatoriaPixKey) String() string {
	return tp.Value()
}

// RestorePixKey shoud be used to restore a instance of PixKey from database
func RestorePixKey(typ, value string) PixKey {
	key, err := NewPixKey(typ, value)

	if err != nil {
		slog.Warn("tempered pix key with invalid values",
			slog.String("pix_key_type", typ),
			slog.String("pix_key", value),
			slog.String("error", err.Error()))

		return restoredPixKey{typ, value}
	}

	return key
}

// restoredPixKey is an struct to wrap tempered pix key value, that means
// if pix key is changed from database to invalid state, it wont break api
type restoredPixKey struct{ typ, value string }

func (r restoredPixKey) Type() string {
	return r.typ
}

func (r restoredPixKey) Value() string {
	return r.value
}

func (r restoredPixKey) String() string {
	return r.Value()
}
