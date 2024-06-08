package starling

import (
	"encoding/json"
)

type Addresses struct {
	Current  Address   `json:"current"`
	Previous []Address `json:"previous"`
}

func (c *Client) GetAddresses() (Addresses, error) {
	var ret Addresses
	url := BaseEndpoint(AddressesEndpoint)
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
