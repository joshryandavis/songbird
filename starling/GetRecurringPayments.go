package starling

import (
	"encoding/json"
)

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
