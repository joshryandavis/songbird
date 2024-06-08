package starling

import (
	"encoding/json"
)

func (c *Client) GetSpaces(a *Account) (Spaces, error) {
	var ret Spaces
	url := AccountEndpoint(a.AccountUID, SpacesEndpoint)
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
