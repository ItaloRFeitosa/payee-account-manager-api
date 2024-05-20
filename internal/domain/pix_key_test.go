package domain_test

import (
	"testing"

	"github.com/italorfeitosa/payee-account-manager-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewPixKey(t *testing.T) {
	type args struct {
		typ   string
		value string
	}
	tests := []struct {
		name       string
		args       args
		wantType   string
		wantValue  string
		wantString string
		wantErr    error
	}{
		// CPF pix key
		{
			name: "given a valid cpf key should return a cpf pix key",
			args: args{
				typ:   domain.CPFPixKeyType,
				value: "606.883.970-26",
			},
			wantType:   domain.CPFPixKeyType,
			wantValue:  "60688397026",
			wantString: "606.883.970-26",
		},
		{
			name: "given an invalid cpf key should return error",
			args: args{
				typ:   domain.CPFPixKeyType,
				value: "60698397026",
			},
			wantErr: domain.ErrInvalidCPF,
		},
		// CNPJ pix key
		{
			name: "given a valid CNPJ key should return a CNPJ pix key",
			args: args{
				typ:   domain.CNPJPixKeyType,
				value: "65.678.974/0001-74",
			},
			wantType:   domain.CNPJPixKeyType,
			wantValue:  "65678974000174",
			wantString: "65.678.974/0001-74",
		},
		{
			name: "given an invalid CNPJ key should return error",
			args: args{
				typ:   domain.CNPJPixKeyType,
				value: "65678974000175",
			},
			wantErr: domain.ErrInvalidCNPJ,
		},
		// Email pix key
		{
			name: "given a valid Email key should return a Email pix key",
			args: args{
				typ:   domain.EmailPixKeyType,
				value: "italo@feitosa.com",
			},
			wantType:   domain.EmailPixKeyType,
			wantValue:  "italo@feitosa.com",
			wantString: "italo@feitosa.com",
		},
		{
			name: "given an invalid Email key should return error",
			args: args{
				typ:   domain.EmailPixKeyType,
				value: "italofeitosacom",
			},
			wantErr: domain.ErrInvalidEmail,
		},
		// Telefone pix key
		{
			name: "given a valid Telefone key should return a Telefone pix key",
			args: args{
				typ:   domain.TelefonePixKeyType,
				value: "99987654321",
			},
			wantType:   domain.TelefonePixKeyType,
			wantValue:  "5599987654321",
			wantString: "+5599987654321",
		},
		{
			name: "given a valid Telefone key should return a Telefone pix key",
			args: args{
				typ:   domain.TelefonePixKeyType,
				value: "+5599987654321",
			},
			wantType:   domain.TelefonePixKeyType,
			wantValue:  "5599987654321",
			wantString: "+5599987654321",
		},
		{
			name: "given an invalid Telefone key should return error",
			args: args{
				typ:   domain.TelefonePixKeyType,
				value: "465454",
			},
			wantErr: domain.ErrInvalidTelefone,
		},
		{
			name: "given an invalid Telefone key should return error",
			args: args{
				typ:   domain.TelefonePixKeyType,
				value: "55999876543210",
			},
			wantErr: domain.ErrInvalidTelefone,
		},
		// ChaveAleatoria pix key
		{
			name: "given a valid ChaveAleatoria key should return a ChaveAleatoria pix key",
			args: args{
				typ:   domain.ChaveAleatoriaPixKeyType,
				value: "9fcacf81-0f7f-47a2-881d-da4f41a39afd",
			},
			wantType:   domain.ChaveAleatoriaPixKeyType,
			wantValue:  "9fcacf81-0f7f-47a2-881d-da4f41a39afd",
			wantString: "9fcacf81-0f7f-47a2-881d-da4f41a39afd",
		},
		{
			name: "given a valid ChaveAleatoria key should return a ChaveAleatoria pix key",
			args: args{
				typ:   domain.ChaveAleatoriaPixKeyType,
				value: "9FCacF81-0F7f-47A2-881D-dA4F41A39AFD",
			},
			wantType:   domain.ChaveAleatoriaPixKeyType,
			wantValue:  "9fcacf81-0f7f-47a2-881d-da4f41a39afd",
			wantString: "9fcacf81-0f7f-47a2-881d-da4f41a39afd",
		},
		{
			name: "given an invalid ChaveAleatoria key should return error",
			args: args{
				typ:   domain.ChaveAleatoriaPixKeyType,
				value: "9fcacf810f7f47a2881dda4f41a39afd",
			},
			wantErr: domain.ErrInvalidChaveAleatoria,
		},
		{
			name: "given an invalid ChaveAleatoria key should return error",
			args: args{
				typ:   domain.ChaveAleatoriaPixKeyType,
				value: "9fcacf810-f7f47a288-1dda4f41a-39afd",
			},
			wantErr: domain.ErrInvalidChaveAleatoria,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewPixKey(tt.args.typ, tt.args.value)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.wantType, got.Type())
				assert.Equal(t, tt.wantValue, got.Value())
				assert.Equal(t, tt.wantString, got.String())
			}
		})
	}
}
