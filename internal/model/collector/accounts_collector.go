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

//TODO: TESTME!
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

//TODO: TESTME
func (l *AccountsCollector) GetAccounts() []*model.Account {
	return l.accounts
}

//TODO: TESTME
func NewAccountsCollector() *AccountsCollector {
	return &AccountsCollector{
		accounts: make([]*model.Account, 0),
	}
}
