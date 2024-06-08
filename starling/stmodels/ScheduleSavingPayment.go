package stmodels

type ScheduleSavingPayment struct {
	TransferUid       string            `json:"transferUid"`
	RecurrenceRule    RecurrenceRule    `json:"recurrenceRule"`
	CurrencyAndAmount CurrencyAndAmount `json:"currencyAndAmount"`
	NextPaymentDate   string            `json:"nextPaymentDate"`
}
