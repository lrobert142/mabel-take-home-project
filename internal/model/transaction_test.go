package model_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	customErrors "mabel-take-home-project/internal/errors"
	"mabel-take-home-project/internal/model"
	"testing"
)

var mockUpdateBalanceError = errors.New("manually thrown UpdateBalance error")

// MockBalanceUpdater is a mock implementation of BalanceUpdater used exclusively for testing purposes
type MockBalanceUpdater struct {
	mustFail                           bool
	expectedNumberOfUpdateBalanceCalls int
	actualNumberOfUpdateBalanceCalls   int
}

func (m *MockBalanceUpdater) UpdateBalance(_ float64) error {
	m.actualNumberOfUpdateBalanceCalls += 1

	if m.mustFail {
		return mockUpdateBalanceError
	}
	return nil
}

func TestNewTransaction(t *testing.T) {
	tests := map[string]struct {
		amount        float64
		expectedError error
	}{
		"Return no error when the amount is valid": {
			amount:        100,
			expectedError: nil,
		},
		"Return an error when the amount is invalid": {
			amount:        -100,
			expectedError: customErrors.TransactionAmountInvalid,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, actualError := model.NewTransaction(&MockBalanceUpdater{}, &MockBalanceUpdater{}, test.amount)
			assert.Equal(t, test.expectedError, actualError)
		})
	}
}

func TestTransact(t *testing.T) {
	type transactionAndAccounts struct {
		from        *MockBalanceUpdater
		to          *MockBalanceUpdater
		transaction *model.Transaction
	}

	tests := map[string]struct {
		transactionAndAccounts *transactionAndAccounts
		expectedError          error
	}{
		"Return no error when no errors are raised": {
			transactionAndAccounts: func() *transactionAndAccounts {
				from := &MockBalanceUpdater{
					mustFail:                           false,
					expectedNumberOfUpdateBalanceCalls: 1,
				}
				to := &MockBalanceUpdater{
					mustFail:                           false,
					expectedNumberOfUpdateBalanceCalls: 1,
				}
				return &transactionAndAccounts{
					from:        from,
					to:          to,
					transaction: model.NewTestTransaction(from, to, 100),
				}
			}(),
			expectedError: nil,
		},
		"Return the underlying error when processing 'from' fails": {
			transactionAndAccounts: func() *transactionAndAccounts {
				from := &MockBalanceUpdater{
					mustFail:                           true,
					expectedNumberOfUpdateBalanceCalls: 1,
				}
				to := &MockBalanceUpdater{
					mustFail: false,
					// Never call `to` when `from` fails
					expectedNumberOfUpdateBalanceCalls: 0,
				}
				return &transactionAndAccounts{
					from:        from,
					to:          to,
					transaction: model.NewTestTransaction(from, to, 100),
				}
			}(),
			expectedError: fmt.Errorf("failed to update balance for 'from': %w", mockUpdateBalanceError),
		},
		"Return the underlying error when processing 'to' fails": {
			transactionAndAccounts: func() *transactionAndAccounts {
				from := &MockBalanceUpdater{
					mustFail: false,
					// Called once initially, then again for the 'revert' when `to` fails
					expectedNumberOfUpdateBalanceCalls: 2,
				}
				to := &MockBalanceUpdater{
					mustFail:                           true,
					expectedNumberOfUpdateBalanceCalls: 1,
				}
				return &transactionAndAccounts{
					from:        from,
					to:          to,
					transaction: model.NewTestTransaction(from, to, 100),
				}
			}(),
			expectedError: fmt.Errorf("failed to update balance for 'to': %w", mockUpdateBalanceError),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedError, test.transactionAndAccounts.transaction.Transact())
			assert.Equal(t, test.transactionAndAccounts.to.expectedNumberOfUpdateBalanceCalls, test.transactionAndAccounts.to.actualNumberOfUpdateBalanceCalls)
			assert.Equal(t, test.transactionAndAccounts.from.expectedNumberOfUpdateBalanceCalls, test.transactionAndAccounts.from.actualNumberOfUpdateBalanceCalls)
		})
	}
}
