package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetAccountIdentifiers(a *stmodels.Account) (stmodels.AccountIdentifiers, error) {
	var ret stmodels.AccountIdentifiers
	url := AccountEndpoint(a.AccountUID, IdentifiersEndpoint)
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
