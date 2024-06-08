package starling

import (
	"encoding/json"
)

type AccountHolder struct {
	AccountHolderUid  string `json:"accountHolderUid"`
	AccountHolderType string `json:"accountHolderType"`
}

type AccountHolderName struct {
	AccountHolderName string `json:"accountHolderName"`
}

func (c *Client) GetAccountHolder() (AccountHolder, error) {
	var ret AccountHolder
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
