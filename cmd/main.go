package cmd

import (
	_ "embed"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/cli"
	"github.com/joshryandavis/songbird/internal/client"
	"github.com/joshryandavis/songbird/internal/config"
)

func Main() {
	_, command := cli.Enter()
	tokens, err := config.LoadTokens()
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	c := client.New(tokens, cfg)
	// loop through all starling clients
	for _, ac := range c.Clients {
		acc := ac.Account
		switch command {
		// walk transactions
		case "walk":
			m := cli.Model.Walk
			var start time.Time
			var end time.Time
			if m.DateStart.Time == (time.Time{}) {
				start = time.Now()
			} else {
				log.Println("no start date specified, using current date")
				start = m.DateStart.Time
			}
			if m.DateEnd.Time == (time.Time{}) {
				end = acc.CreatedAt.Time
				log.Println("no end date specified, using account creation date")
			} else {
				end = m.DateEnd.Time
			}
			log.Println("walking transactions")
			err = c.WalkTransactions(&ac, start, end, m.NewOnly)
			if err != nil {
				log.Panic("error walking transactions", err)
			}
		// list accounts
		case "accounts":
			accounts, err := ac.Client.GetAccounts()
			if err != nil {
				log.Panic("error getting accounts", err)
			}
			cli.PrintResult("accounts", accounts)
		// get account balance
		case "balance":
			balance, err := ac.Client.GetAccountBalance(&acc)
			if err != nil {
				log.Panic("error getting account balance", err)
			}
			cli.PrintResult("balance", balance)
		// get recurring payments
		case "recurring":
			repayments, err := ac.Client.GetRecurringPayments(&acc)
			if err != nil {
				log.Panic("error getting recurring payments", err)
			}
			cli.PrintResult("recurring payments", repayments)
			// get direct debits
		case "dd":
			dd, err := ac.Client.GetDirectDebitMandates()
			if err != nil {
				log.Panic("error getting direct debits", err)
			}
			cli.PrintResult("direct debits", dd)
		// get transactions
		case "client":
			items, err := ac.Client.GetFeedItems(&acc, acc.CreatedAt.Time)
			if err != nil {
				log.Panic("error getting transactions", err)
			}
			cli.PrintResult("transactions", items)
		}
	}
}
