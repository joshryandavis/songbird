package starling

import (
	"encoding/json"
)

func (c *Client) GetAccountIdentifiers(a *Account) (AccountIdentifiers, error) {
	var ret AccountIdentifiers
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
