# mabel-take-home-project

Take home interview project for Mabel

## Quickstart

Make sure you have [Go](https://go.dev/) installed, then run the project via
`go run ./cmd/mabel-take-home-project <account_balance_csv> <transactions_csv>`

- `account_balance_csv` is a CSV file containing the starting information for an account (
  see  [mable_account_balances.csv](./mable_account_balances.csv) as an example)
- `transactions_csv` is a CSV file containing a list of transactions between existing accounts (
  see [mable_transactions.csv](./mable_transactions.csv) as an example)

### Tests

Make sure you have [Go](https://go.dev/) installed, then:

- Unit tests can be run with `go test ./...`
- Integration tests can be run with `go test -tags=integration ./...`

## Summary

You are a developer for a company that runs a very simple banking service. Each
day companies provide you with a CSV file with transfers they want to make
between accounts for customers they are doing business with. Accounts are
identified by a 16 digit number and money cannot be transferred from them if it
will put the account balance below $0. Your job is to implement a simple system
that can load account balances for a single company and then accept a day's
transfers in a CSV file. An example customer balance file is provided as well
as an example days transfers.

eg [mable_account_balances.csv](./mable_account_balances.csv)

| Starting state of accounts for Account | customers of Alpha Sales: Balance |
|---------------------------------------:|----------------------------------:|
|                       1111234522226789 |                           5000.00 |
|                       1111234522221234 |                          10000.00 |
|                       2222123433331212 |                            550.00 |
|                       1212343433335665 |                           1200.00 |
|                       3212343433335755 |                          50000.00 |

Single day transactions for Alpha sales:

eg [mable_transactions.csv](./mable_transactions.csv)

|             From |               To |  Amount |
|-----------------:|-----------------:|--------:|
| 1111234522226789 | 1212343433335665 |  500.00 |
| 3212343433335755 | 2222123433331212 | 1000.00 |
| 3212343433335755 | 1111234522226789 |  320.50 |
| 1111234522221234 | 1212343433335665 |   25.60 |

## Assumptions

- Accounts _may_ start with a negative balance
- "...money cannot be transferred _from_ them if it will put the account balance below $0..." however money can be
  transferred _to_ them, even if the final balance would be below zero
- If a transaction contains a `From` or `To` value that does not correlate to an account, an error is raised and
  processing stops
- A transaction `amount` _cannot_ be negative. Any `amount` with a negative value should return an error stop processing
- If a payment fails partway through both accounts are 'reverted' to their previous state before the transaction was
  run, even if this would cause an account to have a negative balance
- A transaction may specify `from` and `to` as the same account. It will be processed like normal, assuming all other
  restrictions and assumptions are respected
