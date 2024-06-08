package starling

type NextPayment struct {
	Currency   string `json:"currency"`
	MinorUnits int    `json:"minorUnits"`
}
