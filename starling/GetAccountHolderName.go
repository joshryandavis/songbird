package starling

import (
	"encoding/json"
)

func (c *Client) GetAccountHolderName() (AccountHolderName, error) {
	var ret AccountHolderName
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
