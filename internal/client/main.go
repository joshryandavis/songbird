package client

import (
	"github.com/joshryandavis/songbird/starling"

	"github.com/joshryandavis/songbird/internal/config"
)

type Client struct {
	Cfg config.Cfg
	Api []Api
}

type Api struct {
	Token   string
	Account starling.Account
	Client  *starling.Client
}

func New(tokens []string, cfg config.Cfg) *Client {
	ret := new(Client)
	ret.Cfg = cfg
	for _, token := range tokens {
		ret.Api = append(ret.Api, Api{
			Token:  token,
			Client: starling.New(token),
		})
	}
	for i := range ret.Api {
		account, err := GetPrimary(ret.Api[i].Client)
		if err != nil {
			continue
		}
		ret.Api[i].Account = account
	}
	return ret
}
