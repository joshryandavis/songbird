package stmodels

type ScheduledPayment struct {
	AccountHolderUid  string         `json:"accountHolderUid"`
	PaymentOrderUid   string         `json:"paymentOrderUid"`
	CategoryUid       string         `json:"categoryUid"`
	NextPaymentAmount NextPayment    `json:"nextPaymentAmount"`
	Reference         string         `json:"reference"`
	PayeeUid          string         `json:"payeeUid"`
	PayeeAccountUid   string         `json:"payeeAccountUid"`
	RecipientName     string         `json:"recipientName"`
	RecurrenceRule    RecurrenceRule `json:"recurrenceRule"`
	StartDate         string         `json:"startDate"`
	NextDate          string         `json:"nextDate"`
	EndDate           string         `json:"endDate"`
	PaymentType       string         `json:"paymentType"`
	SpendingCategory  string         `json:"spendingCategory"`
}
