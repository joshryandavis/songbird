package stmodels

type LastPayment struct {
	LastDate   string            `json:"lastDate"`
	LastAmount CurrencyAndAmount `json:"lastAmount"`
}
