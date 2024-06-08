package cli

import (
	_ "embed"
	"os"
	"time"

	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

const (
	name        = "songbird"
	description = "a cli interface for starling bank"
)

var Model struct {
	Verbose      verboseFlag `cmd:"" help:"Output verbose logs." short:"v"`
	Account      string      `cmd:"" help:"Account Name or \"all\"." short:"a" default:"all"`
	Balance      struct{}    `cmd:"" help:"List account balances." aliases:"bal"`
	Recurring    struct{}    `cmd:"" help:"List recurring card payments." aliases:"rp"`
	DirectDebit  struct{}    `cmd:"" help:"List direct debits." aliases:"dd"`
	Transactions struct{}    `cmd:"" help:"Parse and list all transactions." aliases:"tx"`
	Accounts     struct{}    `cmd:"" help:"List all accounts." aliases:"acc"`
	Walk         struct {
		NewOnly   bool            `help:"Only walk new transactions."`
		DateStart Date[startDate] `help:"Date to walk transactions from. Defaults to now, if not set."`
		DateEnd   Date[endDate]   `help:"Date to walk transactions to. Defaults to account creation date, if not set."`
	} `cmd:"" help:"Walk the transactions tree." aliases:"w"`
}

const inputDateFormat = "2006-01-02"

type verboseFlag bool
type startDate struct{}
type endDate struct{}

type Date[T any] struct {
	time.Time
}

func (d *Date[T]) UnmarshalFlag(value string) error {
	if value == "" {
		d.Time = time.Now()
		return nil
	}
	t, err := time.Parse(inputDateFormat, value)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func Enter() (*kong.Context, string) {
	setDefaultLogs()
	ctx := kong.Parse(&Model,
		kong.Name(name),
		kong.Description(description),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}))
	command := ctx.Command()
	log.WithFields(log.Fields{
		"command": command,
		"args":    os.Args[2:],
	}).Info("command")
	return ctx, command
}
