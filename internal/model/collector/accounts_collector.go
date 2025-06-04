package collector

import (
	"fmt"
	"mabel-take-home-project/internal/model"
	"strconv"
)

// AccountsCollector is a collector that specifically collects model.Account data
type AccountsCollector struct {
	accounts []*model.Account
}

// Collect collects record data, transforming it into a model.Account value and storing it
func (l *AccountsCollector) Collect(record []string) error {
	rawBalance := record[1]
	balance, err := strconv.ParseFloat(rawBalance, 64)
	if err != nil {
		return fmt.Errorf("unable to read balance value %q from CSV record: %w", rawBalance, err)
	}

	account, err := model.NewAccount(record[0], balance)
	if err != nil {
		return fmt.Errorf("unable to create account: %w", err)
	}

	l.accounts = append(l.accounts, account)
	return nil
}

// GetAccounts returns all stored accounts
func (l *AccountsCollector) GetAccounts() []*model.Account {
	return l.accounts
}

// NewAccountsCollector creates a new AccountsCollector with appropriate defaults
func NewAccountsCollector() *AccountsCollector {
	return &AccountsCollector{
		accounts: make([]*model.Account, 0),
	}
}
