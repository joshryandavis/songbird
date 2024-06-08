package stmodels

type Balance struct {
	ClearedBalance        SignedCurrencyAndAmount `json:"clearedBalance"`
	EffectiveBalance      SignedCurrencyAndAmount `json:"effectiveBalance"`
	PendingTransactions   SignedCurrencyAndAmount `json:"pendingTransactions"`
	AcceptedOverdraft     SignedCurrencyAndAmount `json:"acceptedOverdraft"`
	Amount                SignedCurrencyAndAmount `json:"amount"`
	TotalClearedBalance   SignedCurrencyAndAmount `json:"totalClearedBalance"`
	TotalEffectiveBalance SignedCurrencyAndAmount `json:"totalEffectiveBalance"`
}
