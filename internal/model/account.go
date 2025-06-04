package model

import (
	"mabel-take-home-project/internal/errors"
	"regexp"
)

// BalanceUpdater represents the functionality needed to update the "balance" of an account
type BalanceUpdater interface {
	UpdateBalance(amount float64) error
}

// accountIdRegex is the regular expression used to verify that an account ID contains only numbers
var accountIdRegex = regexp.MustCompile("^[0-9]+$")

// Account represents a customer's banking information
type Account struct {
	id      string
	balance float64
}

// Id returns the ID of the account
func (a *Account) Id() string {
	return a.id
}

// Balance returns the Balance of the account
func (a *Account) Balance() float64 {
	return a.balance
}

// UpdateBalance updates the Balance of an account.
//
// If an account's balance would be negative as a result of this call, an error will be returned and the balance will be unaffected.
func (a *Account) UpdateBalance(amount float64) error {
	// Short-circuit if there is nothing to do
	if amount == 0 {
		return nil
	}

	// Money cannot be transferred *from* an account if it would end up negative, but it can be transferred *to* an
	// account and leave it negative. So only return an error if we're taking money out (negative amount)
	if amount < 0 && a.balance+amount < 0 {
		return errors.AccountBalanceUpdateNegative
	}

	a.balance += amount
	return nil
}

// NewAccount creates a new Account with the specified details.
//
// If data is invalid for initialisation, an error will be returned.
func NewAccount(id string, balance float64) (*Account, error) {
	if len(id) != 16 {
		return nil, errors.AccountIdLengthInvalid
	}

	if !accountIdRegex.MatchString(id) {
		return nil, errors.AccountIdFormatInvalid
	}

	return &Account{
		id:      id,
		balance: balance,
	}, nil
}
