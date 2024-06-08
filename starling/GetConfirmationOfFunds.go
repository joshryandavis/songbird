package starling

import (
	"encoding/json"
	"fmt"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetConfirmationOfFunds(a *stmodels.Account, amountMinorUnits int64) (stmodels.ConfirmationOfFunds, error) {
	var ret stmodels.ConfirmationOfFunds
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
