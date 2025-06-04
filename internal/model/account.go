package model

import (
	"mabel-take-home-project/internal/errors"
	"math"
	"regexp"
)

// BalanceUpdater represents the functionality needed to update the "balance" of an account
type BalanceUpdater interface {
	UpdateBalance(amount int64) error
}

// accountIdRegex is the regular expression used to verify that an account ID contains only numbers
var accountIdRegex = regexp.MustCompile("^[0-9]+$")

// Account represents a customer's banking information
type Account struct {
	id      string
	balance int64
}

// Id returns the ID of the account
func (a *Account) Id() string {
	return a.id
}

// Balance returns the Balance of the account
func (a *Account) Balance() int64 {
	return a.balance
}

// UpdateBalance updates the Balance of an account.
//
// If an account's balance would be negative as a result of this call, an error will be returned and the balance will be unaffected.
func (a *Account) UpdateBalance(amount int64) error {
	// Short-circuit if there is nothing to do
	if amount == 0 {
		return nil
	}

	// Working with negative numbers is strange when checking for min/max boundaries, so make everything positive and
	// do our checks that way
	posBalance := a.balance
	if posBalance < 0 {
		posBalance = -posBalance
	}
	posAmount := amount
	if posAmount < 0 {
		posAmount = -posAmount
	}
	posTotal := uint64(posBalance) + uint64(posAmount)

	// Overflow check, if adding an amount make sure it doesn't go over our upper data limit
	if amount > 0 && posTotal > uint64(math.MaxInt64) {
		return errors.AccountBalanceUpdateOverflow
	} else {
		// Our minimum is a slightly different value to our maximum, so invert it to make sure we are using the correct value
		posMin := math.MinInt64
		posMin = -posMin

		// Underflow check, if subtracting an amount make sure it doesn't go under our lower data limit
		if posTotal > uint64(posMin) {
			return errors.AccountBalanceUpdateUnderflow
		}
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
func NewAccount(id string, balance int64) (*Account, error) {
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
