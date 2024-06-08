package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetAccountHolderName() (stmodels.AccountHolderName, error) {
	var ret stmodels.AccountHolderName
	url := BaseEndpoint(AccountHolderEndpoint)
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
