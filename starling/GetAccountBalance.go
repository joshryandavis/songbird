package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetAccountBalance(a *stmodels.Account) (stmodels.Balance, error) {
	var ret stmodels.Balance
	url := AccountEndpoint(a.AccountUID, BalanceEndpoint)
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
