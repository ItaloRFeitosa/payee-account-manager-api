package domain_test

import (
	"testing"

	"github.com/italorfeitosa/payee-account-manager-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    string
		wantErr error
	}{
		{
			name: "given a valid email should return a valid Email object",
			arg:  "italo@feitosa.com",
			want: "italo@feitosa.com",
		},
		{
			name: "given a valid email should return a valid Email object",
			arg:  "italo_feitosa@domain.com",
			want: "italo_feitosa@domain.com",
		},
		{
			name: "given a valid email when has space around should return a valid Email object",
			arg:  "    italo@feitosa.com       ",
			want: "italo@feitosa.com",
		},
		{
			name:    "given a email when not match with regex should return error",
			arg:     "testnoat",
			wantErr: domain.ErrInvalidEmail,
		},
		{
			name:    "given a email when not match with regex should return error",
			arg:     "$$$$$@email.com",
			wantErr: domain.ErrInvalidEmail,
		},
		{
			name:    "given a email when not match with regex should return error",
			arg:     "$$$$$@/////*.com",
			wantErr: domain.ErrInvalidEmail,
		},
		{
			name:    "given an empty email should return error",
			arg:     "",
			wantErr: domain.ErrInvalidEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := domain.NewEmail(tt.arg)
			if tt.wantErr != nil {
				assert.Equal(t, domain.EmptyEmail, email)
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.want, email.Value())
				assert.NoError(t, err)
			}
		})
	}
}
