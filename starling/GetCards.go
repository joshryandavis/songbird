package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetCards() ([]stmodels.Card, error) {
	var ret []stmodels.Card
	url := BaseEndpoint(CardsEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper stmodels.Cards
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Cards
	return ret, nil
}
