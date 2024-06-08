package starling

import (
	"encoding/json"
	"strconv"

	"golang.org/x/text/currency"
)

type CurrencyAndAmount struct {
	Currency Currency `json:"currency"`
	Amount   Amount   `json:"minorUnits"`
}

type Currency struct {
	Currency currency.Unit
}

func (c *Currency) UnmarshalJSON(data []byte) error {
	var err error
	var tag string
	err = json.Unmarshal(data, &tag)
	if err != nil {
		return err
	}
	c.Currency, err = currency.ParseISO(tag)
	return err
}

func (c *Currency) String() string {
	return c.Currency.String()
}

type Amount struct {
	Amount float64
}

func (c *Amount) UnmarshalJSON(data []byte) error {
	var err error
	c.Amount, err = strconv.ParseFloat(string(data), 64)
	c.Amount = c.Amount / 100
	return err
}

func (c *Amount) String() string {
	return strconv.FormatFloat(c.Amount, 'f', 2, 64)
}

type CurrencyFlag struct {
	Enabled  bool   `json:"enabled"`
	Currency string `json:"currency"`
}

type SignedCurrencyAndAmount struct {
	Currency   string `json:"currency"`
	MinorUnits int64  `json:"minorUnits"`
}
