package starling

import (
	"encoding/json"
)

func (c *Client) GetCards() ([]Card, error) {
	var ret []Card
	url := BaseEndpoint(CardsEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper Cards
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Cards
	return ret, nil
}
