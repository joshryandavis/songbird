package starling

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetAccounts() ([]stmodels.Account, error) {
	var ret []stmodels.Account
	url := fmt.Sprintf("%s/%s", BaseUrlProd, AccountsEndpoint)
	res, err := c.Request("GET", url, "")
	log.Printf("request: %s", res)
	if err != nil {
		log.Panic("request error:", err)
		return ret, err
	}
	var wrapper stmodels.Accounts
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		log.Panic("json unmarshal error:", err)
		return ret, err
	}
	ret = wrapper.Accounts
	return ret, nil
}
