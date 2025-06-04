package model

// NewTestAccount creates a new Account to be used in testing only, allowing special setup to the inner values.
//
// It is exported in this `..._test.go` file to enable ease of testing but will not be included in the final build due
// to how Golang handles `..._test.go` files, so we do not need to worry about accidentally using this in production code.
func NewTestAccount(id string, balance int64) *Account {
	return &Account{
		id:      id,
		balance: balance,
	}
}
