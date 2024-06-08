package starling

type RecurringCardPayment struct {
	FeedItemUID         string            `json:"feedItemUid"`
	RecurringPaymentUID string            `json:"recurringPaymentUid"`
	AccountUID          string            `json:"accountUid"`
	CounterPartyUID     string            `json:"counterPartyUid"`
	CounterPartyName    string            `json:"counterPartyName"`
	Status              string            `json:"status"`
	LatestFeedItemUID   string            `json:"latestFeedItemUid"`
	LatestPaymentDate   string            `json:"latestPaymentDate"`
	LatestPaymentAmount CurrencyAndAmount `json:"latestPaymentAmount"`
}
