package domain_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/italorfeitosa/payee-account-manager-api/internal/domain"
	"github.com/italorfeitosa/payee-account-manager-api/test/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testPayeeMaxIterations = 100

func TestCreatePayee_Success(t *testing.T) {
	for i := 0; i < testPayeeMaxIterations; i++ {
		t.Run("should create a valid payee", func(t *testing.T) {
			wantName := gofakeit.Name()
			wantDocument := gofakeit.RandomString([]string{fake.CNPJ(), fake.CPF()})
			wantEmail := gofakeit.RandomString([]string{gofakeit.Email(), ""})
			wantPixKeyType, wantPixKey := fake.PixKey()

			got, err := domain.CreatePayee(
				wantName,
				wantDocument,
				wantPixKeyType,
				wantPixKey,
				wantEmail,
			)

			require.NoError(t, err)

			assert.Condition(t, func() (success bool) {
				_, err := uuid.Parse(got.ID())
				return err == nil
			})

			assert.Equal(t, wantName, got.Name())
			assert.Equal(t, wantDocument, got.Document().String())
			assert.Equal(t, wantEmail, got.Email())
			assert.Equal(t, domain.PayeeDraftStatus, got.Status())
			assert.Equal(t, wantPixKeyType, got.PixKey().Type())
			assert.Equal(t, wantPixKey, got.PixKey().String())
		})
	}
}

func TestCreatePayee_Errors(t *testing.T) {

	tests := []struct {
		name     string
		createFn func() error
		wantErr  error
	}{
		{
			name: "missing name",
			createFn: func() error {
				wantPixKeyType, wantPixKey := fake.PixKey()

				_, err := domain.CreatePayee(
					"",
					fake.CNPJ(),
					wantPixKeyType,
					wantPixKey,
					gofakeit.Email(),
				)
				return err
			},
			wantErr: domain.ErrNameEmptyString,
		},
		{
			name: "invalid document",
			createFn: func() error {
				wantPixKeyType, wantPixKey := fake.PixKey()

				_, err := domain.CreatePayee(
					gofakeit.Name(),
					"invaliddoc",
					wantPixKeyType,
					wantPixKey,
					gofakeit.RandomString([]string{gofakeit.Email(), ""}),
				)

				return err
			},
			wantErr: domain.ErrInvalidDocument,
		},
		{
			name: "invalid email",
			createFn: func() error {
				wantPixKeyType, wantPixKey := fake.PixKey()

				_, err := domain.CreatePayee(
					gofakeit.Name(),
					fake.CPF(),
					wantPixKeyType,
					wantPixKey,
					"invalidemail",
				)

				return err
			},
			wantErr: domain.ErrInvalidEmail,
		},
		{
			name: "invalid pix key type",
			createFn: func() error {
				wantName := gofakeit.Name()
				wantDocument := fake.CPF()

				_, err := domain.CreatePayee(
					wantName,
					wantDocument,
					"invalidtype",
					"none",
					"",
				)
				return err
			},
			wantErr: domain.ErrInvalidPixKeyType,
		},
		{
			name: "invalid cpf pix key",
			createFn: func() error {
				wantName := gofakeit.Name()
				wantDocument := fake.CPF()

				_, err := domain.CreatePayee(
					wantName,
					wantDocument,
					domain.CPFPixKeyType,
					"none",
					"",
				)
				return err
			},
			wantErr: domain.ErrInvalidCPF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ErrorIs(t, tt.createFn(), tt.wantErr)
		})
	}
}

func TestPayee_EditDetails(t *testing.T) {
	for i := 0; i < testPayeeMaxIterations; i++ {
		t.Run("given a payee when status is DRAFT should be possible edit all information", func(t *testing.T) {
			payee := createRandomPayee()

			wantName := gofakeit.Name()
			wantDocument := gofakeit.RandomString([]string{fake.CNPJ(), fake.CPF()})
			wantEmail := gofakeit.RandomString([]string{gofakeit.Email(), ""})
			wantPixKeyType, wantPixKey := fake.PixKey()

			err := payee.EditDetails(
				wantName,
				wantDocument,
				wantPixKeyType,
				wantPixKey,
				wantEmail,
			)

			require.NoError(t, err)

			assert.Equal(t, wantName, payee.Name())
			assert.Equal(t, wantDocument, payee.Document().String())
			assert.Equal(t, wantEmail, payee.Email())
			assert.Equal(t, domain.PayeeDraftStatus, payee.Status())
			assert.Equal(t, wantPixKeyType, payee.PixKey().Type())
			assert.Equal(t, wantPixKey, payee.PixKey().String())
		})
	}

	for i := 0; i < testPayeeMaxIterations; i++ {
		t.Run("given a payee when has status VALID should edit only email", func(t *testing.T) {

			wantID := domain.NewEntityID()
			wantStatus := domain.PayeeValidStatus
			wantName := gofakeit.Name()
			wantDocument := gofakeit.RandomString([]string{fake.CNPJ(), fake.CPF()})
			wantEmail := gofakeit.RandomString([]string{gofakeit.Email(), ""})
			wantPixKeyType, wantPixKey := fake.PixKey()

			payee := domain.RestorePayee(
				wantID.Value(),
				wantName,
				wantDocument,
				wantStatus.Value(),
				wantEmail,
				wantPixKeyType,
				wantPixKey,
				nil,
			)

			newEmail := gofakeit.RandomString([]string{gofakeit.Email(), ""})
			newPixKeyType, newPixKey := fake.PixKey()
			err := payee.EditDetails(
				gofakeit.Name(),
				gofakeit.RandomString([]string{fake.CNPJ(), fake.CPF()}),
				newPixKeyType,
				newPixKey,
				newEmail,
			)

			require.NoError(t, err)

			assert.Equal(t, newEmail, payee.Email())

			assert.Equal(t, wantID.Value(), payee.ID())
			assert.Equal(t, wantName, payee.Name())
			assert.Equal(t, wantDocument, payee.Document().String())
			assert.Equal(t, domain.PayeeValidStatus, payee.Status())
			assert.Equal(t, wantPixKeyType, payee.PixKey().Type())
			assert.Equal(t, wantPixKey, payee.PixKey().String())
		})
	}

}

func createRandomPayee() *domain.PayeeEntity {
	pixKeyType, pixKey := fake.PixKey()

	payee, err := domain.CreatePayee(
		gofakeit.Name(),
		gofakeit.RandomString([]string{fake.CNPJ(), fake.CPF()}),
		pixKeyType,
		pixKey,
		gofakeit.RandomString([]string{gofakeit.Email(), ""}),
	)
	if err != nil {
		panic(err)
	}

	return payee
}
