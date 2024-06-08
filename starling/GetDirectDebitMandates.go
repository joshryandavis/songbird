package starling

import (
	"encoding/json"
)

func (c *Client) GetDirectDebitMandates() ([]DirectDebitMandate, error) {
	var ret []DirectDebitMandate
	url := BaseEndpoint(DirectDebitsEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper WrapperDirectDebitMandates
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Mandates
	return ret, nil
}
