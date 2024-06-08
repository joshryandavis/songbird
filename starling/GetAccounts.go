package starling

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (c *Client) GetAccounts() ([]Account, error) {
	var ret []Account
	url := fmt.Sprintf("%s/%s", BaseUrlProd, AccountsEndpoint)
	res, err := c.Request("GET", url, "")
	log.Printf("request: %s", res)
	if err != nil {
		log.Panic("request error:", err)
		return ret, err
	}
	var wrapper Accounts
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		log.Panic("json unmarshal error:", err)
		return ret, err
	}
	ret = wrapper.Accounts
	return ret, nil
}
