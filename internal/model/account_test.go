package model_test

import (
	"mabel-take-home-project/internal/errors"
	"mabel-take-home-project/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	validId := "1234567890123456"
	validBalance := float64(100)

	tooLongId := "12345678901234567890"
	tooShortId := "123"
	idWithNonDigits := "a234567~9012345 "

	tests := map[string]struct {
		id              string
		balance         float64
		expectedAccount *model.Account
		expectedError   error
	}{
		"Return account when details are valid": {
			id:              validId,
			balance:         validBalance,
			expectedAccount: model.NewTestAccount(validId, validBalance),
			expectedError:   nil,
		},
		"Return error when ID is too short": {
			id:              tooLongId,
			balance:         validBalance,
			expectedAccount: nil,
			expectedError:   errors.AccountIdLengthInvalid,
		},
		"Return error when ID is too long": {
			id:              tooShortId,
			balance:         validBalance,
			expectedAccount: nil,
			expectedError:   errors.AccountIdLengthInvalid,
		},
		"Return error when ID contains something other than numbers": {
			id:              idWithNonDigits,
			balance:         validBalance,
			expectedAccount: nil,
			expectedError:   errors.AccountIdFormatInvalid,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualAccount, actualError := model.NewAccount(test.id, test.balance)
			assert.Equal(t, test.expectedAccount, actualAccount, "Accounts do not match!")
			assert.Equal(t, test.expectedError, actualError, "Accounts do not match!")
		})
	}
}

func TestUpdateBalance(t *testing.T) {
	id := "1234567890123456"

	tests := map[string]struct {
		account         *model.Account
		amount          float64
		expectedBalance float64
		expectedError   error
	}{
		"Return no error and update the account balance when it finishes positive": {
			account:         model.NewTestAccount(id, 0),
			amount:          100,
			expectedBalance: 100,
			expectedError:   nil,
		},
		"Return no error and update the account balance when it finishes at exactly zero": {
			account:         model.NewTestAccount(id, 100),
			amount:          -100,
			expectedBalance: 0,
			expectedError:   nil,
		},
		"Return no error when the account finishes negative, but amount is positive": {
			account:         model.NewTestAccount(id, -200),
			amount:          100,
			expectedBalance: -100,
			expectedError:   nil,
		},
		"Return error and do not update balance when the account would finish negative (and the amount is negative)": {
			account:         model.NewTestAccount(id, 100),
			amount:          -200,
			expectedBalance: 100,
			expectedError:   errors.AccountBalanceUpdateNegative,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualErr := test.account.UpdateBalance(test.amount)
			assert.Equal(t, test.expectedError, actualErr)
			assert.Equal(t, test.expectedBalance, test.account.Balance())
		})
	}
}
