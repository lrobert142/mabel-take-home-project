package model

import (
	"fmt"
)

// Transaction represents the transfer of funds from one Account to another
type Transaction struct {
	from   BalanceUpdater
	to     BalanceUpdater
	amount int64
}

func (t *Transaction) Process() error {
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

func NewTransaction(from BalanceUpdater, to BalanceUpdater, amount int64) *Transaction {
	return &Transaction{
		from:   from,
		to:     to,
		amount: amount,
	}
}
