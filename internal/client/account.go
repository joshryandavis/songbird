package client

import (
	"fmt"

	"github.com/joshryandavis/songbird/starling"
)

func GetPrimary(c *starling.Client) (starling.Account, error) {
	var ret starling.Account
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
