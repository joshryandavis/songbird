package client

import (
	_ "embed"
	"fmt"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/cli"
	"github.com/joshryandavis/songbird/internal/client/models"
)

const DateFormat = "2006-01-02 15:04"

func (c *Client) WalkTransactions(ac *ApiClient, start time.Time, end time.Time, newOnly bool) error {
	acc := ac.Account
	log.Println("getting transactions from", start, "to", end)
	trans, err := c.GetTransactions(ac.Client, acc, end)
	if err != nil {
		return err
	}
	for _, t := range trans {
		if newOnly && (!noteIsNew(t.Note.Updated) || noteUpdatedInDateRange(t.Note.Updated, start, end)) {
			log.Println("skipping client", t.UID, "as it is not new.")
			continue
		}
		printTransaction(t)
		updatedNote := promptNoteUpdate(t.Note)
		if updatedNote == t.Note {
			log.Println("note is the same, skipping update")
			continue
		}
		t.Note, err = UpdateNote(t.Note, ac.Client, acc, t.UID, t.CategoryUID, updatedNote)
		if err != nil {
			break
		}
	}
	return err
}

func noteIsNew(updated time.Time) bool {
	return updated == (time.Time{})
}

func noteUpdatedInDateRange(updated time.Time, start time.Time, end time.Time) bool {
	return updated.After(start) && updated.Before(end)
}

func promptNoteUpdate(note models.Note) models.Note {
	f := reflect.ValueOf(&note).Elem()
	for i := 0; i < f.NumField(); i++ {
		field := f.Field(i)
		fieldType := field.Type().String()
		fieldName := f.Type().Field(i).Name
		if fieldType == "string" {
			field.SetString(cli.StringPrompt(fmt.Sprintf("%s [%s]:", fieldName, field.String())))
		} else if fieldType == "bool" {
			field.SetBool(cli.YesNoPrompt(fmt.Sprintf("%s [%t]:", fieldName, field.Bool()), false))
		}
	}
	return note
}

func printTransaction(t models.Transaction) {
	fmt.Println(" ")
	fmt.Println(t.Created.Format(DateFormat), t.CounterParty.Name, t.Amount)
	fmt.Println("Spend Category:", t.SpendingCategory)
	if t.Note.Updated == (time.Time{}) {
		fmt.Println("note updated: never")
	} else {
		fmt.Println("note updated:", t.Note.Updated.Format(DateFormat))
	}
	fmt.Println(" ")
}