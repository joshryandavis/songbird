package starling

import (
	"encoding/json"
)

type DirectDebitMandate struct {
	UID            string      `json:"uid"`
	Reference      string      `json:"reference"`
	Status         string      `json:"status"`
	Source         string      `json:"source"`
	Created        DateTime    `json:"created"`
	Cancelled      DateTime    `json:"cancelled"`
	NextDate       string      `json:"nextDate"`
	LastDate       string      `json:"lastDate"`
	OriginatorName string      `json:"originatorName"`
	OriginatorUID  string      `json:"originatorUid"`
	MerchantUID    string      `json:"merchantUid"`
	LastPayment    LastPayment `json:"lastPayment"`
	AccountUID     string      `json:"accountUid"`
	CategoryUID    string      `json:"categoryUid"`
}

type DirectDebitMandates struct {
	Mandates []DirectDebitMandate `json:"mandates"`
}

func (c *Client) GetDirectDebitMandates() ([]DirectDebitMandate, error) {
	var ret []DirectDebitMandate
	url := BaseEndpoint(DirectDebitsEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper DirectDebitMandates
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Mandates
	return ret, nil
}
