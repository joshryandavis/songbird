package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetRecurringPayments(a *stmodels.Account) ([]stmodels.RecurringCardPayment, error) {
	var ret []stmodels.RecurringCardPayment
	url := AccountEndpoint(a.AccountUID, RecurringPaymentsEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper stmodels.RecurringCardPayments
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.RecurringPayments
	return ret, nil
}
