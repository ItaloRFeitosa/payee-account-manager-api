package domain_test

import (
	"testing"

	"github.com/italorfeitosa/payee-account-manager-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    string
		wantErr error
	}{
		{
			name: "given a valid person name should return a valid Name",
			arg:  "Italo Feitosa",
			want: "Italo Feitosa",
		},
		{
			name: "given a valid company name should return a valid Name",
			arg:  "The Fake Company LTDA",
			want: "The Fake Company LTDA",
		},
		{
			name: "given a valid company name when have extra space should return a valid Name without the spaces",
			arg:  "The     Fake   Company  LTDA",
			want: "The Fake Company LTDA",
		},
		{
			name:    "given a name when has only word should return error",
			arg:     "Italo",
			wantErr: domain.ErrNameLessThenTwoWords,
		},
		{
			name:    "given a name when first name is too short should return error",
			arg:     "A Silva",
			wantErr: domain.ErrShortFirstName,
		},
		{
			name:    "given an empty name should return error",
			arg:     "",
			wantErr: domain.ErrNameEmptyString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, err := domain.NewName(tt.arg)
			if tt.wantErr != nil {
				assert.Equal(t, domain.EmptyName, name)
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.want, name.Value())
				assert.NoError(t, err)
			}
		})
	}
}
