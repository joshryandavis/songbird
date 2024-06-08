package starling

type MastercardFeedItem struct {
	MerchantIdentifier string    `json:"merchantIdentifier"`
	Mcc                int32     `json:"mcc"`
	PosTimestamp       LocalTime `json:"posTimestamp"`
	AuthorisationCode  string    `json:"authorisationCode"`
	CardLast4          string    `json:"cardLast4"`
}
