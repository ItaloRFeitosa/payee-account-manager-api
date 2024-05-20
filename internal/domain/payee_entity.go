package domain

import (
	"errors"
	"log/slog"
)

type PayeeEntity struct {
	id          EntityID
	name        Name
	document    Document
	status      PayeeStatus
	email       Email
	pixKey      PixKey
	bankAccount *BankAccount
}

func (p *PayeeEntity) ID() string {
	return p.id.Value()
}

func (p *PayeeEntity) Name() string {
	return p.name.Value()
}

func (p *PayeeEntity) Document() Document {
	return p.document
}

func (p *PayeeEntity) Email() string {
	return p.email.Value()
}

func (p *PayeeEntity) Status() PayeeStatus {
	return p.status
}

func (p *PayeeEntity) PixKey() PixKey {
	return p.pixKey
}

func (p *PayeeEntity) BankAccount() *BankAccount {
	return p.bankAccount
}

// EditDetails updates payee information
// when payee status is VALID, only email can be changed
func (p *PayeeEntity) EditDetails(
	name string,
	document string,
	pixKeyType string,
	pixKey string,
	email string,
) error {
	var err error

	if email != "" {
		p.email, err = NewEmail(email)
		if err != nil {
			return err
		}
	} else {
		p.email = EmptyEmail
	}

	if p.status != PayeeDraftStatus {
		return nil
	}

	p.name, err = NewName(name)
	if err != nil {
		return err
	}

	p.document, err = NewDocument(document)
	if err != nil {
		return err
	}

	p.pixKey, err = NewPixKey(pixKeyType, pixKey)
	if err != nil {
		return err
	}

	return nil
}

// CreatePayee is a factory function to create a valid instance of PayeeEntity
func CreatePayee(
	name string,
	document string,
	pixKeyType string,
	pixKey string,
	email string,
) (*PayeeEntity, error) {
	var err error

	payee := new(PayeeEntity)

	payee.id = NewEntityID()

	payee.name, err = NewName(name)
	if err != nil {
		return nil, err
	}

	payee.document, err = NewDocument(document)
	if err != nil {
		return nil, err
	}

	payee.pixKey, err = NewPixKey(pixKeyType, pixKey)
	if err != nil {
		return nil, err
	}

	if email != "" {
		payee.email, err = NewEmail(email)
		if err != nil {
			return nil, err
		}
	}

	payee.status = PayeeDraftStatus

	return payee, nil
}

// RestorePayee is a factory function to restore a PayeeEntity from database
func RestorePayee(
	id string,
	name string,
	document string,
	status string,
	email string,
	pixKeyType string,
	pixKeyValue string,
	bankAccount *BankAccount,
) *PayeeEntity {

	payeeStatus, err := restorePayeeStatus(status)
	if errors.Is(err, ErrTemperedValue) {
		slog.Warn("tempered payee status", slog.String("payee_id", id))

		payeeStatus = PayeeStatus{status, status}
	}

	pixKey, err := NewPixKey(pixKeyType, pixKeyValue)
	if err != nil {
		slog.Warn("tempered pix key with invalid values", slog.String("payee_id", id), slog.String("error", err.Error()))

		pixKey = restoredPixKey{pixKeyType, pixKeyValue}
	}

	payeeDocument, err := NewDocument(document)
	if err != nil {
		slog.Warn("tempered document with invalid values", slog.String("payee_id", id), slog.String("error", err.Error()))

		payeeDocument = restoredDocument{document}
	}

	return &PayeeEntity{
		id:          EntityID{id},
		name:        Name{name},
		document:    payeeDocument,
		status:      payeeStatus,
		email:       Email{email},
		pixKey:      pixKey,
		bankAccount: bankAccount,
	}
}
