package errors

import (
	"errors"
)

var (
	// AccountIdLengthInvalid is an error used when an account ID is either too long or too short
	AccountIdLengthInvalid = errors.New("account ID must be exactly 16 characters long")
	// AccountIdFormatInvalid is an error used when an account ID is incorrectly formatted
	AccountIdFormatInvalid = errors.New("account ID must only contain digits")
	// AccountBalanceUpdateNegative is an error used when a requested balance update would result in a negative balance
	AccountBalanceUpdateNegative = errors.New("account balance would be made negative by this update")

	// TransactionAmountInvalid is an error when a transaction has an invalid value for `amount`
	TransactionAmountInvalid = errors.New("transaction amount must be a positive number")
)
