package starling

type Client struct {
	Token string
}

func New(token string) *Client {
	ret := new(Client)
	ret.Token = token
	return ret
}
