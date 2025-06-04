package collector

import "mabel-take-home-project/internal/model"

// NewTestAccountsCollector creates a new AccountsCollector to be used in testing only, allowing special setup of the inner values.
//
// It is exported in this `..._test.go` file to enable ease of testing but will not be included in the final build due
// to how Golang handles `..._test.go` files, so we do not need to worry about accidentally using this in production code.
func NewTestAccountsCollector(accounts []*model.Account) *AccountsCollector {
	return &AccountsCollector{
		accounts: accounts,
	}
}

// NewTestTransactionsCollector creates a new TransactionsCollector to be used in testing only, allowing special setup of the inner values.
//
// It is exported in this `..._test.go` file to enable ease of testing but will not be included in the final build due
// to how Golang handles `..._test.go` files, so we do not need to worry about accidentally using this in production code.
func NewTestTransactionsCollector(accountsMap map[string]*model.Account, transactions []*model.Transaction) *TransactionsCollector {
	return &TransactionsCollector{
		accountsMap:  accountsMap,
		transactions: transactions,
	}
}
