package cmd

import (
	_ "embed"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/cli"
	"github.com/joshryandavis/songbird/internal/client"
	"github.com/joshryandavis/songbird/internal/config"
)

func Execute() {
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

	var wg sync.WaitGroup

	var output []string

	for _, ac := range c.Api {
		wg.Add(1)
		go func(ac client.Api) {
			defer wg.Done()
			out, err := execForClient(ac, c, command)
			if err != nil {
				log.Fatal("error processing new transactions", err)
			}
			output = append(output, out...)
		}(ac)
	}

	wg.Wait()

	for _, o := range output {
		fmt.Println(o)
	}
}

func execForClient(ac client.Api, c *client.Client, command string) ([]string, error) {
	acc := ac.Account
	var out []string

	switch command {
	// walk transactions
	case "w":
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
		err := c.WalkItems(&ac, start, end, m.NewOnly)
		if err != nil {
			return out, err
		}

	// list accounts
	case "acc":
	case "accounts":
		accounts, err := ac.Client.GetAccounts()
		if err != nil {
			return out, err
		}
		for _, a := range accounts {
			out = append(out, fmt.Sprintf("Account: %v", a))
		}

	// get account balance
	case "bal":
	case "balance":
		balance, err := ac.Client.GetAccountBalance(&acc)
		if err != nil {
			return out, err
		}
		out = append(out, fmt.Sprintf("Balance: %v", balance))

	// get recurring payments
	case "rec":
	case "recurring":
		repayments, err := ac.Client.GetRecurringPayments(&acc)
		if err != nil {
			return out, err
		}
		for _, r := range repayments {
			out = append(out, fmt.Sprintf("Recurring Payment: %v", r))
		}

	// get direct debits
	case "dd":
	case "direct-debits":
		dd, err := ac.Client.GetDirectDebitMandates()
		if err != nil {
			return out, err
		}
		for _, d := range dd {
			out = append(out, fmt.Sprintf("Direct Debit: %v", d))
		}

	// get transactions
	case "tx":
	case "transactions":
		items, err := c.GetItems(ac.Client, acc, acc.CreatedAt.Time)
		if err != nil {
			return out, err
		}
		for _, i := range items {
			out = append(out, fmt.Sprintf("Transaction: %v", i))
		}
	}
	return out, nil
}
