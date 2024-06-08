package cli

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

const (
	name            = "songbird"
	description     = "a cli interface for starling bank"
	inputDateFormat = "2006-01-02"
)

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

func (d *Date[T]) String() string {
	return d.Time.Format(inputDateFormat)
}

func (d verboseFlag) BeforeApply() error {
	log.SetLevel(log.DebugLevel)
	return nil
}

func setDefaultLogs() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

var Model struct {
	Verbose      verboseFlag `cmd:"" help:"Output verbose logs." short:"v"`
	Account      string      `cmd:"" help:"Account Name or \"all\"." short:"a" default:"all"`
	Balance      struct{}    `cmd:"" help:"Account balance." aliases:"bal"`
	Recurring    struct{}    `cmd:"" help:"Recurring payments." aliases:"rp"`
	DirectDebit  struct{}    `cmd:"" help:"Direct debits." aliases:"dd"`
	Transactions struct{}    `cmd:"" help:"Transactions." aliases:"tx"`
	Accounts     struct{}    `cmd:"" help:"Account list." aliases:"acc"`
	Walk         struct {
		NewOnly   bool            `help:"Only walk new transactions."`
		DateStart Date[startDate] `help:"Date to walk transactions from. Defaults to now, if not set."`
		DateEnd   Date[endDate]   `help:"Date to walk transactions to. Defaults to account creation date, if not set."`
	} `cmd:"" help:"Walk the transactions tree." aliases:"w"`
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

func PrintResult(label string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s: %s\n", label, string(out))
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Fprint(os.Stderr, label+" ")
		if err != nil {
			return ""
		}
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func YesNoPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		_, err := fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)
		if err != nil {
			return false
		}
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "" {
			return def
		}
		s = strings.ToLower(s)
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}
