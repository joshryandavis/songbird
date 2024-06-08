package starling

import (
	"encoding/json"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetDirectDebitMandates() ([]stmodels.DirectDebitMandate, error) {
	var ret []stmodels.DirectDebitMandate
	url := BaseEndpoint(DirectDebitsEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper stmodels.WrapperDirectDebitMandates
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Mandates
	return ret, nil
}
