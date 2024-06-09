# songbird

Install: `go install github.com/joshryandavis/songbird@latest`

Usage: `songbird --help`

---

A little Go program for accessing Starling Bank data in the terminal, planning budgets and managing upcoming payments.

Its USP is that it allows for custom metadata to be addedd to tranactions in the form of JSON embeded in the transaction note. This means no data is kept locally and the state is all stored on Starling. 

Note: the code that calculates future payment dates is hard-coded to use English bank holidays. 

**Work in progress**. Most of what I want to do is not implemented yet.
