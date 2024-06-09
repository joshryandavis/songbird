# Songbird

```
Usage: songbird <command> [flags]

a cli interface for starling bank

Flags:
  -h, --help             Show context-sensitive help.
  -v, --verbose          Output verbose logs.
  -a, --account="all"    Account Name or "all".

Commands:
  balance (bal)        List account balances.
  recurring (rp)       List recurring payments.
  direct-debit (dd)    List direct debits.
  transactions (tx)    List all transactions.
  accounts (acc)       List all accounts.
  walk (w)             Walk the transactions tree.
```

A little Go program for accessing Starling Bank data in the terminal, planning budgets and managing upcoming payments.

Its USP is that it allows for custom metadata to be addedd to tranactions in the form of JSON embeded in the transaction note. This means no data is kept locally and the state is all stored on Starling. 

Note: the code that calculates future payment dates is hard-coded to use English bank holidays. 

**Work in progress**.
