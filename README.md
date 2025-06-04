# mabel-take-home-project

Take home interview project for Mabel

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
- Transactions that would result in account balances being over the maximum allowed data value or under the minimum
  allowed data value will result in an error and no balance change
- A transaction _may_ provide `From`/`To` IDs that do not exist in the system (account not found), in this case that
  transaction should return an error and the remaining transactions continue to be processed
- A transaction `amount` _cannot_ be negative. Any `amount` with a negative value should return an error and the
  remaining transactions continue to be processed
- If a payment fails partway through both accounts are 'reverted' to their previous state before the transaction was
  run, even if this would cause an account to have a negative balance
