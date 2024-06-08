package starling

type Account struct {
	AccountUID      string   `json:"accountUid"`
	AccountType     string   `json:"accountType"`
	DefaultCategory string   `json:"defaultCategory"`
	Currency        string   `json:"currency"`
	CreatedAt       DateTime `json:"createdAt"`
	Name            string   `json:"name"`
}
