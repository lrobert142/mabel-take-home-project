package model

import (
	"fmt"
	"mabel-take-home-project/internal/errors"
)

// Transaction represents the transfer of funds from one Account to another
type Transaction struct {
	from   BalanceUpdater
	to     BalanceUpdater
	amount float64
}

// Transact updates the balances of both `from` and `to` accounts.
//
// If an error occurs partway through, both accounts will be rolled back to their last valid state.
func (t *Transaction) Transact() error {
	if err := t.from.UpdateBalance(-t.amount); err != nil {
		return fmt.Errorf("failed to update balance for 'from': %w", err)
	}
	if err := t.to.UpdateBalance(t.amount); err != nil {
		// Revert `from` to ensure no funds go missing
		if err := t.from.UpdateBalance(t.amount); err != nil {
			return fmt.Errorf("failed to revert balance for 'from': %w", err)
		}
		return fmt.Errorf("failed to update balance for 'to': %w", err)
	}

	return nil
}

// NewTransaction creates a new Transaction with the specified details.
//
// If data is invalid for initialisation, an error will be returned.
func NewTransaction(from BalanceUpdater, to BalanceUpdater, amount float64) (*Transaction, error) {
	if amount < 0 {
		return nil, errors.TransactionAmountInvalid
	}

	return &Transaction{
		from:   from,
		to:     to,
		amount: amount,
	}, nil
}
