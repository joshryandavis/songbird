package stmodels

type Addresses struct {
	Current  Address   `json:"current"`
	Previous []Address `json:"previous"`
}
