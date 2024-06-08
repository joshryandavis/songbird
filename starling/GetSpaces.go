package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetSpaces(a *stmodels.Account) (stmodels.Spaces, error) {
	var ret stmodels.Spaces
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
