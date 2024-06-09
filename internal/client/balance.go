package client

import (
	"github.com/joshryandavis/songbird/starling"
)

type Balances struct {
	Cleared   float64 `json:"cleared"`
	Effective float64 `json:"effective"`
	Pending   float64 `json:"pending"`
}

func (c *Client) GetBalance(api *Api, acc *starling.Account) (Balances, error) {
	var ret Balances
	bal, err := api.Client.GetAccountBalance(acc)
	if err != nil {
		return ret, err
	}
	ret = Balances{
		Cleared:   float64(bal.ClearedBalance.MinorUnits / 100),
		Effective: float64(bal.EffectiveBalance.MinorUnits / 100),
		Pending:   float64(bal.PendingTransactions.MinorUnits / 100),
	}
	return ret, nil
}
