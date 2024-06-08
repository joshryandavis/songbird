package starling

import (
	"encoding/json"
	"fmt"
	"time"
)

type FeedItems struct {
	FeedItems []FeedItem `json:"feedItems"`
}

type AssociatedFeedRoundUp struct {
	GoalCategoryUid string            `json:"goalCategoryUid"`
	Amount          CurrencyAndAmount `json:"amount"`
}

type BatchPaymentDetails struct {
	BatchPaymentUID  string `json:"batchPaymentUid"`
	BatchPaymentType string `json:"batchPaymentType"`
}

type FeedItem struct {
	FeedItemUID                        string                `json:"feedItemUid"`
	CategoryUID                        string                `json:"categoryUid"`
	Amount                             CurrencyAndAmount     `json:"amount"`
	SourceAmount                       CurrencyAndAmount     `json:"sourceAmount"`
	Direction                          string                `json:"direction"`
	UpdatedAt                          DateTime              `json:"updatedAt"`
	TransactionTime                    DateTime              `json:"transactionTime"`
	SettlementTime                     DateTime              `json:"settlementTime"`
	RetryAllocationUntilTime           DateTime              `json:"retryAllocationUntilTime"`
	Source                             string                `json:"source"`
	SourceSubType                      string                `json:"sourceSubType"`
	Status                             string                `json:"status"`
	TransactingApplicationUserUID      string                `json:"transactingApplicationUserUid"`
	CounterPartyType                   string                `json:"counterPartyType"`
	CounterPartyUID                    string                `json:"counterPartyUid"`
	CounterPartyName                   string                `json:"counterPartyName"`
	CounterPartySubEntityUID           string                `json:"counterPartySubEntityUid"`
	CounterPartySubEntityName          string                `json:"counterPartySubEntityName"`
	CounterPartySubEntityIdentifier    string                `json:"counterPartySubEntityIdentifier"`
	CounterPartySubEntitySubIdentifier string                `json:"counterPartySubEntitySubIdentifier"`
	ExchangeRate                       float64               `json:"exchangeRate"`
	TotalFees                          float64               `json:"totalFees"`
	TotalFeeAmount                     CurrencyAndAmount     `json:"totalFeeAmount"`
	Reference                          string                `json:"reference"`
	Country                            string                `json:"country"`
	SpendingCategory                   string                `json:"spendingCategory"`
	UserNote                           string                `json:"userNote"`
	RoundUp                            AssociatedFeedRoundUp `json:"roundUp"`
	HasAttachment                      bool                  `json:"hasAttachment"`
	HasReceipt                         bool                  `json:"hasReceipt"`
	BatchPaymentDetails                BatchPaymentDetails   `json:"batchPaymentDetails"`
}

func (c *Client) GetFeedItems(a *Account, t0 time.Time) ([]FeedItem, error) {
	var ret []FeedItem
	timeSince := ParseTime(t0).String()
	url := BaseEndpoint(fmt.Sprintf("feed/account/%s/category/%s?accountUid=%s&categoryUid=%s&changesSince=%s", a.AccountUID, a.DefaultCategory, a.AccountUID, a.DefaultCategory, timeSince))
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper FeedItems
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.FeedItems
	return ret, nil
}
