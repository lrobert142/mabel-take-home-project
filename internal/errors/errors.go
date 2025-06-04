package errors

import (
	"errors"
)

var (
	AccountIdLengthInvalid        = errors.New("account ID must be exactly 16 characters long")
	AccountIdFormatInvalid        = errors.New("account ID must only contain digits")
	AccountBalanceUpdateNegative  = errors.New("account balance would be made negative by this update")
	AccountBalanceUpdateUnderflow = errors.New("account balance update not allowed: Modification would underflow minimum allowed value")
	AccountBalanceUpdateOverflow  = errors.New("account balance update not allowed: Modification would overflow maximum allowed value")
)
