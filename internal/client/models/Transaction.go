package models

import (
	"time"

	"golang.org/x/text/currency"
)

type CounterParty struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Transaction struct {
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

func NewTransaction(m Transaction) Transaction {
	// make sure the amount is negative for OUT transactions
	if m.Direction == "OUT" {
		m.Amount = m.Amount * -1
	}
	return m
}
