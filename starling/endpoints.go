package starling

import "fmt"

const (
	BaseUrlProd = "https://api.starlingbank.com/api/v2"
	// BaseUrlSb   = "https://api-sandbox.starlingbank.com/api/v2"
	AccountHolderEndpoint       = "account-holder"
	AccountsEndpoint            = "accounts"
	AddressesEndpoint           = "addresses"
	BalanceEndpoint             = "balance"
	ConfirmationOfFundsEndpoint = "confirmation-of-funds"
	IdentifiersEndpoint         = "identifiers"
	CardsEndpoint               = "cards"
	DirectDebitsEndpoint        = "direct-debit/mandates"
	RecurringPaymentsEndpoint   = "recurring-payment"
	SpacesEndpoint              = "spaces"
)

func BaseEndpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", BaseUrlProd, endpoint)
}

func AccountEndpoint(a string, endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", fmt.Sprintf("%s/%s", BaseUrlProd, AccountsEndpoint), a, endpoint)
}
