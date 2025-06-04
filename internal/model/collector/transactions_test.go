package collector_test

import (
	"github.com/stretchr/testify/assert"
	"mabel-take-home-project/internal/model"
	"mabel-take-home-project/internal/model/collector"
	"strconv"
	"testing"
)

func TestTransactionsCollector_Collect(t *testing.T) {
	validAccountId1 := "1234567890123456"
	validAccountId2 := "6543210987654321"
	validAccounts := map[string]*model.Account{
		validAccountId1: {},
		validAccountId2: {},
	}
	validAmount := "123.45"
	validAmountAsFloat, _ := strconv.ParseFloat(validAmount, 64)

	tests := map[string]struct {
		record               []string
		errorContains        string
		expectedTransactions []*model.Transaction
	}{
		"Collect data successfully when the record is valid, and all accounts exist": {
			record:        []string{validAccountId1, validAccountId2, validAmount},
			errorContains: "",
			expectedTransactions: func() []*model.Transaction {
				transaction, _ := model.NewTransaction(validAccounts[validAccountId1], validAccounts[validAccountId2], validAmountAsFloat)
				return []*model.Transaction{transaction}
			}(),
		},
		"Return an error when the from account does not exist": {
			record:               []string{"not an ID", validAccountId2, validAmount},
			errorContains:        "unable to find 'from' account",
			expectedTransactions: []*model.Transaction{},
		},
		"Return an error when the to account does not exist": {
			record:               []string{validAccountId1, "not an ID", validAmount},
			errorContains:        "unable to find 'to' account",
			expectedTransactions: []*model.Transaction{},
		},
		"Return an error when the amount cannot be parsed": {
			record:               []string{validAccountId1, validAccountId2, "not a number"},
			errorContains:        "unable to read amount value",
			expectedTransactions: []*model.Transaction{},
		},
		"Return an error when there is an issue creating a transaction": {
			record:               []string{validAccountId1, validAccountId2, "-100"},
			errorContains:        "unable to create transaction",
			expectedTransactions: []*model.Transaction{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ac := collector.NewTestTransactionsCollector(validAccounts, []*model.Transaction{})
			if test.errorContains == "" {
				assert.NoError(t, ac.Collect(test.record))
			} else {
				assert.ErrorContains(t, ac.Collect(test.record), test.errorContains)
			}
			assert.Equal(t, test.expectedTransactions, ac.GetTransactions())
		})
	}
}

func TestTransactionsCollector_GetTransactions(t *testing.T) {
	transaction, _ := model.NewTransaction(&model.Account{}, &model.Account{}, 123.45)
	transactions := []*model.Transaction{transaction}

	tests := map[string]struct {
		ac                   *collector.TransactionsCollector
		expectedTransactions []*model.Transaction
	}{
		"Return all stored transactions": {
			ac:                   collector.NewTestTransactionsCollector(nil, transactions),
			expectedTransactions: transactions,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedTransactions, test.ac.GetTransactions())
		})
	}
}

func TestTransactionsCollector_NewTransactionsCollector(t *testing.T) {
	assert.NotNil(t, collector.NewTransactionsCollector([]*model.Account{}))
}
