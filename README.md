# songbird

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

A little Go program for accessing Starling Bank data in the terminal.

The USP is that you can walk through transactions and add whatever metadata you want in the transaction note. This way I can build on top of the Starling API and use the notes as my database. 

Work in progress.
