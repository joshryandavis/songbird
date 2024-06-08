package starling

import (
	"encoding/json"
)

type Address struct {
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Line3       string `json:"line3"`
	PostTown    string `json:"postTown"`
	PostCode    string `json:"postCode"`
	CountryCode string `json:"countryCode"`
}

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
