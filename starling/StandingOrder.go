package starling

type StandingOrder struct {
	Description             string                  `json:"description"`
	PaymentOrderUid         string                  `json:"paymentOrderUid"`
	Amount                  CurrencyAndAmount       `json:"amount"`
	Reference               string                  `json:"reference"`
	PayeeUid                string                  `json:"payeeUid"`
	PayeeAccountUid         string                  `json:"payeeAccountUid"`
	StandingOrderRecurrence StandingOrderRecurrence `json:"standingOrderRecurrence"`
	NextDate                string                  `json:"nextDate"`
	CancelledAt             string                  `json:"cancelledAt"`
	UpdatedAt               string                  `json:"updatedAt"`
	SpendingCategory        string                  `json:"spendingCategory"`
	CategoryUid             string                  `json:"categoryUid"`
}
