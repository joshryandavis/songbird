package starling

type ScheduleSavingPayment struct {
	TransferUid       string            `json:"transferUid"`
	RecurrenceRule    RecurrenceRule    `json:"recurrenceRule"`
	CurrencyAndAmount CurrencyAndAmount `json:"currencyAndAmount"`
	NextPaymentDate   string            `json:"nextPaymentDate"`
}

type NextPayment struct {
	Currency   string `json:"currency"`
	MinorUnits int    `json:"minorUnits"`
}

type RecurrenceRule struct {
	StartDate string   `json:"startDate"`
	Frequency string   `json:"frequency"`
	Interval  int      `json:"interval"`
	Count     int      `json:"count"`
	UntilDate string   `json:"untilDate"`
	WeekStart string   `json:"weekStart"`
	Days      []string `json:"days"`
	MonthDay  int      `json:"monthDay"`
	MonthWeek int      `json:"monthWeek"`
}

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

type ScheduledPaymentResponse struct {
	ScheduledPayments []ScheduledPayment `json:"scheduledPayments"`
}
