package collector_test

import (
	"github.com/stretchr/testify/assert"
	"mabel-take-home-project/internal/model"
	"mabel-take-home-project/internal/model/collector"
	"strconv"
	"testing"
)

func TestAccountsCollector_Collect(t *testing.T) {
	validId := "1234567890123456"
	validBalance := "123.45"
	validBalanceAsFloat, _ := strconv.ParseFloat(validBalance, 64)

	tests := map[string]struct {
		record           []string
		errorContains    string
		expectedAccounts []*model.Account
	}{
		"Collect the data successfully when the record is valid": {
			record:        []string{validId, validBalance},
			errorContains: "",
			expectedAccounts: func() []*model.Account {
				acc, _ := model.NewAccount(validId, validBalanceAsFloat)
				return []*model.Account{acc}
			}(),
		},
		"Return an error when the balance cannot be parsed": {
			record:           []string{validId, "not a number"},
			errorContains:    "unable to read balance value",
			expectedAccounts: []*model.Account{},
		},
		"Return an error when there is an issue creating an account": {
			record:           []string{"bad ID", validBalance},
			errorContains:    "unable to create account",
			expectedAccounts: []*model.Account{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ac := collector.NewTestAccountsCollector([]*model.Account{})
			if test.errorContains == "" {
				assert.NoError(t, ac.Collect(test.record))
			} else {
				assert.ErrorContains(t, ac.Collect(test.record), test.errorContains)
			}
			assert.Equal(t, test.expectedAccounts, ac.GetAccounts())
		})
	}
}

func TestAccountsCollector_GetAccounts(t *testing.T) {
	account, _ := model.NewAccount("12345678901234556", 123.45)
	accounts := []*model.Account{account}

	tests := map[string]struct {
		ac               *collector.AccountsCollector
		expectedAccounts []*model.Account
	}{
		"Return all stored accounts": {
			ac:               collector.NewTestAccountsCollector(accounts),
			expectedAccounts: accounts,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedAccounts, test.ac.GetAccounts())
		})
	}
}

func TestAccountsCollector_NewAccountsCollector(t *testing.T) {
	assert.NotNil(t, collector.NewAccountsCollector())
}
