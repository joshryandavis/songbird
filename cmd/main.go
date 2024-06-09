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
	switch command {
	// walk transactions
	case "w":
	case "walk":
		return walkCmd(ac, c)
	// list accounts
	case "acc":
	case "accounts":
		return accountsCmd(ac)
	// get account balance
	case "bal":
	case "balance":
		return balanceCmd(ac, c)
	// get recurring payments
	case "rec":
	case "recurring":
		return recurringCmd(ac)
	// get direct debits
	case "dd":
	case "direct-debits":
		return directDebitCmd(ac, c)
	// get transactions
	case "tx":
	case "transactions":
		return transactionsCmd(ac, c)
	}
	return []string{"no command given"}, nil
}

func walkCmd(ac client.Api, c *client.Client) ([]string, error) {
	var out []string
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
		end = ac.Account.CreatedAt.Time
		log.Println("no end date specified, using account creation date")
	} else {
		end = m.DateEnd.Time
	}
	log.Println("walking transactions")
	err := c.WalkItems(&ac, start, end, m.NewOnly)
	if err != nil {
		return out, err
	}
	return out, nil
}

func transactionsCmd(ac client.Api, c *client.Client) ([]string, error) {
	var out []string
	items, err := c.GetItems(&ac, ac.Account.CreatedAt.Time)
	if err != nil {
		return out, err
	}
	for _, i := range items {
		out = append(out, fmt.Sprintf("\nID: %v", i.UID))
		out = append(out, fmt.Sprintf("Amount: %v", i.Amount))
		out = append(out, fmt.Sprintf("Date: %v", i.Created.String()))
		out = append(out, fmt.Sprintf("Counterparty: %v", i.CounterParty.Name))
	}
	return out, nil
}

func directDebitCmd(ac client.Api, c *client.Client) ([]string, error) {
	var out []string
	dd, err := c.GetDirectDebits(&ac, &ac.Account)
	if err != nil {
		return out, err
	}
	for _, d := range dd {
		out = append(out, fmt.Sprintf("\nID: %v", d.UID))
		out = append(out, fmt.Sprintf("Reference: %v", d.Reference))
		out = append(out, fmt.Sprintf("Originator: %v", d.OriginatorName))
		out = append(out, fmt.Sprintf("Last Date: %v", d.LastDate))
		out = append(out, fmt.Sprintf("Probable Next Date: %v", d.ProbableNextDate))
		out = append(out, fmt.Sprintf("Last Amount: %v", d.LastPayment.LastAmount.Amount.Amount))
		out = append(out, fmt.Sprintf("Status: %v", d.Status))
	}
	return out, nil
}

func recurringCmd(ac client.Api) ([]string, error) {
	var out []string
	repayments, err := ac.Client.GetRecurringPayments(&ac.Account)
	if err != nil {
		return out, err
	}
	for _, r := range repayments {
		out = append(out, fmt.Sprintf("\nID: %v", r.RecurringPaymentUID))
		out = append(out, fmt.Sprintf("Status: %v", r.Status))
		out = append(out, fmt.Sprintf("Receipient: %v", r.CounterPartyName))
		out = append(out, fmt.Sprintf("Last Amount: %v", r.LatestPaymentAmount.Amount.Amount))
		out = append(out, fmt.Sprintf("Last Date: %v", r.LatestPaymentDate))
	}
	return out, nil
}

func balanceCmd(ac client.Api, c *client.Client) ([]string, error) {
	var out []string
	out = append(out, fmt.Sprintf("\nAccount: %v", ac.Account.Name))
	balance, err := c.GetBalance(&ac, &ac.Account)
	if err != nil {
		return out, err
	}
	out = append(out, fmt.Sprintf("Effective: %v", balance.Effective))
	out = append(out, fmt.Sprintf("Cleared: %v", balance.Cleared))
	out = append(out, fmt.Sprintf("Pending: %v", balance.Pending))
	return out, nil
}

func accountsCmd(ac client.Api) ([]string, error) {
	var out []string
	accounts, err := ac.Client.GetAccounts()
	if err != nil {
		return out, err
	}
	for _, a := range accounts {
		out = append(out, fmt.Sprintf("\nAccount: %v", a))
	}
	return out, nil
}
