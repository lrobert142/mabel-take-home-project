package collector

import (
	"fmt"
	"mabel-take-home-project/internal/model"
	"strconv"
)

// TransactionsCollector is a collector that specifically collects model.Transaction data
type TransactionsCollector struct {
	accountsMap  map[string]*model.Account
	transactions []*model.Transaction
}

// Collect collects record data, transforming it into a model.Transaction value, verifying the related model.Account
// items exist and storing it
func (l *TransactionsCollector) Collect(record []string) error {
	from, ok := l.accountsMap[record[0]]
	if !ok {
		return fmt.Errorf("unable to find 'from' account %q", record[0])
	}

	to, ok := l.accountsMap[record[1]]
	if !ok {
		return fmt.Errorf("unable to find 'to' account %q", record[0])
	}

	rawAmount := record[2]
	amount, err := strconv.ParseFloat(rawAmount, 64)
	if err != nil {
		return fmt.Errorf("unable to read amount value %q from CSV record: %w", rawAmount, err)
	}

	transaction, err := model.NewTransaction(from, to, amount)
	if err != nil {
		return fmt.Errorf("unable to create transaction: %w", err)
	}

	l.transactions = append(l.transactions, transaction)

	return nil
}

// GetTransactions returns all stored transactions
func (l *TransactionsCollector) GetTransactions() []*model.Transaction {
	return l.transactions
}

// NewTransactionsCollector creates a new TransactionsCollector with appropriate defaults
func NewTransactionsCollector(accounts []*model.Account) *TransactionsCollector {
	// Convert to a map for easier lookup in future work
	accountsMap := make(map[string]*model.Account)
	for _, account := range accounts {
		accountsMap[account.Id()] = account
	}

	return &TransactionsCollector{
		accountsMap:  accountsMap,
		transactions: make([]*model.Transaction, 0),
	}
}
