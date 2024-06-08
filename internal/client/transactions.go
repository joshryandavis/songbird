package client

import (
	"encoding/json"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/client/models"
	"github.com/joshryandavis/songbird/starling"
	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetTransactions(ac *starling.Client, acc stmodels.Account, dt time.Time) ([]models.Transaction, error) {
	var err error
	var ret []models.Transaction
	log.Println("getting direct debit mandates")
	dd, err := ac.GetDirectDebitMandates()
	if err != nil {
		return ret, err
	}
	log.Println("getting recurring payments")
	recurring, err := ac.GetRecurringPayments(&acc)
	if err != nil {
		return ret, err
	}
	log.Println("getting feed items")
	items, err := ac.GetFeedItems(&acc, dt)
	if err != nil {
		return ret, err
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
			newT, err := processTransaction(ac, acc, recurring, dd, t, c)
			if err != nil {
				log.Fatal("error processing new transactions", err)
			}
			ch <- newT
			ret = append(ret, <-ch)
		}(t)
	}
	return ret, nil
}

func processTransaction(ac *starling.Client, acc stmodels.Account, recurring []stmodels.RecurringCardPayment, dd []stmodels.DirectDebitMandate, t stmodels.FeedItem, c *Client) (models.Transaction, error) {
	ret := models.Transaction{}
	note, err := getNote(t)
	if err != nil {
		return ret, err
	}
	ret = models.NewTransaction(models.Transaction{
		Note:           note,
		UID:            t.FeedItemUID,
		AccountUID:     acc.AccountUID,
		DirectDebitUID: getDirectDebitId(dd, t.Reference, t.CounterPartyName),
		RecurringUID:   getRecurringId(recurring, t.FeedItemUID),
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
	})
	err = c.autoCategorise(ac, acc, &ret)
	if err != nil {
		log.Fatal("error auto categorising", err)
		return ret, err
	}
	return ret, nil
}

func getRecurringId(recurring []stmodels.RecurringCardPayment, feedItemUID string) string {
	for _, r := range recurring {
		if r.FeedItemUID == feedItemUID {
			return r.RecurringPaymentUID
		}
	}
	return ""
}

func getDirectDebitId(dd []stmodels.DirectDebitMandate, reference string, originatorName string) string {
	for _, d := range dd {
		if d.Reference == reference && d.OriginatorName == originatorName {
			return d.UID
		}
	}
	return ""
}

func getNote(t stmodels.FeedItem) (models.Note, error) {
	var ret models.Note
	if t.UserNote != "" && isJson(t.UserNote) {
		err := json.Unmarshal([]byte(t.UserNote), &ret)
		if err != nil {
			log.Fatal("error unmarshalling", t.UserNote, err)
			return ret, err
		}
	}
	return ret, nil
}
