package domain_test

import (
	"testing"

	"github.com/italorfeitosa/payee-account-manager-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	tests := []struct {
		name       string
		arg        string
		wantValue  string
		wantString string
		wantErr    error
	}{
		// Valid CPFs
		{
			name:       "given a valid cpf should return a document instance",
			arg:        "77386735081",
			wantValue:  "77386735081",
			wantString: "773.867.350-81",
		},
		{
			name:       "given a valid cpf should return a document instance",
			arg:        "33860422014",
			wantValue:  "33860422014",
			wantString: "338.604.220-14",
		},
		{
			name:       "given a valid cpf should return a document instance",
			arg:        "128731.950-53",
			wantValue:  "12873195053",
			wantString: "128.731.950-53",
		},
		{
			name:       "given a valid cpf should return a document instance",
			arg:        "616.38824070",
			wantValue:  "61638824070",
			wantString: "616.388.240-70",
		},
		// Invalid CPFs
		{
			name:    "given a invalid cpf should return error",
			arg:     "773867350-82",
			wantErr: domain.ErrInvalidDocument,
		},
		{
			name:    "given a invalid cpf should return error",
			arg:     "33860*422014",
			wantErr: domain.ErrInvalidDocument,
		},
		{
			name:    "given a invalid cpf should return error",
			arg:     "12873295053",
			wantErr: domain.ErrInvalidDocument,
		},
		// Valid CNPJ
		{
			name:       "given a valid cnpj should return a document instance",
			arg:        "19039318000104",
			wantValue:  "19039318000104",
			wantString: "19.039.318/0001-04",
		},
		{
			name:       "given a valid cnpj should return a document instance",
			arg:        "35.952.585/0001-24",
			wantValue:  "35952585000124",
			wantString: "35.952.585/0001-24",
		},
		{
			name:       "given a valid cnpj should return a document instance",
			arg:        "63.2051020001-63",
			wantValue:  "63205102000163",
			wantString: "63.205.102/0001-63",
		},
		{
			name:       "given a valid cnpj should return a document instance",
			arg:        "01042140/000195",
			wantValue:  "01042140000195",
			wantString: "01.042.140/0001-95",
		},
		// Invalid CNPJ
		{
			name:    "given an invalid cnpj should return a error",
			arg:     "19039*318000*104",
			wantErr: domain.ErrInvalidDocument,
		},
		{
			name:    "given an invalid cnpj should return a error",
			arg:     "35.952.585/0002-25",
			wantErr: domain.ErrInvalidDocument,
		},
		{
			name:    "given an invalid cnpj should return a error",
			arg:     "63345102000163",
			wantErr: domain.ErrInvalidDocument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := domain.NewDocument(tt.arg)
			if tt.wantErr != nil {
				assert.Nil(t, doc)
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.wantValue, doc.Value())
				assert.Equal(t, tt.wantString, doc.String())
			}
		})
	}
}
