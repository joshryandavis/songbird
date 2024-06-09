package client

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/text/currency"

	"github.com/joshryandavis/songbird/starling"
)

type CounterParty struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Item struct {
	UID              string        `json:"uid"`
	AccountUID       string        `json:"accountUid"`
	CategoryUID      string        `json:"categoryUid"`
	DirectDebitUID   string        `json:"directDebitUID"`
	RecurringUID     string        `json:"recurringUID"`
	CounterParty     CounterParty  `json:"counterParty"`
	Created          time.Time     `json:"created"`
	Updated          time.Time     `json:"updated"`
	Amount           float64       `json:"amount"`
	Direction        string        `json:"direction"`
	Currency         currency.Unit `json:"currency"`
	SpendingCategory string        `json:"spendingCategory"`
	Reference        string        `json:"reference"`
	Source           string        `json:"source"`
	Status           string        `json:"status"`
	Country          string        `json:"country"`
	Note             Note          `json:"note"`
}

func (c *Client) GetItems(ac *Api, dt time.Time) ([]Item, error) {
	var err error
	var ret []Item
	log.Println("getting direct debit mandates")
	dd, err := ac.Client.GetDirectDebitMandates()
	if err != nil {
		return ret, err
	}
	log.Println("getting recurring payments")
	recurring, err := ac.Client.GetRecurringPayments(&ac.Account)
	if err != nil {
		return ret, err
	}
	log.Println("getting feed items")
	items, err := ac.Client.GetFeedItems(&ac.Account, dt)
	if err != nil {
		return ret, err
	}
	var wg sync.WaitGroup
	// ch := make(chan Item)
	for _, t := range items {
		if t.Amount.Amount.Amount == 0 {
			log.Println("skipping client", t.FeedItemUID, "as it has a zero amount")
			continue
		}
		wg.Add(1)
		go func(t starling.FeedItem) {
			defer wg.Done()
			newT, err := processItem(ac, recurring, dd, t, c)
			if err != nil {
				log.Fatal("error processing new transactions", err)
			}
			// ch <- newT
			// ret = append(ret, <-ch)
			ret = append(ret, newT)
		}(t)
	}
	return ret, nil
}

func processItem(ac *Api, recurring []starling.RecurringCardPayment, dd []starling.DirectDebitMandate, t starling.FeedItem, c *Client) (Item, error) {
	ret := Item{}
	note, err := parseNote(t)
	if err != nil {
		return ret, err
	}
	ret = Item{
		Note:           note,
		Amount:         normaliseAmount(&t),
		DirectDebitUID: getDirectDebitId(dd, t.Reference, t.CounterPartyName),
		RecurringUID:   getRecurringId(recurring, t.FeedItemUID),
		AccountUID:     ac.Account.AccountUID,
		CategoryUID:    t.CategoryUID,
		UID:            t.FeedItemUID,
		CounterParty: CounterParty{
			UID:  t.CounterPartyUID,
			Name: t.CounterPartyName,
			Type: t.CounterPartyType,
		},
		Created:          t.TransactionTime.Time,
		Updated:          t.UpdatedAt.Time,
		Direction:        t.Direction,
		Currency:         t.Amount.Currency.Currency,
		SpendingCategory: t.SpendingCategory,
		Reference:        t.Reference,
		Source:           t.Source,
		Status:           t.Status,
		Country:          t.Country,
	}
	err = c.autoCategorise(ac, &ret)
	if err != nil {
		log.Fatal("error auto categorising", err)
		return ret, err
	}
	return ret, nil
}

func normaliseAmount(i *starling.FeedItem) float64 {
	if i.Direction == "OUT" {
		return i.Amount.Amount.Amount * -1
	}
	return i.Amount.Amount.Amount
}

func getRecurringId(recurring []starling.RecurringCardPayment, feedItemUID string) string {
	for _, r := range recurring {
		if r.FeedItemUID == feedItemUID {
			return r.RecurringPaymentUID
		}
	}
	return ""
}

func getDirectDebitId(dd []starling.DirectDebitMandate, reference string, originatorName string) string {
	for _, d := range dd {
		if d.Reference == reference && d.OriginatorName == originatorName {
			return d.UID
		}
	}
	return ""
}

func parseNote(t starling.FeedItem) (Note, error) {
	var ret Note
	var js map[string]interface{}
	if t.UserNote != "" && json.Unmarshal([]byte(t.UserNote), &js) == nil {
		err := json.Unmarshal([]byte(t.UserNote), &ret)
		if err != nil {
			log.Fatal("error unmarshalling", t.UserNote, err)
			return ret, err
		}
	}
	return ret, nil
}

func (c *Client) autoCategorise(ac *Api, newT *Item) error {
	for _, cat := range c.Cfg.Categories {
		if cat.Parent != newT.SpendingCategory {
			continue
		}
		for _, name := range cat.Auto {
			if normaliseCategory(name) == normaliseCategory(newT.CounterParty.Name) {
				newNote := newT.Note
				newNote.Category = cat.Name
				if newT.Note.Category == newNote.Category {
					continue
				}
				log.Println("new found match", newT.CounterParty.Name, "in", cat.Name)
				_, err := UpdateNote(newT.Note, ac, newT.UID, newT.CategoryUID, newNote)
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

func normaliseCategory(cat string) string {
	replacer := strings.NewReplacer("-", "", "_", "", "'", "", ".", "")
	cat = replacer.Replace(strings.ToLower(strings.TrimSpace(cat)))
	return cat
}
