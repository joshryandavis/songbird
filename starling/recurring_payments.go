package starling

import (
	"encoding/json"
)

type RecurringCardPayment struct {
	FeedItemUID         string            `json:"feedItemUid"`
	RecurringPaymentUID string            `json:"recurringPaymentUid"`
	AccountUID          string            `json:"accountUid"`
	CounterPartyUID     string            `json:"counterPartyUid"`
	CounterPartyName    string            `json:"counterPartyName"`
	Status              string            `json:"status"`
	LatestFeedItemUID   string            `json:"latestFeedItemUid"`
	LatestPaymentDate   string            `json:"latestPaymentDate"`
	LatestPaymentAmount CurrencyAndAmount `json:"latestPaymentAmount"`
}

type RecurringCardPayments struct {
	RecurringPayments []RecurringCardPayment `json:"recurringPayments"`
}

func (c *Client) GetRecurringPayments(a *Account) ([]RecurringCardPayment, error) {
	var ret []RecurringCardPayment
	url := AccountEndpoint(a.AccountUID, RecurringPaymentsEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper RecurringCardPayments
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.RecurringPayments
	return ret, nil
}
