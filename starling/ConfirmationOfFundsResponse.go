package starling

type ConfirmationOfFunds struct {
	RequestedAmountAvailableToSpend                 bool `json:"requestedAmountAvailableToSpend"`
	AccountWouldBeInOverdraftIfRequestedAmountSpent bool `json:"accountWouldBeInOverdraftIfRequestedAmountSpent"`
}
