package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount_Validate(t *testing.T) {
	tests := []struct {
		name    string
		account Account
		wantErr bool
		err     error
	}{
		{
			name: "valid account",
			account: Account{
				Email:        "test@example.com",
				PasswordHash: "hashed_pass",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			account: Account{
				Email:        "invalid",
				PasswordHash: "hashed_pass",
			},
			wantErr: true,
			err:     ErrInvalidEmail,
		},
		{
			name: "missing email",
			account: Account{
				Email:        "",
				PasswordHash: "hashed_pass",
			},
			wantErr: true,
			err:     ErrInvalidEmail,
		},
		{
			name: "missing password",
			account: Account{
				Email:        "test@example.com",
				PasswordHash: "",
			},
			wantErr: true,
			err:     ErrPasswordRequired,
		},
		{
			name: "system account missing password",
			account: Account{
				Email:           "system@hyperx.io",
				PasswordHash:    "",
				IsSystemAccount: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAccount_CanActAsCreator(t *testing.T) {
	acc := Account{IsSystemAccount: false}
	assert.True(t, acc.CanActAsCreator())

	sysAcc := Account{IsSystemAccount: true}
	assert.False(t, sysAcc.CanActAsCreator())
}
