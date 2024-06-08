package starling

import (
	"encoding/json"
	"fmt"
)

type ConfirmationOfFunds struct {
	RequestedAmountAvailableToSpend                 bool `json:"requestedAmountAvailableToSpend"`
	AccountWouldBeInOverdraftIfRequestedAmountSpent bool `json:"accountWouldBeInOverdraftIfRequestedAmountSpent"`
}

func (c *Client) GetConfirmationOfFunds(a *Account, amountMinorUnits int64) (ConfirmationOfFunds, error) {
	var ret ConfirmationOfFunds
	url := AccountEndpoint(a.AccountUID, fmt.Sprintf("%s/%s/%d", ConfirmationOfFundsEndpoint, a.AccountUID, amountMinorUnits))
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(res, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
