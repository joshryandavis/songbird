package client

import (
	"github.com/joshryandavis/songbird/starling"

	"github.com/joshryandavis/songbird/internal/config"
)

type Client struct {
	Cfg     config.Config
	Clients []ApiClient
}

type ApiClient struct {
	Account starling.Account
	Token   string
	Client  *starling.Client
}

func New(tokens []string, cfg config.Config) *Client {
	ret := new(Client)
	ret.Cfg = cfg
	for _, token := range tokens {
		ret.Clients = append(ret.Clients, ApiClient{
			Token:  token,
			Client: starling.New(token),
		})
	}
	for i := range ret.Clients {
		account, err := GetPrimary(ret.Clients[i].Client)
		if err != nil {
			continue
		}
		ret.Clients[i].Account = account
	}
	return ret
}
