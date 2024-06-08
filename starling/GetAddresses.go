package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetAddresses() (stmodels.Addresses, error) {
	var ret stmodels.Addresses
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
