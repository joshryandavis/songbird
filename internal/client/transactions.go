package client

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/client/models"
	"github.com/joshryandavis/songbird/starling"
	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetTransactions(ac *starling.Client, acc stmodels.Account, dt time.Time) ([]models.Transaction, error) {
	var err error
	var rec []models.Transaction
	log.Println("getting direct debit mandates")
	dd, err := ac.GetDirectDebitMandates()
	if err != nil {
		return rec, err
	}
	log.Println("getting recurring payments")
	recurring, err := ac.GetRecurringPayments(&acc)
	if err != nil {
		return rec, err
	}
	log.Println("getting feed items")
	items, err := ac.GetFeedItems(&acc, dt)
	if err != nil {
		return rec, err
	}

	var wg sync.WaitGroup
	ch := make(chan models.Transaction)
	for _, t := range items {
		if t.Amount.Amount.Amount == 0 {
			log.Println("skipping client", t.FeedItemUID, "as it has a zero amount")
			continue
		}
		wg.Add(1)
		go func(t stmodels.FeedItem) {
			defer wg.Done()
			newT, err := processNewTransactions(ac, acc, recurring, dd, t, c)
			if err != nil {
				log.Fatal("error processing new transactions", err)
			}
			ch <- newT
		}(t)
	}
	err = writeToTmp(rec)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func processNewTransactions(ac *starling.Client, acc stmodels.Account, recurring []stmodels.RecurringCardPayment, dd []stmodels.DirectDebitMandate, t stmodels.FeedItem, c *Client) (models.Transaction, error) {
	ret := models.Transaction{}
	var n models.Note
	var directDebitUid string
	var recurringUid string
	for _, r := range recurring {
		if r.FeedItemUID == t.FeedItemUID {
			recurringUid = r.RecurringPaymentUID
			break
		}
	}
	for _, d := range dd {
		if d.Reference == t.Reference && d.OriginatorName == t.CounterPartyName {
			directDebitUid = d.UID
			break
		}
	}
	if t.UserNote != "" && isJson(t.UserNote) {
		err := json.Unmarshal([]byte(t.UserNote), &n)
		if err != nil {
			log.Fatal("error unmarshalling", t.UserNote, err)
			return ret, err
		}
	}
	ret = NewTransaction(models.Transaction{
		UID:            t.FeedItemUID,
		AccountUID:     acc.AccountUID,
		DirectDebitUID: directDebitUid,
		RecurringUID:   recurringUid,
		CategoryUID:    t.CategoryUID,
		CounterParty: models.CounterParty{
			UID:  t.CounterPartyUID,
			Name: t.CounterPartyName,
			Type: t.CounterPartyType,
		},
		Created:          t.TransactionTime.Time,
		Updated:          t.UpdatedAt.Time,
		Amount:           t.Amount.Amount.Amount,
		Direction:        t.Direction,
		Currency:         t.Amount.Currency.Currency,
		SpendingCategory: t.SpendingCategory,
		Reference:        t.Reference,
		Source:           t.Source,
		Status:           t.Status,
		Country:          t.Country,
		Note:             n,
	})
	err := c.autoCategorise(ac, acc, &ret)
	if err != nil {
		log.Fatal("error auto categorising", err)
		return ret, err
	}
	return ret, nil
}

func writeToTmp(output []models.Transaction) error {
	tmp, err := os.CreateTemp("tmp", "transactions-*.json")
	if err != nil {
		log.Fatal("error creating temp file", err)
		return err
	}
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		log.Fatal("error marshalling output to json", err)
		return err
	}
	_, err = tmp.WriteString(string(jsonOutput))
	if err != nil {
		log.Fatal("error writing json to temp file", err)
		return err
	}
	return nil
}

func (c *Client) autoCategorise(ac *starling.Client, acc stmodels.Account, newT *models.Transaction) error {
	for _, cat := range c.Cfg.Categories {
		if cat.Parent != newT.SpendingCategory {
			continue
		}
		for _, name := range cat.Auto {
			if normaliseCat(name) == normaliseCat(newT.CounterParty.Name) {
				newNote := newT.Note
				newNote.Category = cat.Name
				if newT.Note.Category == newNote.Category {
					continue
				}
				fmt.Println("new found match", newT.CounterParty.Name, "in", cat.Name)
				_, err := UpdateNote(newT.Note, ac, acc, newT.UID, newT.CategoryUID, newNote)
				if err != nil {
					log.Fatal("error updating note", err)
					return err
				}
				break
			}
		}
	}
	return nil
}

func normaliseCat(cat string) string {
	cat = strings.ToLower(cat)
	cat = strings.Replace(cat, "-", "", -1)
	cat = strings.Replace(cat, "_", "", -1)
	cat = strings.Replace(cat, "'", "", -1)
	cat = strings.Replace(cat, ".", "", -1)
	cat = strings.Trim(cat, " ")
	return cat
}

func isJson(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func NewTransaction(m models.Transaction) models.Transaction {
	// make sure the amount is negative for OUT transactions
	if m.Direction == "OUT" {
		m.Amount = m.Amount * -1
	}
	return m
}

func UpdateNote(n models.Note, c *starling.Client, acc stmodels.Account, tUID string, categoryUID string, newNote models.Note) (models.Note, error) {
	var rec = n
	if newNote == n {
		log.WithFields(log.Fields{
			"newNote": newNote,
			"oldNote": n,
		}).Printf("note is the same, skipping update")
		return rec, nil
	}
	rec = newNote
	noteJson, err := json.Marshal(rec)
	if err != nil {
		return rec, err
	}
	log.Println("updating note for client", tUID)
	err = c.PutFeedItemNote(&acc, categoryUID, tUID, string(noteJson))
	return rec, err
}

func (c *Client) GetSpendingCategories() ([]string, error) {
	return nil, nil
}
