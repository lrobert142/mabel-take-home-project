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

// TODO TESTME
func (l *TransactionsCollector) Collect(record []string) error {
	from, ok := l.accountsMap[record[0]]
	if !ok {
		return fmt.Errorf("unable to find 'from' account %q", record[0])
	}

	to, ok := l.accountsMap[record[0]]
	if !ok {
		return fmt.Errorf("unable to find 'to' account %q", record[0])
	}

	rawAmount := record[1]
	amount, err := strconv.ParseFloat(rawAmount, 64)
	if err != nil {
		return fmt.Errorf("unable to read amount value %q from CSV record: %w", rawAmount, err)
	}

	transaction, err := model.NewTransaction(from, to, amount)
	if err != nil {
		return fmt.Errorf("unable to create tranasction: %w", err)
	}

	l.transactions = append(l.transactions, transaction)

	return nil
}

// TODO: TESTME
func (l *TransactionsCollector) GetTransactions() []*model.Transaction {
	return l.transactions
}

// TODO: TESTME?
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
