package client

import (
	"fmt"

	"github.com/joshryandavis/songbird/starling"
	"github.com/joshryandavis/songbird/starling/stmodels"
)

func GetPrimary(c *starling.Client) (stmodels.Account, error) {
	var ret stmodels.Account
	accounts, err := c.GetAccounts()
	if err != nil {
		return ret, err
	}
	if len(accounts) == 0 {
		return ret, fmt.Errorf("no accounts found")
	}
	ret = accounts[0]
	return ret, nil
}
