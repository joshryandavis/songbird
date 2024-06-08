package stmodels

type StandingOrderRecurrence struct {
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	Frequency   string `json:"frequency"`
	Interval    int32  `json:"interval"`
	Count       int32  `json:"count"`
	UntilDate   string `json:"untilDate"`
}
